package handler

import (
	"encoding/json"
	"net/http"
	"uangq-be/common/http_helper"
	"uangq-be/common/logger"
	"uangq-be/modules/user"
)

func HandleLogin() http.HandlerFunc {
	f := func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		// Parse json request
		loginRequest := &http_helper.LoginRequest{}
		err := json.NewDecoder(r.Body).Decode(loginRequest)
		if err != nil {
			logger.Log("Error decoding response: " + err.Error())
			http_helper.SendInvalidResponse(w)
			return
		}

		// Validate Request
		errs := loginRequest.Validate()
		if len(errs) > 0 {
			logger.Log("Bad request. Validate fail.")
			http_helper.SendInvalidResponseWithErrs(w, errs)
			return
		}

		// Handle logic
		_, u, jwt, msg := user.Login(loginRequest.Username, loginRequest.Password)
		http_helper.SendResponse(w, http.StatusOK, http_helper.Response{
			ErrorMessage: msg,
			Data:         http_helper.LoginResponse{u, jwt},
		})
	}

	return f
}

func HandleSignup() http.HandlerFunc {
	f := func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		// Parse json request
		signupRequest := &http_helper.SignupRequest{}
		err := json.NewDecoder(r.Body).Decode(signupRequest)
		if err != nil {
			logger.Log("Error decoding response: " + err.Error())
			http_helper.SendInvalidResponse(w)
			return
		}

		// Validate Request
		errs := signupRequest.Validate()
		if len(errs) > 0 {
			logger.Log("Bad request. Validate fail.")
			http_helper.SendInvalidResponseWithErrs(w, errs)
			return
		}

		// Handle logic
		_, u, jwt, msg := user.Signup(signupRequest.Username, signupRequest.Password)
		http_helper.SendResponse(w, http.StatusOK, http_helper.Response{
			ErrorMessage: msg,
			Data:         http_helper.SignupResponse{u, jwt},
		})
	}

	return f
}
