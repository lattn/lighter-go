package client

const (
	CodeOK = 200
)

type ResultCode struct {
	Code    int32  `json:"code,example=200"`
	Message string `json:"message,omitempty"`
}

type NextNonce struct {
	ResultCode
	Nonce int64 `json:"nonce,example=722"`
}

type ApiKey struct {
	AccountIndex int64  `json:"account_index,example=3"`
	ApiKeyIndex  uint8  `json:"api_key_index,example=0"`
	Nonce        int64  `json:"nonce,example=722"`
	PublicKey    string `json:"public_key"`
}

type AccountApiKeys struct {
	ResultCode
	ApiKeys []*ApiKey `json:"api_keys"`
}

type TxHash struct {
	ResultCode
	TxHash string `json:"tx_hash,example=0x70997970C51812dc3A010C7d01b50e0d17dc79C8"`
}

type TransferFeeInfo struct {
	ResultCode
	TransferFee int64 `json:"transfer_fee_usdc"`
}

type Position struct {
	MarketId               int    `json:"market_id"`
	Symbol                 string `json:"symbol"`
	InitialMarginFraction  string `json:"initial_margin_fraction"`
	OpenOrderCount         int    `json:"open_order_count"`
	PendingOrderCount      int    `json:"pending_order_count"`
	PositionTiedOrderCount int    `json:"position_tied_order_count"`
	Sign                   int    `json:"sign"`
	Position               string `json:"position"`
	AvgEntryPrice          string `json:"avg_entry_price"`
	PositionValue          string `json:"position_value"`
	UnrealizedPnl          string `json:"unrealized_pnl"`
	RealizedPnl            string `json:"realized_pnl"`
	LiquidationPrice       string `json:"liquidation_price"`
	MarginMode             int    `json:"margin_mode"`
	AllocatedMargin        string `json:"allocated_margin"`
}

type Account struct {
	Code                     int           `json:"code"`
	AccountType              int           `json:"account_type"`
	Index                    int           `json:"index"`
	L1Address                string        `json:"l1_address"`
	CancelAllTime            int           `json:"cancel_all_time"`
	TotalOrderCount          int           `json:"total_order_count"`
	TotalIsolatedOrderCount  int           `json:"total_isolated_order_count"`
	PendingOrderCount        int           `json:"pending_order_count"`
	AvailableBalance         string        `json:"available_balance"`
	Status                   int           `json:"status"`
	Collateral               string        `json:"collateral"`
	AccountIndex             int           `json:"account_index"`
	Name                     string        `json:"name"`
	Description              string        `json:"description"`
	CanInvite                bool          `json:"can_invite"`
	ReferralPointsPercentage string        `json:"referral_points_percentage"`
	Positions                []Position    `json:"positions"`
	TotalAssetValue          string        `json:"total_asset_value"`
	CrossAssetValue          string        `json:"cross_asset_value"`
	Shares                   []interface{} `json:"shares"`
}

type GetAccountResult struct {
	ResultCode
	Total    int       `json:"total"`
	Accounts []Account `json:"accounts"`
}

type SubAccount struct {
	Code                    int    `json:"code"`
	AccountType             int    `json:"account_type"`
	Index                   int    `json:"index"`
	L1Address               string `json:"l1_address"`
	CancelAllTime           int    `json:"cancel_all_time"`
	TotalOrderCount         int    `json:"total_order_count"`
	TotalIsolatedOrderCount int    `json:"total_isolated_order_count"`
	PendingOrderCount       int    `json:"pending_order_count"`
	AvailableBalance        string `json:"available_balance"`
	Status                  int    `json:"status"`
	Collateral              string `json:"collateral"`
}

type GetAccountsByL1AddressResult struct {
	ResultCode
	L1Address   string       `json:"l1_address"`
	SubAccounts []SubAccount `json:"sub_accounts"`
}

type GetAccountLimitsResult struct {
	ResultCode
	MaxLlpPercentage    int    `json:"max_llp_percentage"`
	UserTier            string `json:"user_tier"`
	CanCreatePublicPool bool   `json:"can_create_public_pool"`
}

