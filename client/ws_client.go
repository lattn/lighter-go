package client

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/bytedance/sonic"
	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
)

const MainnetWsEndpoint = "wss://mainnet.zklighter.elliot.ai/stream"

type AuthFunc func(ctx context.Context, dur time.Duration) (string, error)

type WsClient struct {
	lk       sync.RWMutex
	subjects []map[string]any
	write    chan any
	read     chan []byte
	hooks    []func(b []byte)
	onError  func(err error)
	dialOpts *websocket.DialOptions
}

type WsClientOption func(*WsClient)

func WithWsOnError(onError func(err error)) WsClientOption {
	return func(client *WsClient) {
		client.onError = onError
	}
}

func WithDialOptions(opts *websocket.DialOptions) WsClientOption {
	return func(client *WsClient) {
		client.dialOpts = opts
	}
}

func NewWsClient(opts ...WsClientOption) *WsClient {
	c := &WsClient{
		write: make(chan any, 128),
		read:  make(chan []byte, 128),
	}
	for _, opt := range opts {
		opt(c)
	}
	if c.onError != nil {
		c.onError = func(err error) {}
	}
	return c
}

func (c *WsClient) Dial(ctx context.Context, endpoint string) {
	conn, _, err := websocket.Dial(ctx, endpoint, c.dialOpts)
	if err != nil {
		c.onError(err)
		return
	}
	defer conn.Close(websocket.StatusNormalClosure, "")

	c.lk.RLock()
	for _, sub := range c.subjects {
		err = wsjson.Write(ctx, conn, sub)
		if err != nil {
			c.onError(err)
			c.lk.RUnlock()
			return
		}
	}
	c.lk.RUnlock()

	closed := make(chan struct{})
	once := sync.Once{}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			once.Do(func() {
				close(closed)
			})
		}()
		for {
			var b json.RawMessage
			err = wsjson.Read(ctx, conn, &b)
			if err != nil {
				c.onError(err)
				return
			}
			select {
			case c.read <- b:
			case <-closed:
				return
			}
		}
	}()
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			once.Do(func() {
				close(closed)
			})
		}()
		for {
			select {
			case <-ctx.Done():
				return
			case <-closed:
				return
			case msg := <-c.write:
				err = wsjson.Write(ctx, conn, msg)
				if err != nil {
					c.onError(err)
					return
				}
			}
		}
	}()

	wg.Wait()
}

func (c *WsClient) Keepalive(ctx context.Context, endpoint string) {
	for {
		c.Dial(ctx, endpoint)
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Second):
		}
	}
}

func (c *WsClient) AddHook(hook func(b []byte)) {
	c.lk.Lock()
	c.hooks = append(c.hooks, hook)
	c.lk.Unlock()
}

func (c *WsClient) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case read := <-c.read:
			node, _ := sonic.Get(read, "type")
			if node.Exists() {
				v, err := node.StrictString()
				if err != nil {
					c.onError(fmt.Errorf("parse type: %w", err))
					continue
				}
				if v == "ping" {
					c.write <- map[string]any{"type": "pong"}
					continue
				}
			}
			c.lk.RLock()
			for _, hook := range c.hooks {
				hook(read)
			}
			c.lk.RUnlock()
		}
	}
}

func (c *WsClient) Send(typ, channel string, auth AuthFunc) {
	v := map[string]any{"type": typ, "channel": channel}
	if auth != nil {
		a, err := auth(context.TODO(), time.Minute)
		if err != nil {
			c.onError(fmt.Errorf("get auth: %w", err))
			return
		}
		v["auth"] = a
	}
	c.lk.Lock()
	c.subjects = append(c.subjects, v)
	c.lk.Unlock()
	c.write <- v
}

func (c *WsClient) SubscribeMarketStats(marketIdxs ...int64) {
	if len(marketIdxs) == 0 {
		c.Send("subscribe", "market_stats/all", nil)
		return
	}
	for _, idx := range marketIdxs {
		c.Send("subscribe", fmt.Sprintf("market_stats/%d", idx), nil)
	}
}

func (c *WsClient) SubscribeAccountOrders(auth AuthFunc, accountIndex int64, marketIdxs ...int64) {
	if len(marketIdxs) == 0 {
		c.Send("subscribe", fmt.Sprintf("account_all_orders/%d", accountIndex), auth)
		return
	}
	for _, idx := range marketIdxs {
		c.Send("subscribe", fmt.Sprintf("account_orders/%d/%d", idx, accountIndex), auth)
	}
}
