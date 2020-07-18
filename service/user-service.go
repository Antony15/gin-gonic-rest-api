// Service package for servicing request data
package service

import (
	"errors"
	"log"
	"math"
	"strconv"

	"github.com/Antony15/wednesday-go-test/db"
	"github.com/Antony15/wednesday-go-test/entity"
)

// UserService interface
type UserService interface {
	ViewBookings(entity.User) []entity.Bookings
	GetCabs(entity.Cabs) ([]entity.CabsAvailable, error)
	UserBooking(entity.NewBooking) (entity.BookedCab, error)
}

// userService struct
type userService struct {
	bookings []entity.Bookings
	cabs     []entity.CabsAvailable
	booked   entity.BookedCab
}

// Constructor Function to implement interface
func New() UserService {
	return &userService{
		bookings: []entity.Bookings{},
		cabs:     []entity.CabsAvailable{},
		booked:   entity.BookedCab{},
	}
}

// Method for viewing user's bookings
func (service *userService) ViewBookings(user entity.User) []entity.Bookings {
	var eub []entity.Bookings
	rows, err := db.DbQuery("select user_id,source,destination,total_amount,discount_amount,actual_amount,coupon_applied,coupon_id from user_bookings where user_id = ?", user.UserID)
	if err != nil {
		log.Println("Mysql error:", err.Error())
		return service.bookings
	}
	for rows.Next() {
		var (
			userid          int
			source          string
			destination     string
			total_amount    float64
			discount_amount float64
			actual_amount   float64
			coupon_applied  int
			coupon_id       string
		)
		err = rows.Scan(&userid, &source, &destination, &total_amount, &discount_amount, &actual_amount, &coupon_applied, &coupon_id)
		if err != nil {
			log.Println("Mysql error:", err.Error())
		} else {
			geteub := entity.Bookings{UserId: userid, Source: source, Destination: destination, Total_amount: total_amount, Discount_amount: discount_amount, Actual_amount: actual_amount, Coupon_applied: coupon_applied, Coupon_id: coupon_id}
			eub = append(eub, geteub)
		}

	}
	return eub
}

// Method for getting cabs nearby user's postion
func (service *userService) GetCabs(cabs entity.Cabs) ([]entity.CabsAvailable, error) {
	var arrCabs []entity.CabsAvailable
	cUser := checkUser(cabs.UserID)
	if cUser {
		rows, err := db.DbQuery("select c.cab_id,c.lattitude,c.longitude,cb.driverNname,cb.vehicleNo from cabs c inner join cab_driver_details cb on c.cab_id = cb.id")
		if err != nil {
			log.Println("Mysql error:", err.Error())
			return service.cabs, err
		}
		for rows.Next() {
			var (
				cab_id     int
				lattitude  float64
				longitude  float64
				drivername string
				vehicle_no string
			)
			err = rows.Scan(&cab_id, &lattitude, &longitude, &drivername, &vehicle_no)
			if err != nil {
				log.Println("Mysql error:", err.Error())
			} else {
				meters := Distance(cabs.Lat, cabs.Long, lattitude, longitude)
				if meters <= 2000 {
					getCabs := entity.CabsAvailable{CabID: cab_id, DriverName: drivername, VehicleNo: vehicle_no, WithinDistance: math.Round((meters/1000)*100) / 100}
					arrCabs = append(arrCabs, getCabs)
				}
			}
		}
		if len(arrCabs) > 0 {
			return arrCabs, nil
		} else {
			return service.cabs, errors.New("No cabs Available")
		}
	}
	return service.cabs, errors.New("User ID is wrong")
}

// Method for new booking of cab
func (service *userService) UserBooking(nbookings entity.NewBooking) (entity.BookedCab, error) {
	var booked entity.BookedCab
	cabs := entity.Cabs{UserID: nbookings.UserID, Lat: nbookings.PickupLat, Long: nbookings.PickupLong}
	arrCabs, err := service.GetCabs(cabs)
	if len(arrCabs) > 0 && err == nil {
		var index int
		lessDistance := arrCabs[0].WithinDistance
		for k, v := range arrCabs {
			if v.WithinDistance < lessDistance {
				lessDistance = v.WithinDistance
				index = k
			}
		}
		booked.Message = "Cab Booked successfully"
		booked.CabID = arrCabs[index].CabID
		booked.DriverName = arrCabs[index].DriverName
		booked.VehicleNo = arrCabs[index].VehicleNo
		booked.WithinDistance = lessDistance
		booked.PickupLat = nbookings.PickupLat
		booked.PickupLong = nbookings.PickupLong
		booked.DestinationLat = nbookings.DestinationLat
		booked.DestinationLong = nbookings.DestinationLong
		booked.Fare = 50
		return booked, nil
	}
	booked.Message = "No Cabs Available!"
	return booked, err
}

// Function for checking user
func checkUser(userid int) bool {
	rows, err := db.DbNumRows("select id from users where id = " + strconv.Itoa(userid))
	if rows > 0 && err == nil {
		return true
	}
	return false
}

// Function to calculate distance returned is METERS
func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}

// Function for haversin(θ)
func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}
