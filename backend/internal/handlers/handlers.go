package handlers

import (
	"log"
	"net/http"
	"time"

	"rzd-sales/backend/internal/config"
	"rzd-sales/backend/internal/models"
	"rzd-sales/backend/internal/rzd"

	"github.com/gin-gonic/gin"
)

// Handler представляет HTTP обработчики
type Handler struct {
	rzdClient *rzd.Client
	config    *config.Config
}

// NewHandler создает новый экземпляр Handler
func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		rzdClient: rzd.NewClient(&cfg.RZD),
		config:    cfg,
	}
}

// SearchStations обрабатывает запрос на поиск станций
func (h *Handler) SearchStations(c *gin.Context) {
	query := c.Query("query")
	if len(query) < 2 {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Query must be at least 2 characters long",
		})
		return
	}

	stations, err := h.rzdClient.SearchStations(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to search stations",
		})
		return
	}

	c.JSON(http.StatusOK, stations)
}

// SearchTrains обрабатывает запрос на поиск поездов
func (h *Handler) SearchTrains(c *gin.Context) {
	// Получаем все параметры запроса
	queryParams := c.Request.URL.Query()
	log.Printf("All query parameters: %v", queryParams)

	// Получаем параметры напрямую из запроса
	fromCode := c.Query("fromCode")
	toCode := c.Query("toCode")
	dateStr := c.Query("date")

	// Логируем все параметры запроса
	log.Printf("Received request parameters:")
	log.Printf("fromCode: %s", fromCode)
	log.Printf("toCode: %s", toCode)
	log.Printf("date: %s", dateStr)

	// Проверяем обязательные параметры
	if fromCode == "" || toCode == "" || dateStr == "" {
		log.Printf("Missing parameters: fromCode=%v, toCode=%v, date=%v", fromCode, toCode, dateStr)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Missing required parameters: fromCode, toCode, and date are required",
		})
		return
	}

	// Parse date
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		log.Printf("Date parsing error: %v for date string: %s", err, dateStr)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid date format. Use YYYY-MM-DD",
		})
		return
	}

	// Validate date is not in the past
	if date.Before(time.Now().Truncate(24 * time.Hour)) {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Date cannot be in the past",
		})
		return
	}

	trains, err := h.rzdClient.SearchTrains(fromCode, toCode, date)
	if err != nil {
		log.Printf("Error searching trains: %v", err)
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to search trains",
		})
		return
	}

	c.JSON(http.StatusOK, models.SearchResponse{
		Trains: trains,
	})
}
