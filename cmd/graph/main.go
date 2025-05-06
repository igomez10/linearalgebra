package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"

	"github.com/igomez10/linearalgebra"
)

var redColor = color.RGBA{255, 0, 0, 255}       // red
var grayColor = color.RGBA{200, 200, 200, 255}  // gray
var blackColor = color.RGBA{0, 0, 0, 255}       // black
var whiteColor = color.RGBA{255, 255, 255, 255} // white
var blueColor = color.RGBA{0, 0, 255, 255}      // blue
var yellowColor = color.RGBA{255, 255, 0, 255}  // yellow
var greenColor = color.RGBA{0, 255, 0, 255}     // green
var brownColor = color.RGBA{165, 42, 42, 255}   // brown

func main() {
	const width, height = 1000, 1000
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Fill background
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, whiteColor)
		}
	}

	Draw2DVector(width, 0, img, &blackColor)   // x
	Draw2DVector(-width, 0, img, &blackColor)  // -x
	Draw2DVector(0, height, img, &blackColor)  // y
	Draw2DVector(0, -height, img, &blackColor) // -y

	for i := 0; i < width; i += width / 10 {
		for j := 0; j < height; j++ {
			if i == width/2 || j == height/2 {
				continue
			}
			img.Set(i, j, grayColor) // x-axis
			img.Set(j, i, grayColor) // y-axis
		}
	}

	vectorA := []float64{100, 100}
	vectorB := []float64{100, 200}
	summedVector := linearalgebra.AddMatrices([][]float64{vectorA}, [][]float64{vectorB})[0]

	fmt.Println("vectorA", vectorA)
	Draw2DVector(vectorA[0], vectorA[1], img, &redColor)
	Draw2DVector(vectorB[0], vectorB[1], img, &brownColor)
	Draw2DVector(summedVector[0], summedVector[1], img, &blackColor)

	fmt.Println("summed vector ", summedVector)
	fmt.Println("the angle between vector A and B is ", linearalgebra.GetAngleBetweenVectors(vectorA, vectorB))

	// Save to file
	f, err := os.Create("3dplot.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		panic(err)
	}

	println("3D plot saved as 3dplot.png")
}

func GetVectorLength(vector []int) float64 {
	var powed float64 = 0
	for i := range vector {
		powed += math.Pow(float64(vector[i]), 2)
	}

	return math.Sqrt(powed)
}

// Draw2DVector draws a 2D vector on the image
// with the origin at the center of the image.
func Draw2DVector(x, y float64, img *image.RGBA, color *color.RGBA) {
	if color == nil {
		color = &blackColor
	}

	originX := img.Bounds().Max.X / 2
	originY := img.Bounds().Max.Y / 2

	linePoints := BresenhamLine(originX, originY, originX+int(x), originY-int(y))
	for i := range linePoints {
		img.Set(linePoints[i][0], linePoints[i][1], color)
	}
}

func Max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func Equal(imgA, imgB *image.RGBA) bool {
	if imgA.Bounds().Max.X != imgB.Bounds().Max.X {
		return false
	}

	if imgA.Bounds().Max.Y != imgB.Bounds().Max.Y {
		return false
	}

	for i := imgA.Bounds().Min.X; i <= imgA.Bounds().Max.X; i++ {
		for j := imgB.Bounds().Min.Y; j <= imgB.Bounds().Max.Y; j++ {
			if !EqualColor(imgA.At(i, j), imgB.At(i, j)) {
				return false
			}
		}
	}

	return true
}

func EqualColor(colorA, colorB color.Color) bool {
	redA, greenA, blueA, alphaA := colorA.RGBA()
	redB, greenB, blueB, alphaB := colorB.RGBA()

	if redA != redB {
		return false
	}

	if greenA != greenB {
		return false
	}

	if blueA != blueB {
		return false
	}

	if alphaA != alphaB {
		return false
	}

	return true
}

func BresenhamLine(x0, y0, x1, y1 int) [][]int {
	points := [][]int{}
	dx := int(math.Abs(float64(x1 - x0)))
	dy := int(math.Abs(float64(y1 - y0)))

	sx := 1
	if x0 > x1 {
		sx = -1
	}
	sy := 1
	if y0 > y1 {
		sy = -1
	}

	err := dx - dy

	for {
		points = append(points, []int{x0, y0})
		if x0 == x1 && y0 == y1 {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}

	return points
}
