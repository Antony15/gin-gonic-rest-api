package main

import (
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/Antony15/wednesday-go-test/controller"
	"github.com/Antony15/wednesday-go-test/entity"
	"github.com/Antony15/wednesday-go-test/middlewares"
	"github.com/Antony15/wednesday-go-test/service"
	"github.com/gin-gonic/gin"
)

// Declare global variables
var (
	user           entity.User
	cabs           entity.Cabs
	userService    service.UserService       = service.New()
	userController controller.UserController = controller.New(userService)
)

// Function for logging in file
func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

// Function for viewing user's bookings
func ViewBookings(ctx *gin.Context) {
	user.UserID, _ = strconv.Atoi(ctx.Param("userid"))
	ctx.JSON(200, userController.ViewBookings(user))
}

// Function for getting cabs nearby user's postion
func GetCabs(ctx *gin.Context) {
	cabs, err := userController.GetCabs(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, cabs)
	}
}

// Function for new booking of cab
func NewBooking(ctx *gin.Context) {
	booked, err := userController.UserBooking(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, booked)
	}
}

// Function for setting routes
func SetupRouter() *gin.Engine {
	setupLogOutput()
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())
	server.GET("/getBookings/:userid", ViewBookings)
	server.POST("/getCabs", GetCabs)
	server.POST("/newBooking", NewBooking)
	return server
}

// Main Function
func main() {
	router := SetupRouter()
	router.Run(":8090")
}
