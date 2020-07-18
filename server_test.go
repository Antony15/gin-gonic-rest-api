package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ViewBookings testing function
// A user should be able to view his past bookings
func TestViewBookings(t *testing.T) {
	testRouter := SetupRouter()

	req, err := http.NewRequest("GET", "/getBookings/1", nil)
	req.SetBasicAuth("admin", "admin123")
	if err != nil {
		log.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 200)
}

// GetCabs testing function
// A user should be able to get cabs that are nearby
func TestGetCabs(t *testing.T) {
	testRouter := SetupRouter()
	var jsonStr = []byte(`{"userid":1,"latitude":11.0797,"longitude":76.9997}`)
	req, err := http.NewRequest("POST", "/getCabs", bytes.NewBuffer(jsonStr))
	req.SetBasicAuth("admin", "admin123")
	if err != nil {
		log.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 200)
}

// NewBooking testing function
// A user should be able to request a booking for a cab from pickup location A to pickup location B
func TestNewBooking(t *testing.T) {
	testRouter := SetupRouter()
	var jsonStr = []byte(`{"userid":1,"pickuplat":11.0797,"pickuplong":76.9997,"destinationlat":11.0797,"destinationlong":76.9997
	}`)
	req, err := http.NewRequest("POST", "/newBooking", bytes.NewBuffer(jsonStr))
	req.SetBasicAuth("admin", "admin123")
	if err != nil {
		log.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 200)
}
