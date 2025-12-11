package bootstrap

import (
	httpcountry "github.com/codesayhi/golang-clean/internal/http/country"
	"github.com/gin-gonic/gin"
)

func NewServer(app *Application) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	httpcountry.RegisterRoutes(api, app.CountryHandler)

	return r
}
