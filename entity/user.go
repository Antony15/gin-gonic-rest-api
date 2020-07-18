// Entity package is used for defined all structs
package entity

// User struct
type User struct {
	UserID int `json:"userid" binding:"required"`
}

// Bookings struct
type Bookings struct {
	UserId          int     `json:"userid" binding:"required"`
	Source          string  `json:"source" binding:"required"`
	Destination     string  `json:"destination" binding:"required"`
	Total_amount    float64 `json:"total_amount" binding:"required"`
	Discount_amount float64 `json:"discount_amount" binding:"required"`
	Actual_amount   float64 `json:"actual_amount" binding:"required"`
	Coupon_applied  int     `json:"coupon_applied"`
	Coupon_id       string  `json:"coupon_id"`
}

// Cabs struct
type Cabs struct {
	UserID int     `json:"userid" binding:"required"`
	Lat    float64 `json:"latitude" binding:"required"`
	Long   float64 `json:"longitude" binding:"required"`
}

// CabsAvailable struct
type CabsAvailable struct {
	CabID          int     `json:"cabid"`
	DriverName     string  `json:"drivername"`
	VehicleNo      string  `json:"vehicleno"`
	WithinDistance float64 `json:"within_distance"`
}

// NewBooking struct
type NewBooking struct {
	UserID          int     `json:"userid" binding:"required"`
	PickupLat       float64 `json:"pickuplat" binding:"required"`
	PickupLong      float64 `json:"pickuplong" binding:"required"`
	DestinationLat  float64 `json:"destinationlat" binding:"required"`
	DestinationLong float64 `json:"destinationlong" binding:"required"`
}

// BookedCab struct
type BookedCab struct {
	Message         string  `json:"message"`
	CabID           int     `json:"cabid,omitempty"`
	DriverName      string  `json:"drivername,omitempty"`
	VehicleNo       string  `json:"vehicleno,omitempty"`
	WithinDistance  float64 `json:"within_distance,omitempty"`
	PickupLat       float64 `json:"pickuplat,omitempty"`
	PickupLong      float64 `json:"pickuplong,omitempty"`
	DestinationLat  float64 `json:"destinationlat,omitempty"`
	DestinationLong float64 `json:"destinationlong,omitempty"`
	Fare            float64 `json:"fare,omitempty"`
}
