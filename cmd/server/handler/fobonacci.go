package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vincentconace/app-fibonacci/internal/fibonacci"
	"github.com/vincentconace/app-fibonacci/pkg/web"
)

type Fibonacci struct {
	service fibonacci.Service
}

func NewHandler(service fibonacci.Service) *Fibonacci {
	return &Fibonacci{
		service: service,
	}
}

type Request struct {
	Init  int `json:"init"`
	Limit int `json:"limit"`
	Index int `json:"index"`
}

func (f *Fibonacci) Fibonacci() gin.HandlerFunc {
	return func(c *gin.Context) {
		var fRequest Request

		err := c.ShouldBindJSON(&fRequest)
		if err != nil {
			web.Error(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		fNumber, err := f.service.CalculateFibonacci(fRequest.Init, fRequest.Limit, fRequest.Index)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		web.Success(c, http.StatusOK, gin.H{"The number in the index is": fNumber})

	}
}
