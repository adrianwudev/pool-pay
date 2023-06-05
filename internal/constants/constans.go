package constants

const (
	FriendShipStatusPending  = "pending"
	FriendShipStatusAccepted = "accepted"
	FriendShipStatusRejected = "rejected"
)

const (
	ERRORCODE_INTERNALERROR = 500
	// friendship
	ERRORCODE_FRIENDREQUESTALREADYEXISTS = 4001
	ERRORCODE_FAILEDTOCREATEFRIENDSHIP   = 4002
	// token
	ERRORCODE_KEYNOTFOUND     = 5001
	ERRORCODE_TOKENISREQUIRED = 5002
	ERRORCODE_TOKENISINVALID  = 5003
	ERRORCODE_TOKENISEXPIRED  = 5004
	// others
	ERRORCODE_OTHERS = 7000
)

var ErrorCodeMessageMap = map[int]string{
	ERRORCODE_FRIENDREQUESTALREADYEXISTS: "friend request already exists",
	ERRORCODE_FAILEDTOCREATEFRIENDSHIP:   "failed to create friendship",
	// token
	ERRORCODE_KEYNOTFOUND:     "key not found",
	ERRORCODE_TOKENISREQUIRED: "token is required",
	ERRORCODE_TOKENISINVALID:  "token is invalid",
	ERRORCODE_TOKENISEXPIRED:  "token is expired",
}
