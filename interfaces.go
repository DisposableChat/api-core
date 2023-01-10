package core

type JSONPayload struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Exit    int8        `json:"exited_code"`
}

type Session struct {
	ID         string `json:"id"`         // user id
	Expiration int64  `json:"expiration"` // maximum 7 days
}

// User is abstraction of what user is really, we don't store users, users are stored in cache as temporary accounts
// This for offer the maximum privacy possible
type User struct {
	ID         string `json:"id"`
	PublicName string `json:"public_name"`
	Groups    []string `json:"groups"`
	Expiration int64  `json:"expiration"` // maximum 7 days, this is used to create the token used in every request
}

type GetSessionRequest struct {
	Token string `json:"token"`
	HashPassword string `json:"hash_password"`
}

type CreateSessionRequest struct {
	PublicName string `json:"public_name"`
	Expiration int64  `json:"expiration"` // maximum 7 days, in hours
	HashPassword string `json:"hash_password"` // for hash session, and only user is capable of unhash it
}

const (
	Created      = "created"
	Authorized   = "authorized"
	Unauthorized = "unauthorized"
)

const (
	BodyError              = "invalid_json_payload"
	JSONEmptyValuesError   = "invalid_json_payload"
	InvalidQueryError      = "invalid_query"
	UserNotFoundError      = "user_not_found"
	SessionError           = "session_error"
	InvalidTokenError      = "invalid_token"
	InternalServerError    = "internal_server_error"
	NameLengthError        = "name_length_error"
	ExpirationSessionError = "expiration_must_be_max_7_days_or_min_15_minutes"

	DBInsertionUser = "db_insertion_user"
	ExpiredTokenError = "expired_token"
)