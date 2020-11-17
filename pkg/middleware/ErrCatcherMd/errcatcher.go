package ErrCatcherMd

import "net/http"

func ErrCatcher(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Some error"))
			}
		}()

		handler.ServeHTTP(w, r)
	})
}
