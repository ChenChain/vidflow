package facade

import "github.com/gin-gonic/gin"

type Hello struct {
}

func NewHello() *Hello {
	return &Hello{}
}

func (h *Hello) Healthy(c *gin.Context) {
}
