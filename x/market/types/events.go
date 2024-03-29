package types

// Slashing module event types
const (
	EventTypeCreatePool = "create_pool"
	EventTypeUpdatePool = "update_pool"

	EventTypeCreateBurnings = "create_burnings"
	EventTypeUpdateBurnings = "update_burnings"
	EventTypeBurn           = "burn"

	EventTypeCreateMember = "new_member"
	EventTypeUpdateMember = "update_member"

	EventTypeCreateDrop = "create_drop"
	EventTypeUpdateDrop = "update_drop"
	EventTypeRedeemDrop = "redeem_drop"

	EventTypeOrder = "order"

	AttributeKeyActive  = "active"
	AttributeKeyAmount  = "amount"
	AttributeKeyBalance = "balance"
	AttributeKeyDenom   = "denom"
	// Alpha-numeric ordered denom for pool pair
	AttributeKeyDenom1 = "denom_1"
	AttributeKeyDenom2 = "denom_2"
	// Sequenced denom pair to identify member
	AttributeKeyDenomA     = "denom_a"
	AttributeKeyDenomB     = "denom_b"
	AttributeKeyDenomAsk   = "denom_ask"
	AttributeKeyDenomBid   = "denom_bid"
	AttributeKeyDrops      = "drops"
	AttributeKeyLeaders    = "leaders"
	AttributeKeyLimit      = "limit"
	AttributeKeyNext       = "next"
	AttributeKeyOwner      = "owner"
	AttributeKeyOrderType  = "order_type"
	AttributeKeyPair       = "pair"
	AttributeKeyPrev       = "prev"
	AttributeKeyProduct    = "product"
	AttributeKeyRate       = "rate"
	AttributeKeyStatus     = "status"
	AttributeKeyStop       = "stop"
	AttributeKeyBeginTime  = "begin-time"
	AttributeKeyUpdateTime = "update-time"
	AttributeKeyUid        = "uid"
)
