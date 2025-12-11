package httpcountry

import (
	domain "github.com/codesayhi/golang-clean/internal/domain/country"
	usecase "github.com/codesayhi/golang-clean/internal/service/country"
	"github.com/codesayhi/golang-clean/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	svc usecase.Service
}

func NewHandler(svc usecase.Service) *Handler {
	return &Handler{svc: svc}
}

// POST /countries
func (h *Handler) CreateCountry(c *gin.Context) {
	var req CreateCountryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseValidationError(c, err.Error())
		return
	}

	input := usecase.CreateCountryInput{
		Name:     req.Name,
		Slug:     req.Slug,
		Code:     req.Code,
		Position: req.Position,
	}

	country, err := h.svc.Create(c.Request.Context(), input)
	if err != nil {
		// Có thể custom thêm logic map lỗi trùng slug/code...
		utils.ResponseServerError(c)
		return
	}

	utils.ResponseCreated(c, toCountryResponse(country))
}

// GET /countries
func (h *Handler) ListCountries(c *gin.Context) {
	var req ListCountriesRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.ResponseValidationError(c, err.Error())
		return
	}

	input := usecase.ListCountriesInput{
		Search:  req.Search,
		Page:    req.Page,
		PerPage: req.PerPage,
	}

	out, err := h.svc.List(c.Request.Context(), input)
	if err != nil {
		utils.ResponseServerError(c)
		return
	}

	resp := toListCountriesResponse(out.Items, out.Total)
	utils.ResponseOK(c, resp)
}

// GET /countries/:id
func (h *Handler) GetCountry(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		utils.ResponseBadRequest(c, "ID không hợp lệ")
		return
	}

	country, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		if err == domain.ErrNotFound {
			utils.ResponseNotFound(c, "Country không tồn tại")
			return
		}
		utils.ResponseServerError(c)
		return
	}

	utils.ResponseOK(c, toCountryResponse(country))
}

// PATCH /countries/:id
func (h *Handler) UpdateCountry(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		utils.ResponseBadRequest(c, "ID không hợp lệ")
		return
	}

	var req UpdateCountryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseValidationError(c, err.Error())
		return
	}

	input := usecase.UpdateCountryInput{
		ID:       id,
		Name:     req.Name,
		Slug:     req.Slug,
		Code:     req.Code,
		Position: req.Position,
	}

	country, err := h.svc.Update(c.Request.Context(), input)
	if err != nil {
		if err == domain.ErrNotFound {
			utils.ResponseNotFound(c, "Country không tồn tại")
			return
		}
		utils.ResponseServerError(c)
		return
	}

	utils.ResponseOK(c, toCountryResponse(country))
}

// DELETE /countries/:id
func (h *Handler) DeleteCountry(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		utils.ResponseBadRequest(c, "ID không hợp lệ")
		return
	}

	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		if err == domain.ErrNotFound {
			utils.ResponseNotFound(c, "Country không tồn tại")
			return
		}
		utils.ResponseServerError(c)
		return
	}

	utils.ResponseDeleted(c)
}
