package http_helper

import (
	"net/url"
	"uangq-be/modules/user"
)

type Request struct{}
type Response struct {
	ErrorMessage string      `json:"error_msg"`
	Data         interface{} `json:"data"`
}

type LoginRequest struct {
	Request
	Username string
	Password string
}

type LoginResponse struct {
	User  user.User `json:"user"`
	Token string    `json:"token"`
}

func (r *LoginRequest) Validate() url.Values {
	errs := url.Values{}

	if r.Username == "" {
		errs.Add("username", "Username is required.")
	}

	if r.Password == "" {
		errs.Add("password", "Password is required.")
	}

	return errs
}

type SignupRequest struct {
	Request
	Username string
	Password string
}

type SignupResponse struct {
	User  user.User `json:"user"`
	Token string    `json:"token"`
}

func (r *SignupRequest) Validate() url.Values {
	errs := url.Values{}

	if r.Username == "" {
		errs.Add("username", "Username is required.")
	}

	if r.Password == "" {
		errs.Add("password", "Password is required.")
	}

	if len(r.Password) < 4 {
		errs.Add("password", "Password length minimum is 4.")
	}

	return errs
}
