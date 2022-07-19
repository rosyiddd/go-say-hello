// package main
package gosayhello

// import (
// 	"errors"
// 	"fmt"
// )

func sayHello() string {
	return "Hello All"
}

// type Filter func(string) string

// type Customer struct {
// 	Name, Address string
// 	Age           int
// 	Married       bool
// }

// func (customer Customer) getName() string {
// 	return customer.Name
// }
// func Kosong(i int) interface{} {
// 	if i > 0 {
// 		return "Hai"
// 	}
// 	return i
// }

// type HasName interface {
// 	getName() string
// }

// func sayHuuu(hasName HasName) {
// 	fmt.Println("Huuuuu from", hasName.getName())
// }

// func (customer Customer) sayHi(name string) {
// 	fmt.Println("Hello", name, "My name is", customer.Name)
// }
// func Defer() {
// 	fmt.Println("Aplikasi selesai")
// 	err := recover()
// 	fmt.Println(err)
// }

// func endApp(err bool) {
// 	defer Defer()
// 	fmt.Println("Aplikasi Berjalan")
// 	if err {
// 		panic("Aplikasi error")
// 	}
// 	fmt.Println("Hello")
// }

// func sumAll(numbers ...int) int {
// 	total := 0
// 	for _, value := range numbers {
// 		total += value
// 	}
// 	return total
// }

// type Blacklist func(string) bool

// func register(name string, blacklist Blacklist) string {
// 	if blacklist(name) {
// 		return "You blocked " + name
// 	} else {
// 		return name
// 	}
// }

// func filter(name string) string {
// 	if name == "Anjing" {
// 		return "..."
// 	} else {
// 		return name
// 	}
// }

// func sayHelloWithFilter(name string, filter Filter) {
// 	nameFiltered := filter(name)
// 	fmt.Println("hello", nameFiltered)
// }

// func Pembagi(nilai int, pembagi int) (int, error) {
// 	if pembagi == 0 {
// 		return 0, errors.New("Pembagi tidak boleh nol")
// 	} else {
// 		result := nilai / pembagi
// 		return result, nil
// 	}
// }

// func main() {

// 	hasil, err := Pembagi(10, 0)
// 	if err == nil {
// 		fmt.Println(hasil)
// 	} else {
// 		fmt.Println(err.Error())
// 	}
// 	var inko = Kosong(0)
// 	fmt.Println(inko)
// 	var Joko Customer
// 	Joko.Name = "Joko Widodo"
// 	Joko.Address = "Solo"
// 	Joko.Age = 58
// 	Joko.sayHi("Eko")
// 	Bowo := Customer{
// 		Name:    "Prabowo",
// 		Address: "Kalimantan",
// 		Age:     70,
// 	}
// 	sayHuuu(Joko)
// 	fmt.Println(Joko)
// 	fmt.Println(Bowo)
// 	endApp(true)
// 	fmt.Println("Hello World")
// 	fmt.Println("Satu", 1)
// 	fmt.Println("Dua", 2)

// 	var name string = "Awan"
// 	fmt.Println(name)

// 	email := "Awan@gmail.com"
// 	fmt.Println(email)

// 	var age = 27
// 	fmt.Println(age)

// 	const (
// 		FirstName = "Uzumaki"
// 		LastName  = "Naruto"
// 	)

// 	type NoKtp string
// 	type Married bool

// 	var ktpAwan NoKtp = "11111111111111"
// 	var MarriedStatus Married = false

// 	fmt.Println(ktpAwan)
// 	fmt.Println(MarriedStatus)
// 	fmt.Println(FirstName)
// 	fmt.Println(LastName)

// 	fmt.Println("Benar", true)
// 	fmt.Println("Salah", false)
// 	fmt.Println(len("hello"))

// 	var Sayur [2]string
// 	Sayur[0] = "Terong"
// 	Sayur[1] = "Sawi"
// 	fmt.Println(Sayur)

// 	var Buah = [3]string{
// 		"Mangga",
// 		"Jeruk",
// 		"Apel",
// 	}
// 	fmt.Println(Buah)

// 	Days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
// 	daysSlice := Days[5:]
// 	daysSlice[0] = "Sabtu Lagi"
// 	daysSlice[1] = "Minggu Lagi"

// 	fmt.Println(Days)

// 	daysSlice2 := append(daysSlice, "Libur baru")
// 	daysSlice2[0] = "Ups"

// 	fmt.Println(daysSlice2)
// 	fmt.Println(Days)

// 	slice := make([]string, 2, 5)
// 	slice[0] = "Hai"
// 	slice[1] = "Hello"
// 	fmt.Println(slice)

// 	slice2 := make([]string, len(slice), cap(slice))
// 	copy(slice2, slice)
// 	fmt.Println(slice2)

// 	person := map[string]string{
// 		"name":    "Awan",
// 		"address": "Jakarta",
// 	}

// 	fmt.Println(person)

// 	person1 := make(map[string]string)
// 	person1["name"] = "Joko Widodo"
// 	person1["pekerjaan"] = "Presiden"
// 	person1["salah"] = "Ups"

// 	fmt.Println(person1)

// 	delete(person1, "salah")

// 	fmt.Println(person1)

// 	if length := len(name); length < 10 {
// 		fmt.Println("nama terlalu pendek")
// 	} else {
// 		fmt.Println("nama sudah benar")
// 	}

// 	username := "potato"

// 	switch {
// 	case username == "potato":
// 		fmt.Println("username salah")
// 	case username == "misterpotato":
// 		fmt.Println("username sudah benar")
// 	}

// 	names := []string{
// 		"Roy",
// 		"Agung",
// 	}

// 	for i, value := range names {
// 		fmt.Println(i, value)
// 	}

// 	sayHelloWithFilter("Joko", filter)
// 	fmt.Println(sumAll(1, 2, 3, 4, 5))
// 	fmt.Println(register("Penjahat", func(name string) bool {
// 		return name == "Penjahat"
// 	}))
// }
