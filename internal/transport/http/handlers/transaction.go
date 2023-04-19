package handlers

import (
	"awesomeProject/internal/logger"
	"awesomeProject/internal/models"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

// @Summary CreateTrans
// @Security ApiKeyAuth
// @Tags transaction
// @Description create transaction
// @ID create-transaction
// @Accept  json
// @Produce  json
// @Param input body models.Transaction true "transaction info"
// @Success 201 {string} string "id"
// @Failure 400,404 {string} string "message"
// @Failure 500 {string} string "message"
// @Failure default {string} string "message"
// @Router /trans/create [post]
func (h Manager) CreateTrans() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.Transaction
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, fmt.Sprintf("Bind err:%s", err))
			return
		}
		var Validate = validator.New()
		validationErr := Validate.Struct(req)
		if validationErr != nil {
			logger.Logger().Println(validationErr)
			c.JSON(http.StatusBadRequest, fmt.Sprintf("Validation err:%s", validationErr))
			return
		}
		response, err := h.srv.Transaction.Create(c.Request.Context(), &req)
		if err != nil {
			logger.Logger().Println(err)
			c.JSON(http.StatusBadRequest, fmt.Sprintf("Create err:%s", err))
			return
		}

		c.JSON(http.StatusCreated, fmt.Sprintf("transaction successfully passed!!", response))
	}
}

// @Summary GetTransByID
// @Tags transaction
// @Description get transaction by ID
// @ID get-transaction-by-id
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Transaction
// @Failure 400,404 {object} string "message"
// @Failure 500 {object} string "message"
// @Failure default {object} string "message"
// @Router /trans/get_by_id/{id} [get]
func (h Manager) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		queryParam := c.Param("id")
		get, err := h.srv.Transaction.Get(ctx, queryParam)
		if err != nil {
			logger.Logger().Println(err)
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusFound, get)
	}
}

// @Summary ListTrans
// @Tags transaction
// @Description list transactions
// @ID list-transactions
// @Accept json
// @Produce json
// @Success 200 {object} []models.Transaction
// @Failure 400,404 {object} string "message"
// @Failure 500 {object} string "message"
// @Failure default {object} string "message"
// @Router /trans/list [get]
func (h Manager) ListTrans() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := h.srv.Transaction.List(c.Request.Context())
		if err != nil {
			logger.Logger().Println(err)
			c.JSON(http.StatusBadRequest, err)
		}

		c.JSON(200, resp)
	}
}

// @Summary DeleteTrans
// @Security ApiKeyAuth
// @Tags transaction
// @Description delete transaction by id
// @ID delete-transaction
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200
// @Failure 400,404 {object} string "message"
// @Failure 500 {object} string "message"
// @Failure default {object} string "message"
// @Router /trans/delete/{id} [get]
func (h Manager) DeleteTrans() gin.HandlerFunc {
	return func(c *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		queryParam := c.Param("id")
		err := h.srv.Transaction.Delete(c.Request.Context(), queryParam)
		if err != nil {
			logger.Logger().Println(err)
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, "User is deleted")
	}
}
