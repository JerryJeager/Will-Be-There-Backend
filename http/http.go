package http

type UserIDPathParm struct {
	UserID string `uri:"user-id" binding:"required,uuid_rfc4122"`
}

type EventIDPathParam struct {
	EventID string `uri:"event-id" binding:"required,uuid_rfc4122"`
}
type EventTypeIDPathParam struct {
	EventTypeID string `uri:"event-type-id" binding:"required,uuid_rfc4122"`
}
type InviteeIDPathParam struct {
	InviteeID string `uri:"invitee-id" binding:"required,uuid_rfc4122"`
}


