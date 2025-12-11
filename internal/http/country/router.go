package httpcountry

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup, h *Handler) {
	countries := r.Group("/countries")
	{
		countries.GET("", h.ListCountries)
		countries.POST("", h.CreateCountry)
		countries.GET("/:id", h.GetCountry)
		countries.PATCH("/:id", h.UpdateCountry)
		countries.DELETE("/:id", h.DeleteCountry)
	}
}
