package rooter

import (
	"log"
	"net/http"
	"runtime"

	"app/controller"

	"github.com/pkg/errors"
)

func recoverer(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// TODO
				log.Printf("[ERROR] %s\n", err)
				for depth := 0; ; depth++ {
					_, file, line, ok := runtime.Caller(depth)
					if !ok {
						break
					}
					log.Printf("======> %d: %v:%d", depth, file, line)
				}

				controller.RespondError(w, errors.New("panic"))
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(f)
}
