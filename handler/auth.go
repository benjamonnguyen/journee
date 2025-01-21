package handler

import (
	"net/http"
	"strings"

	"github.com/pocketbase/pocketbase/core"
)

func AuthMiddleware(e *core.RequestEvent) error {
	if e.Request.URL.Path == "/login" {
		e.Next()
	}

	cookie, _ := e.Request.Cookie("Authorization")
	if cookie != nil {
		token := strings.TrimPrefix(cookie.Value, "Bearer ")
		record, err := e.App.FindAuthRecordByToken(token, core.TokenTypeAuth)
		if err != nil {
			e.App.Logger().Debug("loadAuthToken failure", "error", err)
		} else if record != nil {
			e.Auth = record
		}
	}

	if e.Auth == nil {
		if e.Request.Header.Get("HX-Request") != "" {
			e.Response.Header().Set("HX-Redirect", "/login")
			return nil
		}

		return e.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	return e.Next()
}

func Login(e *core.RequestEvent) error {
	var form struct {
		Email    string
		Password string
	}
	if err := e.BindBody(&form); err != nil {
		return e.BadRequestError("", err)
	}

	record, _ := e.App.FindFirstRecordByData("users", "email", form.Email)
	if record == nil {
		return e.BadRequestError(
			"We couldn't find an account with that email.\nPlease try again or sign up.",
			nil,
		)
	}

	if !record.ValidatePassword(form.Password) {
		return e.UnauthorizedError("Invalid credentials.", nil)
	}

	token, _ := record.NewAuthToken()

	e.SetCookie(&http.Cookie{
		Name:  "Authorization",
		Value: token,
		// TOOD Secure: true,
		HttpOnly: true,
	})

	e.SetCookie(&http.Cookie{
		Name:  "user_email",
		Value: form.Email,
	})

	return nil
}

func Logout(e *core.RequestEvent) error {
	e.SetCookie(&http.Cookie{
		Name:  "Authorization",
		Value: "",
	})

	return e.Redirect(http.StatusTemporaryRedirect, "/login")
}
