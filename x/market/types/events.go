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
	EventTypeRedeemDrop = "redeem_drop"

	EventTypeCreateOrder  = "create_order"
	EventTypeExecuteOrder = "execute_order"
	EventTypeRemoveOrder  = "remove_order"
	EventTypeUpdateOrder  = "update_order"

	AttributeKeyActive  = "active"
	AttributeKeyAmount  = "amount"
	AttributeKeyBalance = "balance"
	AttributeKeyDenom   = "denom"
	// Alpha-numeric ordered denom for pool pair
	AttributeKeyDenom1 = "denom_1"
	AttributeKeyDenom2 = "denom_2"
	// Sequenced denom pair to identify member
	AttributeKeyDenomA   = "denom_a"
	AttributeKeyDenomB   = "denom_b"
	AttributeKeyDenomAsk = "denom_ask"
	AttributeKeyDenomBid = "denom_bid"
	AttributeKeyDrops    = "drops"
	AttributeKeyLeader   = "leader"
	AttributeKeyLimit    = "limit"
	AttributeKeyNext     = "next"
	AttributeKeyNext1    = "next1"
	AttributeKeyNext2    = "next2"
	AttributeKeyPair     = "pair"
	AttributeKeyPrev     = "prev"
	AttributeKeyPrev1    = "prev1"
	AttributeKeyPrev2    = "prev2"
	AttributeKeyRate     = "rate"
	AttributeKeyRate1    = "rate1"
	AttributeKeyRate2    = "rate2"
	AttributeKeyStop     = "stop"
	AttributeKeyUid      = "uid"
)
