package main

import "fmt"

type Restaurant struct {
	Name           string
	Manager        string
	AvailableSeats int
	Subscribers    int
}

func NewRestaurant(name, manager string, availableSeats, subscribers int) *Restaurant {
	return &Restaurant{
		Name:           name,
		Manager:        manager,
		AvailableSeats: availableSeats,
		Subscribers:    subscribers,
	}
}

func (r *Restaurant) GetInfo() {
	fmt.Printf("Restaurant name: %s\n", r.Name)
	fmt.Printf("Restaurant manager: %s\n", r.Manager)
	fmt.Printf("Number of restaurant subscribers: %d\n", r.Subscribers)
}

func (r *Restaurant) BookTable(date, time string, partySize int, userProfile *UserProfile, paymentMethod PaymentMethod) bool {
	// Check available seats
	if r.AvailableSeats < partySize {
		fmt.Println("Sorry, we don't have available seats. Please come back to your profile.")
		return false
	}

	// Calculate the reservation price
	price := partySize * 100

	// Make a reservation
	if userProfile.Pay(price) {
		r.AvailableSeats -= partySize
		fmt.Printf("Reservation successful! The price is $%d.\n", price)
		return true
	}

	fmt.Println("Payment failed. Please come back to your profile.")
	return false
}

/*package main

import "sync"

// RestaurantManager implements the Singleton pattern
type RestaurantManager struct {
	Name string
}

var restaurantManager *RestaurantManager
var once sync.Once

func GetRestaurantManager() *RestaurantManager {
	once.Do(func() {
		restaurantManager = &RestaurantManager{Name: "GOLANG"}
	})
	return restaurantManager
}


func (r *RestaurantManager) CheckAvailableSeats(date, time string, partySize int) (bool, float64) {
	// Implement seat availability and price calculation logic
	// Return true if seats are available, and the price based on party size
	return true, calculatePrice(partySize)
}

func calculatePrice(partySize int) float64 {
	// Calculate the reservation price based on party size and other factors
	// Return the price
	return 50.0 * float64(partySize)
}


func (r *RestaurantManager) BookTable(date string, partySize int, price float64) Reservation {
	// Implement table booking logic and return the reservation
	reservationFactory := &ReservationFactory{}
	reservation := reservationFactory.CreateReservation(date, partySize, price)
	return reservation
}
*/
