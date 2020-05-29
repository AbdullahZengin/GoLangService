package app

import (
	u "../utils"
	"context"
	"fmt"
	"net/http"
)

var LoginMiddleWare = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		notAuth := []string{"/", "/api/user/login"} //List of endpoints that doesn't require auth
		requestPath := r.URL.Path                   //current request path

		//check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range notAuth {

			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization") //get the token from the header

		if tokenHeader == "" { //Token is missing, returns with error code 403 Unauthorized
			response = u.Message(false, "Oturum bilgileriniz boş olamaz.")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}
		if tokenHeader != "Gizem-Abdullah-SoftwareLab" {
			response = u.Message(false, "Oturum bilgisini yanlış gönderildi.")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return

		} else {
			fmt.Sprintf("User %", "Gizem-Abdullah") //Useful for monitoring
			ctx := context.WithValue(r.Context(), "user", "Gizem-Abdullah")
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r) //proceed in the middleware chain!
		}
	})
}
