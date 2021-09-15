package main

import "fmt"

type bill struct {
	name  string
	items map[string]float32
	tip   float32
}

func createBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float32{},
		tip:   0,
	}
	return b
}

func (b *bill) format() string {
	titleFs := fmt.Sprintf("%v's bill \n", b.name)

	var total float32 = b.tip
	itemsFs := ""
	for k, v := range b.items {
		itemsFs += fmt.Sprintf("Item: %-10v ...$%0.2f \n", k, v)
		total += v
	}

	tipFs := fmt.Sprintf("Tip: %-11v ...$%0.2f \n", "", b.tip)

	totalFs := fmt.Sprintf("Total: %-9v ...$%0.2f \n", "", total)

	return fmt.Sprint(titleFs + itemsFs + tipFs + totalFs)
}

func (b *bill) addItem(name string, price float32) {
	b.items[name] = price
}

func (b *bill) addTip(value float32) {
	b.tip = value
}
