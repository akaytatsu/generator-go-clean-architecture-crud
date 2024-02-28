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
// @Param id path int true "<%= entityUpCase %> ID"
// @Success 200 {object} entity.Entity<%= entityUpCase %> "success"
// @Router /api/<%= entityLowerCase %>/{id} [get]
func (h *<%= entityUpCase %>Handler) Get<%= entityUpCase %>(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	<%= entityLowerCase %>, err := h.Usecase<%= entityUpCase %>.Get<%= entityUpCase %>ByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, <%= entityLowerCase %>)
}

// @Summary Get All <%= entityUpCase %>s
// @Description Get All <%= entityUpCase %>s
// @Tags <%= entityUpCase %>
// @Accept  json
// @Produce  json
// @Success 200 {object} []entity.Entity<%= entityUpCase %> "success"
// @Router /api/<%= entityLowerCase %> [get]
func (h *<%= entityUpCase %>Handler) GetAll<%= entityUpCase %>s(c *gin.Context) {
	<%= entityLowerCase %>s, err := h.Usecase<%= entityUpCase %>.GetAll<%= entityUpCase %>s()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, <%= entityLowerCase %>s)
}

// @Summary Create <%= entityUpCase %>
// @Description Create <%= entityUpCase %>
// @Tags <%= entityUpCase %>
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param entity.Entity<%= entityUpCase %> body entity.Entity<%= entityUpCase %> true "<%= entityUpCase %>"
// @Success 200 {object} entity.Entity<%= entityUpCase %> "success"
// @Router /api/<%= entityLowerCase %> [post]
func (h *<%= entityUpCase %>Handler) Create<%= entityUpCase %>(c *gin.Context) {
	var <%= entityLowerCase %> entity.Entity<%= entityUpCase %>
	c.BindJSON(&<%= entityLowerCase %>)

	err := h.Usecase<%= entityUpCase %>.Create<%= entityUpCase %>(&<%= entityLowerCase %>)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, <%= entityLowerCase %>)
}

// @Summary Update <%= entityUpCase %>
// @Description Update <%= entityUpCase %>
// @Tags <%= entityUpCase %>
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param entity.Entity<%= entityUpCase %> body entity.Entity<%= entityUpCase %> true "<%= entityUpCase %>"
// @Success 200 {object} entity.Entity<%= entityUpCase %> "success"
// @Router /api/<%= entityLowerCase %> [put]
func (h *<%= entityUpCase %>Handler) Update<%= entityUpCase %>(c *gin.Context) {
	var <%= entityLowerCase %> entity.Entity<%= entityUpCase %>
	c.BindJSON(&<%= entityLowerCase %>)

	err := h.Usecase<%= entityUpCase %>.Update<%= entityUpCase %>(&<%= entityLowerCase %>)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.Usecase<%= entityUpCase %>.Delete<%= entityUpCase %>(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "<%= entityUpCase %> deleted successfully"})
}

func Mount<%= entityUpCase %>Handler(gin *gin.Engine, conn *gorm.DB, usecase usecase_<%= entityLowerCase %>.IUsecase<%= entityUpCase %>) {
	handler := New<%= entityUpCase %>Handler(usecase)

	group := gin.Group("/api/<%= entityLowerCase %>")

	group.GET("/:id", handler.Get<%= entityUpCase %>)
	group.GET("/", handler.GetAll<%= entityUpCase %>s)

	SetAdminMiddleware(conn, group)

	group.POST("/", handler.Create<%= entityUpCase %>)
	group.PUT("/", handler.Update<%= entityUpCase %>)
	group.DELETE("/:id", handler.Delete<%= entityUpCase %>)
}
