package handler

import (
	"net/http"
	"strconv"
	"vk-server-task/internal/models"
	"vk-server-task/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) Create(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		logrus.Errorf("user_id is not found in context")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user_id not found"})
		return
	}

	userId, ok := id.(int)
	if !ok {
		logrus.Errorf("user_id is of invalid type")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user id is of invalid type"})
		return
	}

	var input service.CreateRequest
	if err := c.BindJSON(&input); err != nil {
		logrus.Errorf("invalid create ad request: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	resp, err := h.service.Create(c.Request.Context(), userId, &input)
	if err != nil {
		logrus.Errorf("failed to create ad: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create ad"})
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"created ad": resp,
	})
}

func (h *Handler) Get(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		id = 0
	}

	userId, ok := id.(int)
	if !ok {
		logrus.Errorf("user_id is of invalid type")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user id is of invalid type"})
		return
	}

	page := getQueryInt(c, "page", 1)
	orderField := c.DefaultQuery("order_field", "created_at")
	order := c.DefaultQuery("order", "ASC")

	var minP, maxP *float64
	if v := c.Query("min_price"); v != "" {
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			minP = &f
		}
	}
	if v := c.Query("max_price"); v != "" {
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			maxP = &f
		}
	}

	params := models.AdsParams{
		Page:       page,
		OrderField: orderField,
		Order:      order,
		MinPrice:   minP,
		MaxPrice:   maxP,
	}

	output, err := h.service.Get(c.Request.Context(), userId, params)
	if err != nil {
		logrus.Errorf("failed to get list of ads: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": output,
	})
}

func getQueryInt(c *gin.Context, key string, defaultVal int) int {
	valStr := c.Query(key)
	if valStr == "" {
		return defaultVal
	}
	val, err := strconv.Atoi(valStr)
	if err != nil || val < 0 {
		return defaultVal
	}
	return val
}
