package core

type HTTPServerMessage struct {
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

type Message struct {
	ID      string `json:"id"`
	From    string `json:"from"`
	Content string `json:"content"`
	Length  int16  `json:"length"` // maximum 1024 characters
	Deleted bool   `json:"deleted"`
}

type Group struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Logo        string   `json:"logo"`
	Members     []string `json:"members"`
	Owner       string   `json:"owner_id"`
	CreationAt  int64    `json:"creation_at"`
	Expiration  int64    `json:"expiration"` // maximum 7 days or maximum the expiration of the user who created the group
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
	RecordedUser = "recorded_user"
)

const (
	BodyError              = "empty_body_or_invalid_json"
	JSONEmptyValuesError   = "empty_values"
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