package middleware

import (
	"cherry-disk/core/helper"
	"net/http"
)

type AuthMiddleWare struct {
}

func NewAuthMiddleWare() *AuthMiddleWare {
	return &AuthMiddleWare{}
}

func (m *AuthMiddleWare) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

		token, err := helper.AnalyzeToken(auth)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}

		r.Header.Set("UserId", string(rune(token.Id)))
		r.Header.Set("UserIdentity", token.Identity)
		r.Header.Set("UserName", token.Name)
		next(w, r)
	}
}
