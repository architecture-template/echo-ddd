package parameter

type UserKey struct {
	UserKey string `json:"user_key"`
}

type RegisterUser struct {
	UserName  string  `json:"user_name"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
}
