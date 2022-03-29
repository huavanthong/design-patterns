// Shape is the component
type Shape interface {
	Draw(drawer *Drawer) error
}

// Square and Circle are leaves
type Square struct {
	Location Point
	Size float64
}

func (square *Square) Draw(drawer *Drawer) error {
	return drawer.DrawRect(Rect{
		Location: square.Location,
		Size: Size{
			Height: square.Side,
			Width:  square.Side,
		},
	})
}

type Circle struct {
	Center Point
	Radius float64
}

func (circle *Circle) Draw(drawer *Drawer) error {
	rect := Rect{
		Location: Point{
			X: circle.Center.X - circle.Radius,
			Y: circle.Center.Y - circle.Radius,
		},
		Size: Size{
			Width:  2 * circle.Radius,
			Height: 2 * circle.Radius,
		},
	}

	return drawer.DrawEllipseInRect(rect)
}

// Layer is the composite
type Layer struct {
	Shapes []Shape
}

func (layer *Layer) Draw(drawer *Drawer) error {
	for _, shape := range layer.Shapes {
		if err := shape.Draw(drawer); err != nil {
			return err
		}
		fmt.Println()
	}

	return nil
}

// usage
circle := &photoshop.Circle{
	Center: photoshop.Point{X: 100, Y: 100},
	Radius: 50,
}

square := &photoshop.Square{
	Location: photoshop.Point{X: 50, Y: 50},
	Side:     20,
}

layer := &photoshop.Layer{
	Elements: []photoshop.Shapes{
		circle,
		square,
	},
}

circle.Draw(&photoshop.Drawer{})
square.Draw(&photoshop.Drawer{})
// or
layer.Draw(&photoshop.Drawer{})