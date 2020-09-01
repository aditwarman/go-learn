package main

import "os"

func main() {
	a := App{}

	a.Initialize(
		os.Getenv("PAYMENT_DB_USERNAME"),
		os.Getenv("PAYMENT_DB_PASSWORD"),
		os.Getenv("PAYMENT_DB_NAME"))

	a.Run(":8010")
}
