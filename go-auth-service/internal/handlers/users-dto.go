package handlers

// SignupUserDto user data-object
type SignupUserDto struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

// LoginUserDto login data-object
type LoginUserDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type VerifyUserDto struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}
