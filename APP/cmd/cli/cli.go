package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	bio "github.com/huavanthong/design-patterns/APP/boundary/io"
	"github.com/huavanthong/design-patterns/APP/builder"
	"github.com/huavanthong/design-patterns/APP/interactor"
	"github.com/huavanthong/design-patterns/APP/repository"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "your-app",
		Short: "A brief description of your application",
	}

	rectangleCmd := &cobra.Command{
		Use:   "rectangle",
		Short: "Create a rectangle shape",
		Run: func(cmd *cobra.Command, args []string) {
			// Parse command-line options
			width, _ := cmd.Flags().GetFloat64("width")
			height, _ := cmd.Flags().GetFloat64("height")

			// Create input boundary
			inputBoundary := bio.NewRectangleInput(width, height)

			// Create output boundary
			outputBoundary := &bio.RectangleOutput{}

			// Create builder
			shapeBuilder := builder.NewShapeBuilder()

			// Create interactor
			interactor := interactor.NewRectangleInteractor(shapeBuilder, repository.NewRectangleRepository())

			// Create and display the rectangle shape
			rectangle := interactor.Create(inputBoundary)
			outputBoundary.Display(rectangle)
		},
	}

	rectangleCmd.Flags().Float64("width", 0, "Width of the rectangle")
	rectangleCmd.Flags().Float64("height", 0, "Height of the rectangle")

	rootCmd.AddCommand(rectangleCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
