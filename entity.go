package main

type Member struct {
	id     int
	Nama   Name
	alamat string
	point  float32
}

type Name struct {
	first_name  string
	middle_name string
	last_name   string
}
