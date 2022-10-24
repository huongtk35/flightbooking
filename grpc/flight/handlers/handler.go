package handlers

import (
	"errors"
	"flightbooking/grpc/flight/repositories"
	"flightbooking/util"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	v "github.com/go-playground/validator/v10"
)

type flightHandler struct {
	validator     *v.Validate
	flightService repositories.IFlightService
	mu            *sync.Mutex
}

func NewFlightHandler(flightService repositories.IFlightService) (*flightHandler, error) {
	return &flightHandler{
		validator:     v.New(),
		flightService: flightService,
		mu:            &sync.Mutex{},
	}, nil
}

func (h *flightHandler) HandleHealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (h *flightHandler) HandleGetList(ctx *gin.Context) {
	list, err := h.flightService.GetList()
	if err != nil {
		log.Println("Error occurred", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, list)
}

func (h *flightHandler) HandleGetByCode(ctx *gin.Context) {
	code := ctx.Param("code")
	if len(code) < 1 {
		log.Println("no username path code provided")
		util.SendValidationError(ctx, errors.New("no username path code provided"))
		return
	}
	result, err := h.flightService.Get(code)
	if err != nil {
		log.Println("Error occurred", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}
