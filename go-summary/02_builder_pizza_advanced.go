package main

import "fmt"

type PizzaBuilder interface {
	PrepareDough()
	AddSauce()
	AddToppings()
	BuildPizza() Pizza
}

func GetBuilder(builderType string) PizzaBuilder {
	if builderType == "margherita" {
		return &MargheritaBuilder{}
	}
	if builderType == "pepperoni" {
		return &PepperoniBuilder{}
	}
	return nil
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

func (m *MargheritaBuilder) WithQuality(q string) PizzaBuilder {
	m.Quality = q
	return m
}

func (m *MargheritaBuilder) GetQuality() string {
	return m.Quality
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
	pepperoniBuilder := GetBuilder("pepperoni")

	pizzaDirector := &PizzaDirector{}

	pizzaDirector.builder = margheritaBuilder.WithQuality("high quality")
	fmt.Println(margheritaBuilder.GetQuality())
	margheritaPizza := pizzaDirector.BuildPizza()

	pizzaDirector.builder = pepperoniBuilder
	pepperoniPizza := pizzaDirector.BuildPizza()

	fmt.Println(margheritaPizza)
	fmt.Println(pepperoniPizza)
}
