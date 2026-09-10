package middleware

import (
	"net/http"
	"strings"

	"github.com/Dew-F/v-chat/internal/auth"
)

func Auth(
	jwtManager *auth.JWTManager,
) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(
			func(
				w http.ResponseWriter,
				r *http.Request,
			) {

				header :=
					r.Header.Get("Authorization")

				token :=
					strings.TrimPrefix(
						header,
						"Bearer ",
					)

				if token == "" {

					http.Error(
						w,
						"unauthorized",
						http.StatusUnauthorized,
					)

					return
				}

				_, err := jwtManager.Verify(token)

				if err != nil {

					http.Error(
						w,
						"unauthorized",
						http.StatusUnauthorized,
					)

					return
				}

				next.ServeHTTP(w, r)
			},
		)
	}
}
