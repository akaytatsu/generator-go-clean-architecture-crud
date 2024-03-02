package handlers

import (
	"app/entity"
	usecase_<%= entityLowerCase %> "app/usecase/<%= entityLowerCase %>"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type <%= entityUpCase %>Handler struct {
	Usecase<%= entityUpCase %> usecase_<%= entityLowerCase %>.IUsecase<%= entityUpCase %>
}

func New<%= entityUpCase %>Handler(usecase usecase_<%= entityLowerCase %>.IUsecase<%= entityUpCase %>) *<%= entityUpCase %>Handler {
	return &<%= entityUpCase %>Handler{Usecase<%= entityUpCase %>: usecase}
}

// @Summary Get <%= entityUpCase %>
// @Description Get <%= entityUpCase %>
// @Tags <%= entityUpCase %>
// @Accept  json
// @Produce  json
// @Param year path int true "Year"
// @Success 200 {object} entity.Entity<%= entityUpCase %> "success"
// @Router /api/<%= entityLowerCase %>/{year} [get]
func (h *<%= entityUpCase %>Handler) Get<%= entityUpCase %>(c *gin.Context) {
	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		handleError(c, err)
		return
	}

	<%= entityLowerCase %>, err := h.Usecase<%= entityUpCase %>.Get(year)

	if exception := handleError(c, err); exception {
		return
	}

	c.JSON(http.StatusOK, <%= entityLowerCase %>)
}

// @Summary Get All <%= entityUpCase %>
// @Description Get All <%= entityUpCase %>
// @Tags <%= entityUpCase %>
// @Accept  json
// @Produce  json
// @Success 200 {object} []entity.Entity<%= entityUpCase %> "success"
// @Router /api/<%= entityLowerCase %> [get]
func (h *<%= entityUpCase %>Handler) GetAll<%= entityUpCase %>(c *gin.Context) {
	orderBy, sortOrder := getOrderByParams(c, "updated_at")
	pagina, tamanhoPagina := getPaginationParams(c)

	params := entity.SearchEntity<%= entityUpCase %>Params{
		OrderBy:   orderBy,
		SortOrder: sortOrder,
		Page:      pagina,
		PageSize:  tamanhoPagina,
		Q:         c.Query("q"),
		CreatedAt: c.Query("created_at"),
	}

	<%= entityLowerCase %>, totalRegs, err := h.Usecase<%= entityUpCase %>.GetAll(params)

	if exception := handleError(c, err); exception {
		return
	}

	paginationResponse := PaginationResponse{
		TotalPages:     getTotalPaginas(totalRegs, tamanhoPagina),
		Page:           pagina,
		TotalRegisters: int(totalRegs),
		Registers:      <%= entityLowerCase %>,
	}

	c.JSON(http.StatusOK, paginationResponse)
}

// @Summary Create <%= entityUpCase %>
// @Description Create <%= entityUpCase %>
// @Tags <%= entityUpCase %>
// @Accept  json
// @Produce  json
// @Param <%= entityLowerCase %> body entity.Entity<%= entityUpCase %> true "<%= entityUpCase %>"
// @Success 201 {object} entity.Entity<%= entityUpCase %> "success"
// @Router /api/<%= entityLowerCase %> [post]
func (h *<%= entityUpCase %>Handler) Create<%= entityUpCase %>(c *gin.Context) {
	var <%= entityLowerCase %> entity.Entity<%= entityUpCase %>

	if err := c.ShouldBindJSON(&<%= entityLowerCase %>); err != nil {
		handleError(c, err)
		return
	}

	err := h.Usecase<%= entityUpCase %>.Create(&<%= entityLowerCase %>)

	if exception := handleError(c, err); exception {
		return
	}

	c.JSON(http.StatusCreated, <%= entityLowerCase %>)
}

// @Summary Update <%= entityUpCase %>
// @Description Update <%= entityUpCase %>
// @Tags <%= entityUpCase %>
// @Accept  json
// @Produce  json
// @Param <%= entityLowerCase %> body entity.Entity<%= entityUpCase %> true "<%= entityUpCase %>"
// @Success 200 {object} entity.Entity<%= entityUpCase %> "success"
// @Router /api/<%= entityLowerCase %> [put]
func (h *<%= entityUpCase %>Handler) Update<%= entityUpCase %>(c *gin.Context) {
	var <%= entityLowerCase %> entity.Entity<%= entityUpCase %>

	if err := c.ShouldBindJSON(&<%= entityLowerCase %>); err != nil {
		handleError(c, err)
		return
	}

	err := h.Usecase<%= entityUpCase %>.Update(&<%= entityLowerCase %>)

	if exception := handleError(c, err); exception {
		return
	}

	c.JSON(http.StatusOK, <%= entityLowerCase %>)
}

// @Summary Delete <%= entityUpCase %>
// @Description Delete <%= entityUpCase %>
// @Tags <%= entityUpCase %>
// @Accept  json
// @Produce  json
// @Param id path int true "<%= entityUpCase %> ID"
// @Success 200 {object} entity.Entity<%= entityUpCase %> "success"
// @Router /api/<%= entityLowerCase %>/{id} [delete]
func (h *<%= entityUpCase %>Handler) Delete<%= entityUpCase %>(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, err)
		return
	}

	err = h.Usecase<%= entityUpCase %>.Delete(id)

	if exception := handleError(c, err); exception {
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "<%= entityUpCase %> deleted successfully"})
}

func Mount<%= entityUpCase %>Routes(gin *gin.Engine, conn *gorm.DB, usecase usecase_<%= entityLowerCase %>.IUsecase<%= entityUpCase %>) {
	handler := New<%= entityUpCase %>Handler(usecase)

	group := gin.Group("/api/<%= entityLowerCase %>")

	group.GET("/:year", handler.Get<%= entityUpCase %>)
	group.GET("/", handler.GetAll<%= entityUpCase %>)

	SetAdminMiddleware(conn, group)

	group.POST("/", handler.Create<%= entityUpCase %>)
	group.PUT("/:id", handler.Update<%= entityUpCase %>)
	group.DELETE("/:id", handler.Delete<%= entityUpCase %>)
}
