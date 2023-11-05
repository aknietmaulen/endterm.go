// userProfile.go
package main

import "fmt"

type UserProfile struct {
	Name          string
	Surname       string
	Email         string
	PaymentMethod PaymentMethod
}

func NewUserProfile(name, surname, email string) *UserProfile {
	return &UserProfile{
		Name:    name,
		Surname: surname,
		Email:   email,
	}
}

func (u *UserProfile) PrintProfile() {
	fmt.Printf("Your name: %s\n", u.Name)
	fmt.Printf("Your surname: %s\n", u.Surname)
	fmt.Printf("Your email: %s\n", u.Email)
}

func (u *UserProfile) Pay(amount int) bool {
	if u.PaymentMethod != nil {
		return u.PaymentMethod.Pay(amount)
	}
	return false
}

/*package main

import "fmt"

type UserProfile struct {
    Name    string
    Surname string
    Email   string
}

func NewUserProfile(name, surname, email string) *UserProfile {
    return &UserProfile{
        Name:    name,
        Surname: surname,
        Email:   email,
    }
}

func (u *UserProfile) PrintProfile() {
    fmt.Printf("Your name: %s\n", u.Name)
    fmt.Printf("Your surname: %s\n", u.Surname)
    fmt.Printf("Your email: %s\n", u.Email)
}
*/
