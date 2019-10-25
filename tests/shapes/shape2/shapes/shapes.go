package shapes

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var saneLength, saneRadius, saneSides func(int) int

func init() {
	saneLength = makeBoundedIntFunc(1, 4096)
	saneRadius = makeBoundedIntFunc(1, 1024)
	saneSides = makeBoundedIntFunc(3, 60)
}

type Shaper interface {
	Drawer
	Filler
}
type Drawer interface {
	Draw(igm draw.Image, x, y int) error
}

type Filler interface {
	Fill() color.Color
	SetFill(fill color.Color)
}

type Radiuser interface {
	Radius() int
	SetRadius(radius int)
}

type Sideser interface {
	Sides() int
	SetSideser(sides int)
}

type shape struct {
	fill color.Color
}

func newShape(fill color.Color) shape {
	if fill == nil {
		fill = color.Black
	}
	return shape{fill}
}

func (shape shape) Fill() color.Color {
	return shape.fill
}

func (shape *shape) SetFill(fill color.Color) {
	if fill == nil {
		fill = color.Black
	}
	shape.fill = fill
}

type Circle struct {
	shape
	radius int
}

func NewCircle(fill color.Color, radius int) *Circle {
	return &Circle{newShape(fill), saneRadius(radius)}
}

func (circle *Circle) Radius() int {
	return circle.radius
}

func (circle *Circle) SetRadius(radius int) {
	circle.radius = saneRadius(radius)
}

func (circle *Circle) String() string {
	return fmt.Sprintf("circle(fill=%v, radius=%d)", circle.Fill(), circle.Radius())
}

func (circle *Circle) Draw(img draw.Image, x, y int) error {
	if err := checkBounds(img, x, y); err != nil {
		return err
	}

	fill, radius := circle.fill, circle.radius
	x0, y0 := x, y
	f := 1 - radius
	ddFx, ddFy := 1, -2*radius
	x, y = 0, radius

	img.Set(x0, y0+radius, fill)
	img.Set(x0, y0-radius, fill)
	img.Set(x0+radius, y0, fill)
	img.Set(x0-radius, y0, fill)

	for x < y {
		if f >= 0 {
			y--
			ddFy += 2
			f += ddFy
		}
		x++
		ddFx += 2
		f += ddFx
		img.Set(x0+x, y0+y, fill)
		img.Set(x0-x, y0+y, fill)
		img.Set(x0+x, y0-y, fill)
		img.Set(x0-x, y0-y, fill)
		img.Set(x0+y, y0+x, fill)
		img.Set(x0-y, y0+x, fill)
		img.Set(x0+y, y0-x, fill)
		img.Set(x0-y, y0-x, fill)
	}
	return nil
}

func checkBounds(img image.Image, x, y int) error {
	if !image.Rect(x, y, x, y).In(img.Bounds()) {
		return fmt.Errorf("%s(): point (%d, %d) is outside the image\n", caller(1), x, y)
	}
	return nil
}

type RadularPolygon struct {
	*Circle
	sides int
}

func NewRegularPolygon(fill color.Color, radius, sides int) *RadularPolygon {
	return &RadularPolygon{NewCircle(fill, radius), saneSides(sides)}
}

func (polygon *RadularPolygon) Sides() int {
	return polygon.sides
}

func (polygon *RadularPolygon) SetSides(sides int) {
	polygon.sides = saneSides(sides)
}

func (polygon *RadularPolygon) String() string {
	return fmt.Sprintf("Polygon fill = %v, radius = %d, sides = %d ", polygon.Fill(), polygon.Radius(), polygon.Sides())
}

func (polygon *RadularPolygon) Draw(img draw.Image, x, y int) error {
	if err := checkBounds(img, x, y); err != nil {
		return err
	}

	points := getPoint(x, y, polygon.Sides(), float64(polygon.Radius()))
	for i := 0; i < polygon.Sides(); i++ {
		drawLine(img, points[i], points[i+1], polygon.Fill())
	}
	return nil
}

func getPoint(x, y, sides int, radius float64) []image.Point {
	points := make([]image.Point, sides+1)
	fullCircle := 2 * math.Pi
	x0, y0 := float64(x), float64(y)
	for i := 0; i < sides; i++ {
		sig := float64(float64(i) * fullCircle / float64(sides))
		x1 := x0 + (radius * math.Sin(sig))
		y1 := y0 + (radius * math.Cos(sig))
		points[i] = image.Pt(int(x1), int(y1))
	}
	points[sides] = points[0]
	return points
}

func drawLine(img draw.Image, start, end image.Point, fill color.Color) {
	x0, x1 := start.X, end.X
	y0, y1 := start.Y, end.Y
	dx := math.Abs(float64(x1 - x0))
	dy := math.Abs(float64(y1 - y0))
	if dx >= dy {
		if x0 > x1 {
			x0, y0, x1, y1 = x1, y1, x0, y0
		}
		y := y0
		yStep := 1
		if y0 > y1 {
			yStep = -1
		}
		remainder := float64(int(dx/2)) - dx
		for x := x0; x <= x1; x++ {
			img.Set(x, y, fill)
			remainder += dy
			if remainder >= 0.0 {
				remainder -= dx
				y += yStep
			}
		}
	} else {
		if y0 > y1 {
			x0, y0, x1, y1 = x1, y1, x0, y0
		}
		x := x0
		xStep := 1
		if x0 > x1 {
			xStep = -1
		}
		remainder := float64(int(dy/2)) - dy
		for y := y0; y <= y1; y++ {
			img.Set(x, y, fill)
			remainder += dx
			if remainder >= 0.0 {
				remainder -= dy
				x += xStep
			}
		}

	}

}

type Option struct {
	Fill   color.Color
	Radius int
}

func New(shape string, option Option) (Shaper, error) {
	sidesForShape := map[string]int{"triangle": 3, "square": 4, "pentagon": 5, "hexgon": 6, "heptagon": 7, "octagon": 8, "enneagon": 9, "nonagon": 9, "decagon": 10}
	if sides, found := sidesForShape[shape]; found {
		return NewRegularPolygon(option.Fill, option.Radius, sides), nil
	}
	if shape != "circle" {
		return nil, fmt.Errorf("shapes.New(): invalid shape '%s'", shape)
	}
	return NewCircle(option.Fill, option.Radius), nil
}

func makeBoundedIntFunc(min, max int) func(int) int {
	return func(x int) int {
		valid := x
		switch {
		case x < min:
			valid = min
		case x > max:
			valid = max
		}
		if valid != x {
			log.Printf("%s(): replaced %d with %d\n", caller(1), x, valid)
		}
		return valid
	}
}

func caller(steps int) string {
	name := "?"
	if pc, _, _, ok := runtime.Caller(steps + 1); ok {
		name = filepath.Base(runtime.FuncForPC(pc).Name())
	}
	return name
}

func DrawShapes(img draw.Image, x, y int, shapes ...Shaper) error {
	for _, shape := range shapes {
		if err := shape.Draw(img, x, y); err != nil {
			return err
		}
	}
	return nil
}

func SaveImage(img image.Image, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	switch strings.ToLower(filepath.Ext(filename)) {
	case ".jpg", "jpeg":
		return jpeg.Encode(file, img, nil)
	case ".png":
		return png.Encode(file, img)
	}
	return fmt.Errorf("shapes.SaveImage(): '%s' has an unrecognized "+"suffix", filename)
}

func FilledImage(width, height int, fill color.Color) draw.Image {
	if fill == nil {
		fill = color.Black
	}
	width = saneLength(width)
	height = saneLength(height)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), &image.Uniform{fill}, image.ZP, draw.Src)
	return img
}
