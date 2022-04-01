package main

import "fmt"

func main() {
	// ini adalah print hellow
	print("hello world\n")

	var andre *int

	print(andre, "\n")

	var fery string

	fery = "nama manusia"

	println(fery)

	gandi := "nama yang belajar"

	println(gandi)

	fmt.Printf("Halo Gandi %s \numur saya %v \n", gandi, andre)

	// multi variable

	var agus, joko, budi string
	agus, joko, budi = "Nama orang", "nama parkir", "Orang bos"
	println(agus, budi, joko)

	nama := new(string)

	println(&nama)

	println(*nama)

	var positifThinking uint8 = 89

	println(positifThinking)

	var exists bool = true

	println(exists)

	nama = nil

	println(nama)

	andre = nil

	println("halo, ", andre)

	println(melempar("Agus"))
}

func melempar(namaorang string) *string {
	return nil
}
