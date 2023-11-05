package main

import (
	"fmt"
	"os"
	"regexp"
)

type RestaurantApp struct {
	restaurant    *Restaurant
	userProfile   *UserProfile
	paymentMethod PaymentMethod
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
func isUpper(s string) bool {
    if len(s) == 0 {
        return false
    }
    return 'A' <= s[0] && s[0] <= 'Z'
}
func isAlphabetic(char rune) bool {
    return ('A' <= char && char <= 'Z') || ('a' <= char && char <= 'z')
}


func isValidName(name string) bool {
    if len(name) < 3 || !isUpper(name) {
        return false
    }

    for _, char := range name {
        if !isAlphabetic(char) {
            return false
        }
    }

    return true
}

func isValidSurname (surname string) bool {
    if len(surname) < 3 || !isUpper(surname) {
        return false
    }

    for _, char := range surname {
        if !isAlphabetic(char) {
            return false
        }
    }

    return true
}

func isValidEmail(email string) bool {
	emailPattern := `^[a-zA-Z0-9!#$%&'*+\-/=?^_{}|~]+(\.[a-zA-Z0-9!#$%&'*+\-/=?^_{}|~]+)*@([a-zA-Z0-9](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`
    return regexp.MustCompile(emailPattern).MatchString(email)
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

	fmt.Println("Choose date (dd.mm.yyyy):")
	fmt.Scan(&date)
	fmt.Println("Choose time:")
	fmt.Scan(&time)
	fmt.Println("Party size:")
	_, err := fmt.Scan(&partySize)
	if err != nil || partySize <= 0 {
		fmt.Println("Invalid party size:", err)
		return
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
	_, err = fmt.Scan(&choice)
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

	fmt.Println("Enter card number:")
	fmt.Scan(&cardNumber)
	fmt.Println("Enter CVV:")
	fmt.Scan(&cvv)
	fmt.Println("Amount of money in card:")
	_, err := fmt.Scan(&amount)
	if err != nil || amount < 0 {
		return nil, fmt.Errorf("invalid amount: %v", err)
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

	fmt.Println("Enter email:")
	fmt.Scan(&email)
	fmt.Println("Amount of money in PayPal:")
	_, err := fmt.Scan(&amount)
	if err != nil || amount < 0 {
		return nil, fmt.Errorf("invalid amount: %v", err)
	}

	payPal := NewPayPal(email, amount)
	return payPal, nil
}

/*
    if choice != 1 {
        return
    }

    fmt.Println("Choose payment method:")
    fmt.Println("1. Card    2. PayPal")
    _, err = fmt.Scan(&choice)
    if err != nil {
        fmt.Println("Invalid choice:", err)
        return
    }


    if choice == 1 {
        app.paymentMethod = NewCreditCard("1234-5678-9876-5432", "123", 1000)
    } else if choice == 2 {
        app.paymentMethod = NewPayPal("example@example.com", 1500)
    }

    if app.paymentMethod == nil {
        fmt.Println("Invalid payment method choice.")
        return
    }

    if app.restaurant.BookTable(date, time, partySize, app.userProfile, app.paymentMethod) {
        fmt.Println("Congrats! Payment was successful.")
    } else {
        fmt.Println("Sorry, not enough money.")
    }
}

*/
