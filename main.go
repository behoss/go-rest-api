package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	myBill := createBill("Behzad")
	myBill.addItem("cake", 2.49)
	myBill.addItem("latte", 1.99)
	myBill.addTip(0.99)

	fmt.Fprintf(w, myBill.format())
}

func main() {
	port := 3001
	fmt.Printf("Server listening on http://localhost:%v \n", port)

	http.HandleFunc("/", index)
	err := http.ListenAndServe((fmt.Sprintf(":%v", port)), nil)
	if err != nil {
		panic(err)
	}
}
