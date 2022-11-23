package auth

type Credentials struct {
	Email 	 string `form:"email"`
	Password string `form:"password"`
}