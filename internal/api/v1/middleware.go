package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func disableCacheInDevMode3(next http.Handler) http.Handler {
	if !dev {
		return next
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}

var dev = true

func DisableCacheInDevMode(c *gin.Context) {
	if dev {

		c.Header("Cache-Control", "no-store")
	}
	c.Next()

}
