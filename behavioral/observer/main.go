package main

import (
	"design_patterns/behavioral/observer/observer"
	"design_patterns/behavioral/observer/subject"
)

func main() {
	item := subject.NewItem("Nike Shirt")
	ob1 := observer.NewCustomer("abc@gmail.com")
	ob2 := observer.NewCustomer("xyz@gmail.com")
	item.Register(ob1)
	item.Register(ob2)
	item.UpdateAvailability()
}
