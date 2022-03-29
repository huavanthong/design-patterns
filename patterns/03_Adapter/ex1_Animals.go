package animal

type Animal interface {
  Move()
}

// Cat is a concrete animal since it implements the method Move
type Cat struct {}

func (c *Cat) Move() {}

// and somewhere in the code we need to use the crocodile type which is often not our code and this Crocodile type does not implement the Animal interface
// but we need to use a crocodile as an animal

type Crocodile struct {}

func (c *this.Crocodile) Slither() {}

// we create an CrocodileAdapter struct that dapts an embeded crocodile so that it can be usedd as an Animal

type CrocodileAdapter struct {
  *Crocodile
}

func NeweCrocodile() *CrocodileAdapter {
  return &CrocodileAdapter{new(Crocodile)}
}

func (this *CrocodileAdapter) Move() {
  this.Slither()
}

// usage
import "animal"

var animals []animal.Animal
animals = append(animals, new(animals.Cat))
animals = append(animals, new(animals.NewCrocodile()))

for  entity := range animals {
	entity.Move()
} 