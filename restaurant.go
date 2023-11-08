package main

import "fmt"

type Restaurant struct {
	Name           string
	Manager        string
	AvailableSeats int
	Observers      []Observer
}

type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}

type Observer interface {
	Update()
}

func NewRestaurant(name, manager string, availableSeats int) *Restaurant {
	restaurant := &Restaurant{
		Name:           name,
		Manager:        manager,
		AvailableSeats: availableSeats,
	}
	restaurant.Observers = make([]Observer, 0)
	return restaurant
}

func (r *Restaurant) RegisterObserver(observer Observer) {
	r.Observers = append(r.Observers, observer)
}

func (r *Restaurant) RemoveObserver(observer Observer) {
	for i, obs := range r.Observers {
		if obs == observer {
			r.Observers = append(r.Observers[:i], r.Observers[i+1:]...)
			break
		}
	}
}

func (r *Restaurant) NotifyObservers() {
	for _, observer := range r.Observers {
		observer.Update()
	}
}

func (r *Restaurant) GetInfo() {
	fmt.Printf("Restaurant name: %s\n", r.Name)
	fmt.Printf("Restaurant manager: %s\n", r.Manager)
	fmt.Printf("Number of seats: %d\n", r.AvailableSeats)
}
