package main

import (
	"fmt"
	"os"
	"sync"
)

type RestaurantApp struct {
	restaurant    *Restaurant
	userProfile   *UserProfile
	paymentMethod PaymentMethod
}

var instance *RestaurantApp
var once sync.Once


func getInstance() *RestaurantApp {
    once.Do(func() {
        instance = NewRestaurantApp()
    })
    return instance
}

func NewRestaurantApp() *RestaurantApp {
	restaurant := NewRestaurant("GOLANG", "Jackson Wang", 100, 4567)
	return &RestaurantApp{
		restaurant:    restaurant,
		userProfile:   nil,
		paymentMethod: nil,
	}
}


func (app *RestaurantApp) Run() {
	fmt.Println("Hello! Welcome to our restaurant '", app.restaurant.Name, "'")
	for {
		fmt.Println("Choose one of the options:")
		fmt.Println("1. Make reservation    2. Get Info   3. Exit")
		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid choice:", err)
			os.Exit(1)
		}

		switch choice {
		case 1:
			app.userProfile = app.createProfile()
			app.userProfile.PrintProfile()
			app.reservationMenu()
		case 2:
			app.restaurant.GetInfo()
		case 3:
			fmt.Println("Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func (app *RestaurantApp) createProfile() *UserProfile {
	var name, surname, email string
	for {
		fmt.Println("Enter your name:")
		fmt.Scan(&name)
		if !isValidName(name) {
			fmt.Println("Invalid name. Please try again.")
		} else {
			break
		}
	}
	for {
		fmt.Println("Enter your surname:")
		fmt.Scan(&surname)
		if !isValidSurname(surname) {
			fmt.Println("Invalid surname. Please try again.")
		} else {
			break
		}
	}
	for {
		fmt.Println("Enter your email:")
		fmt.Scan(&email)
		if !isValidEmail(email) {
			fmt.Println("Invalid email. Please try again.")
		} else {
			break
		}
	}
	userProfile := NewUserProfile(name, surname, email)
	app.userProfile = userProfile
	return userProfile
}

func (app *RestaurantApp) reservationMenu() {
	for {
		fmt.Println("1. Make a reservation    2. My profile   3. Cancel")
		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid choice:", err)
			os.Exit(1)
		}

		switch choice {
		case 1:
			app.makeReservation()
		case 2:
			if app.userProfile != nil {
				app.userProfile.PrintProfile()
			} else {
				fmt.Println("You need to create a profile first.")
			}
		case 3:
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func (app *RestaurantApp) makeReservation() {
	if app.userProfile == nil {
		fmt.Println("You need to create a profile first.")
		return
	}

	var date, time string
	var partySize int

	for {
		fmt.Println("Choose date (dd.mm.yyyy):")
		fmt.Scan(&date)
		if !isValidDate(date) {
			fmt.Println("Invalid date format. Please try again.")
		} else {
			break
		}
	}

	for {
		fmt.Println("Choose time:")
		fmt.Scan(&time)
		if !isValidTime(time) {
			fmt.Println("Invalid time format. Please try again.")
		} else {
			break
		}
	}

	for {
		fmt.Println("Party size:")
		_, err := fmt.Scan(&partySize)
		if err != nil || partySize <= 0 {
			fmt.Println("Invalid party size:", err)
		} else {
			break
		}
	}

	reservationPrice := partySize * 100
	fmt.Printf("The price for the reservation is $%d.\n", reservationPrice)

	if app.restaurant.AvailableSeats < partySize {
		fmt.Println("Sorry, we don't have available seats. Please come back to your profile.")
		return
	}

	fmt.Println("Do you want to make a reservation?")
	fmt.Println("1. Yes    2. No")
	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Invalid choice:", err)
		return
	}
	if choice != 1 {
		return
	}

	paymentMethod, err := app.selectPaymentMethod()
	if err != nil {
		fmt.Println(err)
		return
	}

	if paymentMethod.Pay(reservationPrice) {
		app.restaurant.AvailableSeats -= partySize
		fmt.Println("Congrats! Reservation successful.")
	} else {
		fmt.Println("Sorry, not enough money. Please come back to your profile.")
	}
}

func (app *RestaurantApp) selectPaymentMethod() (PaymentMethod, error) {
	fmt.Println("Choose payment method:")
	fmt.Println("1. Credit Card   2. PayPal")
	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil {
		return nil, fmt.Errorf("invalid choice: %v", err)
	}

	switch choice {
	case 1:
		return app.createCreditCardPayment()
	case 2:
		return app.createPayPalPayment()
	default:
		return nil, fmt.Errorf("invalid payment method choice")
	}
}

func (app *RestaurantApp) createCreditCardPayment() (PaymentMethod, error) {
	if app.userProfile == nil {
		return nil, fmt.Errorf("you need to create a profile first")
	}

	var cardNumber, cvv string
	var amount int

	for {
		fmt.Println("Enter card number:")
		fmt.Scan(&cardNumber)
		if len(cardNumber) != 16 {
			fmt.Println("Card number should consist of 16 digits. Please try again.")
			continue
		} else {
			break
		}
	}

	for {
		fmt.Println("Enter CVV:")
		fmt.Scan(&cvv)
		if len(cvv) != 3 {
			fmt.Println("CVV should consist of 3 digits. Please try again.")
			continue
		} else {
			break
		}
	}
	for {
		fmt.Println("Amount of money in card:")
		_, err := fmt.Scan(&amount)
		if err != nil || amount < 0 {
			fmt.Println("Invalid amount:", err)
		} else {
			break
		}
	}

	creditCard := NewCreditCard(cardNumber, cvv, amount)
	return creditCard, nil
}

func (app *RestaurantApp) createPayPalPayment() (PaymentMethod, error) {
	if app.userProfile == nil {
		return nil, fmt.Errorf("you need to create a profile first")
	}

	var email string
	var amount int

	for {
		fmt.Println("Enter email:")
		fmt.Scan(&email)
		if !isValidEmail(email) {
			fmt.Println("Invalid email. Please try again.")
			continue
		} else {
			break
		}
	}

	for {
		fmt.Println("Amount of money in PayPal:")
		_, err := fmt.Scan(&amount)
		if err != nil || amount < 0 {
			fmt.Println("Invalid amount:", err)
		} else {
			break
		}
	}

	payPal := NewPayPal(email, amount)
	return payPal, nil
}
