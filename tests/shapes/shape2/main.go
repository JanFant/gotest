package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"

	//"os"

	"./shapes"
)

func main() {
	log.SetFlags(0)
	const width, height = 1000, 700
	img := shapes.FilledImage(width, height, color.RGBA{0xff, 0xff, 0xff, 0xff})
	x, y := width/4, height/4
	red := color.RGBA{0xff, 0, 0, 0xff}
	blue := color.RGBA{0, 0, 0xff, 0xff}
	if len(os.Args) == 1 {
		fmt.Println("Using NewCircle() & NewRegularPolygon()")
		circle := shapes.NewCircle(blue, 90)
		circle.SetFill(red)
		octagon := shapes.NewRegularPolygon(red, 75, 8)
		octagon.SetFill(blue)
		polygon := shapes.NewRegularPolygon(image.Black, 65, 4)
		if err := shapes.DrawShapes(img, x, y, circle, octagon, polygon); err != nil {
			fmt.Println(err)
		}
		sanityCheack("circle", circle)
		sanityCheack("octagon", octagon)
		sanityCheack("polygon", polygon)
	} else {
		fmt.Println("Using New")
		if _, err := shapes.New("Misshapen", shapes.Option{blue, 5}); err != nil {
			fmt.Println("unexpectedly gor a non-nil invalid shape!")
		}
		circle, _ := shapes.New("circle", shapes.Option{red, 6})
		circle.SetFill(red)
		circle.(shapes.Radiuser).SetRadius(90)
		octagon, _ := shapes.New("octagon", shapes.Option{blue, 2})
		octagon.SetFill(blue)
		if octagon, ok := octagon.(shapes.Radiuser); ok {
			octagon.SetRadius(20)
		}
		polygon, _ := shapes.New("square", shapes.Option{Radius: 65})
		if err := shapes.DrawShapes(img, x, y, circle, octagon, polygon); err != nil {
			fmt.Println(err)
		}
		sanityCheack("circle", circle)
		sanityCheack("octagon", octagon)
		sanityCheack("polygon", polygon)
	}

	polygon := shapes.NewRegularPolygon(color.RGBA{0, 0x7F, 0, 0xFF}, 65, 4)
	showShapeDetais(polygon)
	y = 30
	for i, radius := range []int{60, 55, 50, 45, 40, 80, 70, 100} {
		polygon.SetRadius(radius)
		polygon.SetSides(i + 5)
		x += radius
		y += height / 10
		if err := shapes.DrawShapes(img, x, y, polygon); err != nil {
			fmt.Println(err)
		}
	}

	filename := "shapes.png"
	if err := shapes.SaveImage(img, filename); err != nil {
		log.Println(err)
	} else {
		fmt.Println("Saved", filename)
	}
	fmt.Println("OK")

	img = shapes.FilledImage(width, height, image.White)
	x, y = width/3, height/4

}

func sanityCheack(name string, shape shapes.Shaper) {
	fmt.Print("name=", name, " ")
	fmt.Print("fill=", shape.Fill(), " ")
	if shape, ok := shape.(shapes.Radiuser); ok {
		fmt.Print("radius=", shape.Radius(), " ")
	}
	if shape, ok := shape.(shapes.Sideser); ok {
		fmt.Print("sides=", shape.Sides(), " ")
	}
	fmt.Println()
}

func showShapeDetais(shape shapes.Shaper) {
	fmt.Print("fill = ", shape.Fill(), "; ")
	if shape, ok := shape.(shapes.Radiuser); ok {
		fmt.Print("radius = ", shape.Radius(), "; ")
	}
	if shape, ok := shape.(shapes.Sideser); ok {
		fmt.Print("sides = ", shape.Sides(), " ")
	}
	fmt.Println()

}
