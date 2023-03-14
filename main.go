package main

import "fmt"

var DBMember = [10]Member{}
var counterMember = 0

func InsertMember(anggota Member) {
	DBMember[counterMember].id = anggota.id
	DBMember[counterMember].Nama.first_name = anggota.Nama.first_name
	counterMember++
}

func ReadAllMember() {
	for i := 0; i < counterMember; i++ {
		fmt.Println("Id Member : ", DBMember[i].id)
		fmt.Println("First Name Member : ", DBMember[i].Nama.first_name)
		fmt.Println("_______________")
	}
}

func main() {
	customer1 := Member{
		id:     1,
		Nama:   Name{"kurniawan", "", ""},
		alamat: "Surabaya",
		point:  0.4,
	}

	namee := Name{
		first_name:  "muchamad",
		middle_name: "aan",
		last_name:   "kurniawan",
	}

	customer2 := Member{
		id:     2,
		Nama:   namee,
		alamat: "Malang",
		point:  4.4,
	}

	InsertMember(customer1)
	InsertMember(customer2)
	ReadAllMember()
}
