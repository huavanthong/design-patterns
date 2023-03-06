package main

import "fmt"

type PizzaBuilder interface {
	PrepareDough()
	AddSauce()
	AddToppings()
	BuildPizza() Pizza
}

type Pizza struct {
	Dough    string
	Sauce    string
	Toppings []string
}

type MargheritaBuilder struct {
	Pizza
	Quality string
}

func (m *MargheritaBuilder) PrepareDough() {
	m.Dough = "thin crust"
}

func (m *MargheritaBuilder) AddSauce() {
	m.Sauce = "tomato"
}

func (m *MargheritaBuilder) AddToppings() {
	m.Toppings = []string{"mozzarella cheese", "basil leaves"}
}

func (m *MargheritaBuilder) BuildPizza() Pizza {
	return m.Pizza
}

type PepperoniBuilder struct {
	Pizza
}

func (p *PepperoniBuilder) PrepareDough() {
	p.Dough = "thick crust"
}

func (p *PepperoniBuilder) AddSauce() {
	p.Sauce = "tomato"
}

func (p *PepperoniBuilder) AddToppings() {
	p.Toppings = []string{"mozzarella cheese", "pepperoni slices"}
}

func (p *PepperoniBuilder) BuildPizza() Pizza {
	return p.Pizza
}

type PizzaDirector struct {
	builder PizzaBuilder
}

func (pd *PizzaDirector) BuildPizza() Pizza {
	pd.builder.PrepareDough()
	pd.builder.AddSauce()
	pd.builder.AddToppings()
	return pd.builder.BuildPizza()
}

func main() {
	margheritaBuilder := &MargheritaBuilder{}
	pepperoniBuilder := &PepperoniBuilder{}

	pizzaDirector := &PizzaDirector{}

	pizzaDirector.builder = margheritaBuilder
	margheritaPizza := pizzaDirector.BuildPizza()

	pizzaDirector.builder = pepperoniBuilder
	pepperoniPizza := pizzaDirector.BuildPizza()

	fmt.Println(margheritaPizza)
	fmt.Println(pepperoniPizza)
}
