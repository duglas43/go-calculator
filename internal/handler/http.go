package handler

import (
	"net/http"

	"github.com/duglas43/calculator/internal/executor"
	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	Exec *executor.Executor
}

func NewHTTPHandler(exec *executor.Executor) *HTTPHandler {
	return &HTTPHandler{Exec: exec}
}

// @Summary Execute instructions
// @Accept json
// @Produce json
// @Param request body executor.ExecuteRequest true "Instruction list"
// @Success 200 {object} executor.Result
// @Router /execute [post]
func (h *HTTPHandler) HandleExecute(c *gin.Context) {
	var req executor.ExecuteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := h.Exec.Execute(req.Instructions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
