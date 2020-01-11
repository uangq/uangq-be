package user

type User struct {
	Id           int64    `json:"id"`
	Username     string `json:"username"`
	PasswordHash []byte `json:"-"`
}
