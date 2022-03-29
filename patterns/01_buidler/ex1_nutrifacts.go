package main

import "fmt"

// Step 1: Create Concrete class
// ===> Meaning that: from requirement, we must list all member into a structure.
type nutritionFacts struct {
	serveringSize int
	serverings    int
	calories      int
	sodium        int
	carbohydrate  int
}

// Step 2: Create methods
// ===> If you want to access to the above structure,
//		we can define all methods using interface as below
type NutritionFacts interface {
}

// Step 3: Create builder for NutrionFacts
// ===> Define t?t c? c?c methods cho vi?c kh?i t?o m?t object.
// 		Ph??ng ph?p n?y g?i l? Builder Pattern
type NutritionFactsBuilder interface {
	Serving(servingSize int, servings int) NutritionFactsBuilder
	Calories(calories int) NutritionFactsBuilder
	Sodium(sodium int) NutritionFactsBuilder
	Carbohydrate(carbohydrate int) NutritionFactsBuilder
	Build() NutritionFacts
}

// Step 4: Implement methods for builder
//		===> Kinh nghi?m cho th?y, ph?i implement h?t t?t c? c?c interface th? m?i kh?ng
//			 b? b?o l?i v? vi?c ch?a implement h?t c?c methods trong interface
func (nutri *nutritionFacts) Serving(servingSize int, servings int) NutritionFactsBuilder {
	nutri.serveringSize = servingSize
	nutri.serverings = servings
	return nutri
}

// Step 4:
func (nutri *nutritionFacts) Calories(calories int) NutritionFactsBuilder {
	nutri.calories = calories
	return nutri
}

// Step 4:
func (nutri *nutritionFacts) Sodium(sodium int) NutritionFactsBuilder {
	nutri.sodium = sodium
	return nutri
}

// Step 4:
func (nutri *nutritionFacts) Carbohydrate(carbohydrate int) NutritionFactsBuilder {
	nutri.carbohydrate = carbohydrate
	return nutri
}

// Step 5: Implement method Build() ?? set up h?t c?c paramter
// ===> Meaning that: Build() methods must return a reference pointer to object NutritionFacts
func (nutri *nutritionFacts) Build() NutritionFacts {
	return nutri
}

// Step 6: Must create a initial object from nutritionFacts
func NewNutritionFacts() NutritionFactsBuilder {
	// Important: This will return a variable from nutritionFacts
	return &nutritionFacts{}
}

func main() {
	// Case 1: Full parameter
	//		We must use NewNutritionFacts() to return a variable,
	//		then we can set parameter from input
	nutri1 := NewNutritionFacts().Serving(270, 80).Calories(100).Sodium(35).Carbohydrate(27).Build()
	fmt.Print(nutri1)
	fmt.Printf("\nEnd Case 1\n\n")

	// Case 2:
	nutri2 := NewNutritionFacts().Build()
	fmt.Print(nutri2)
	fmt.Printf("\nEnd Case 2\n\n")

	// Case 3:
	nutri3 := NewNutritionFacts().Serving(270, 80).Build()
	fmt.Print(nutri3)
	fmt.Printf("\nEnd Case 3\n\n")

}
