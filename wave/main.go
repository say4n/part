package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"math"
	"math/rand"
)

type Point struct {
	x float64
	y float64
}

func randIntInRange(lo, hi int64) int64 {
	return rand.Int63n(hi-lo) + int64(lo)
}

func randNormal(mu, sigma float64) float64 {
	return rand.NormFloat64()*sigma + mu
}

func normalPDF(x, mu, sigma float64) float64 {
	sigma2 := math.Pow(sigma, 2)
	num := math.Exp(-math.Pow((x-mu), 2) / (2 * sigma2))
	den := math.Sqrt(2 * math.Pi * sigma2)

	return num / den
}

func main() {
	fmt.Println("Waves")

	rand.Seed(100)

	const (
		nLines  = 80.0
		ySep    = 10.0
		xLen    = 600.0
		nPoints = 100.0
		size    = 1024

		yPad = (size - nLines*ySep) / 2
		xPad = (size - xLen) / 2
		dx   = xLen / nPoints
	)

	ctx := gg.NewContext(size, size)

	ctx.DrawRectangle(0, 0, float64(size), float64(size))
	ctx.SetHexColor("#000000")
	ctx.Fill()

	for i := 0.0; i < nLines; i += 1 {
		x1 := xPad
		x2 := xPad + xLen
		y := yPad + i*ySep

		midpoint := int64((x2 + x1) / 2)
		nModes := randIntInRange(1, 2)
		mus := make([]float64, nModes)
		sigmas := make([]float64, nModes)

		for m := 0; m < int(nModes); m++ {
			mus[m] = rand.NormFloat64()*float64(randIntInRange(-50, 50)) + float64(midpoint)
			sigmas[m] = randNormal(36, 60)
		}

		points := make([]Point, 0)
		var w = y

		for x := x1; x < x2; x += dx {
			noise := 0.0

			for m := 0; m < int(nModes); m++ {
				noise += normalPDF(x, mus[m], sigmas[m])
			}

			yy := 0.3*w + 0.7*(y-600*noise+noise*rand.NormFloat64()*200+rand.NormFloat64())

			points = append(points, Point{x, yy})
			ctx.LineTo(x, yy)

			w = yy
		}

		ctx.ClosePath()
		ctx.SetHexColor("#000000")
		ctx.FillPreserve()

		ctx.ClearPath()

		for _, p := range points {
			ctx.LineTo(p.x, p.y)
		}

		ctx.SetLineWidth(0.65)
		ctx.SetHexColor("#ffffff")
		ctx.Stroke()

	}

	ctx.SavePNG("wave.png")
}