type OrderDetails struct {
	OrderIndex          int64  `json:"order_index"`
	ClientOrderIndex    int    `json:"client_order_index"`
	OrderId             string `json:"order_id"`
	ClientOrderId       string `json:"client_order_id"`
	MarketIndex         int    `json:"market_index"`
	OwnerAccountIndex   int    `json:"owner_account_index"`
	InitialBaseAmount   string `json:"initial_base_amount"`
	Price               string `json:"price"`
	Nonce               int64  `json:"nonce"`
	RemainingBaseAmount string `json:"remaining_base_amount"`
	IsAsk               bool   `json:"is_ask"`
	BaseSize            int    `json:"base_size"`
	BasePrice           int    `json:"base_price"`
	FilledBaseAmount    string `json:"filled_base_amount"`
	FilledQuoteAmount   string `json:"filled_quote_amount"`
	Side                string `json:"side"`
	Type                string `json:"type"`
	TimeInForce         string `json:"time_in_force"`
	ReduceOnly          bool   `json:"reduce_only"`
	TriggerPrice        string `json:"trigger_price"`
	OrderExpiry         int64  `json:"order_expiry"`
	Status              string `json:"status"`
	TriggerStatus       string `json:"trigger_status"`
	TriggerTime         int    `json:"trigger_time"`
	ParentOrderIndex    int    `json:"parent_order_index"`
	ParentOrderId       string `json:"parent_order_id"`
	ToTriggerOrderId0   string `json:"to_trigger_order_id_0"`
	ToTriggerOrderId1   string `json:"to_trigger_order_id_1"`
	ToCancelOrderId0    string `json:"to_cancel_order_id_0"`
	BlockHeight         int    `json:"block_height"`
	Timestamp           int    `json:"timestamp"`
	CreatedAt           int    `json:"created_at"`
	UpdatedAt           int    `json:"updated_at"`
}

type GetAccountActiveOrdersResult struct {
	ResultCode
	Orders []OrderDetails `json:"orders"`
}

type GetAccountInactiveOrdersResult struct {
	ResultCode
	NextCursor string         `json:"next_cursor"`
	Orders     []OrderDetails `json:"orders"`
}

type OrderBookDetails struct {
	Symbol                       string  `json:"symbol"`
	MarketId                     int     `json:"market_id"`
	Status                       string  `json:"status"`
	TakerFee                     string  `json:"taker_fee"`
	MakerFee                     string  `json:"maker_fee"`
	LiquidationFee               string  `json:"liquidation_fee"`
	MinBaseAmount                string  `json:"min_base_amount"`
	MinQuoteAmount               string  `json:"min_quote_amount"`
	OrderQuoteLimit              string  `json:"order_quote_limit"`
	SupportedSizeDecimals        int     `json:"supported_size_decimals"`
	SupportedPriceDecimals       int     `json:"supported_price_decimals"`
	SupportedQuoteDecimals       int     `json:"supported_quote_decimals"`
	SizeDecimals                 int     `json:"size_decimals"`
	PriceDecimals                int     `json:"price_decimals"`
	QuoteMultiplier              int     `json:"quote_multiplier"`
	DefaultInitialMarginFraction int     `json:"default_initial_margin_fraction"`
	MinInitialMarginFraction     int     `json:"min_initial_margin_fraction"`
	MaintenanceMarginFraction    int     `json:"maintenance_margin_fraction"`
	CloseoutMarginFraction       int     `json:"closeout_margin_fraction"`
	LastTradePrice               float64 `json:"last_trade_price"`
	DailyTradesCount             int     `json:"daily_trades_count"`
	DailyBaseTokenVolume         float64 `json:"daily_base_token_volume"`
	DailyQuoteTokenVolume        float64 `json:"daily_quote_token_volume"`
	DailyPriceLow                float64 `json:"daily_price_low"`
	DailyPriceHigh               float64 `json:"daily_price_high"`
	DailyPriceChange             float64 `json:"daily_price_change"`
	OpenInterest                 float64 `json:"open_interest"`
	DailyChart                   struct {
	} `json:"daily_chart"`
	MarketConfig struct {
		MarketMarginMode          int   `json:"market_margin_mode"`
		InsuranceFundAccountIndex int64 `json:"insurance_fund_account_index"`
	} `json:"market_config"`
}

type GetOrderBookDetailResult struct {
	ResultCode
	OrderBookDetails []OrderBookDetails `json:"order_book_details"`
}
