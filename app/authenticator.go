package app

import (
	"fmt"
	"net/http"
	"strings"
)

func (app *application) authenticator(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// получить токен из заголовка
		// Authorization: Bearer
		header := strings.Split(r.Header.Get("Authorization"), " ")

		if len(header) == 2 {
			uid, err := app.auth.GetUID(header[1])
			if err != nil {
				fmt.Println("Error", err.Error())
				// если не удалось - вернуть 401
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			// если удалось аутентифицировать пользователя, положить в контекст user id
			ctx = WithUserID(ctx, uid)
			fmt.Println("uid:", uid)
		} else {
			// если не удалось - вернуть 401
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
