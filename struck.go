package main

import (
	"errors"
	"fmt"
)

// type Laptop struct {
// 	name string
// 	ram int
// 	harddisk int
// }

// func main() {
// 	l := Laptop(name :"ASUS", ram : 32, harddisk : 1000)
// fmt.Println(l)
// }

// package main

// import "fmt"

// type Laptop struct {
// 	merk string
// 	hdd  int
// 	ram  int
// }

// func (laptop *Laptop) changeMerk(newMerk string) {
// 	laptop.merk = newMerk
// }

// func main() {
// 	l := Laptop{"ASUS", 512, 2048}
// 	l.changeMerk("ACER")
// 	fmt.Println(l)
// }

// package main

// import "fmt"

// type Names struct {
// 	Fullname string
// 	Lastname string
// }

// type Person struct {
// 	Names string
// 	Age   int
// }

// func (p Person) Bicara() string {
// 	return "Hello"
// }

// func (p Person) Langkah() string {
// 	return "Berjalan"
// }

// func (p Person) sayHello(helo string) string {
// 	return p.Names + helo
// }

// func main() {
// 	p := Person{Names: "Hadyd", Age: 12}
// 	//	p.Names = Names{Fullname: "Al", Lastname: "Hadyd"}

// 	fmt.Println(p)
// 	fmt.Println(p.Bicara())
// 	fmt.Println(p.Langkah())
// 	fmt.Println(p.sayHello("Hello"))
// }

// package main

// import "fmt"

type Person struct {
	Names string
	Age   int
}

func (p *Person) ChangeName(names string) (bool, error) {
	if names == "" {
		return false, errors.New("Error tidak bisa ganti karena kosong")
	}
	p.Names = names
	return true, nil
}

// func main() {
// 	p := Person{Names: "Hadyd", Age: 22}

// 	ok, err := p.ChangeName("Hadyd")
// 	if ok {
// 		fmt.Println("Ganti nama Berhasil")
// 	} else {
// 		fmt.Println(err)
// 	}

// 	fmt.Println(p)
// }

type Anak struct {
	Person
	School string
}

func (a *Anak) setSchool(name string) error {
	if name == "" {
		return errors.New("Error name must be set")
	}
	a.School = name
	return nil
}

func main() {
	a := Anak{}

	fmt.Println(a)
	err := a.setSchool("SMP N 1 Jakarta")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(a)

	ok, err := a.ChangeName("Angga")
	if !ok {
		fmt.Println(err)
	}
	fmt.Println(a)
}
