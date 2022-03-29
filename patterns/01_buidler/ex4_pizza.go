package main

// Step 1: Create a concrete class
type pizza struct {
	dough   string
	sauce   string
	topping string
}

type Pizza interface {
	setDough(dough string)
	setSauce(sauce string)
	setTopping(sauce string)
}

func (p *pizza) setDough(dough string) {
	p.dough = dough
}

func (p *pizza) setSauce(sauce string) {
	p.sauce = sauce
}

func (p *pizza) setTopping(topping string) {
	p.topping = topping
}

type PizzaBuilder interface {
	HawaiPizza(Color) PizzaBuilder
	SpicyPizza(Wheels) PizzaBuilder
	Build() Interface
}
