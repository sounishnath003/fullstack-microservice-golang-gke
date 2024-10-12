package handlers

type UserSignupDto struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`

	Username string `json:"username"`
	Password string `json:"password"`
}
