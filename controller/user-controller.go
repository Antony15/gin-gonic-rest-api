// Controller package for controlling data flow
package controller

import (
	"github.com/Antony15/gin-gonic-rest-api/entity"
	"github.com/Antony15/gin-gonic-rest-api/service"
	"github.com/gin-gonic/gin"
)

// UserController interface
type UserController interface {
	ViewBookings(entity.User) []entity.Bookings
	GetCabs(ctx *gin.Context) ([]entity.CabsAvailable, error)
	UserBooking(ctx *gin.Context) (entity.BookedCab, error)
}

// controller struct
type controller struct {
	service service.UserService
}

// Constructor Function to implement interface
func New(service service.UserService) UserController {
	return &controller{
		service: service,
	}
}

// Method for viewing user's bookings
func (c *controller) ViewBookings(user entity.User) []entity.Bookings {
	return c.service.ViewBookings(user)
}

// Method for getting cabs nearby user's postion
func (c *controller) GetCabs(ctx *gin.Context) ([]entity.CabsAvailable, error) {
	var (
		arrCabs []entity.CabsAvailable
		cabs    entity.Cabs
	)
	err := ctx.ShouldBindJSON(&cabs)
	if err != nil {
		return arrCabs, err
	}
	return c.service.GetCabs(cabs)
}

// Method for new booking of cab
func (c *controller) UserBooking(ctx *gin.Context) (entity.BookedCab, error) {
	var (
		booked   entity.BookedCab
		nBooking entity.NewBooking
	)
	err := ctx.ShouldBindJSON(&nBooking)
	if err != nil {
		return booked, err
	}
	return c.service.UserBooking(nBooking)
}
