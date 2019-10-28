package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"math"
	"math/rand"
	"time"
)

const (
	minWidth  float64 = 20
	minHeight float64 = 20
)

var (
	colors = []string{
		"#fff001",
		"#ff0101",
		"#0101fd",
		"#f9f9f9",
		"#f9f9f9",
		"#f9f9f9",
		"#f9f9f9",
		"#30303a",
	}
)

type Point struct {
	x float64
	y float64
}

type Rectangle struct {
	bottomleft Point
	topright   Point
}

func (r Rectangle) width() float64 {
	return math.Abs(r.topright.x - r.bottomleft.x)
}

func (r Rectangle) height() float64 {
	return math.Abs(r.topright.y - r.bottomleft.y)
}

func randIntInRange(lo, hi float64) float64 {
	return float64(rand.Int63n(int64(hi)-int64(lo)) + int64(lo))
}

func generate(r Rectangle, depth, limit int64, rects *[]Rectangle) {
	if depth == limit {
		return
	}

	if r.height() < minHeight || r.width() < minWidth {
		return
	}

	*rects = append(*rects, r)

	r1 := Rectangle{}
	r2 := Rectangle{}

	if r.width() > r.height() {
		// do a left-right split
		loc := randIntInRange(r.bottomleft.x, r.bottomleft.x+r.width())

		r1 = Rectangle{r.bottomleft, Point{loc, r.topright.y}}
		r2 = Rectangle{Point{loc, r.bottomleft.y}, r.topright}

	} else {
		// do a top-bottom split

		loc := randIntInRange(r.bottomleft.y, r.bottomleft.y+r.height())

		r1 = Rectangle{Point{r.bottomleft.x, loc}, r.topright}
		r2 = Rectangle{r.bottomleft, Point{r.topright.y, loc}}
	}

	generate(r1, depth+1, limit, rects)
	generate(r2, depth+1, limit, rects)
}

func main() {
	fmt.Println("Pete Mondrian — Grid Art")

	pad := float64(4)
	size := int(1000 + 2*pad)

	rand.Seed(time.Now().UnixNano())

	base := Rectangle{Point{pad, pad}, Point{1000, 1000}}
	rects := []Rectangle{}
	rects = append(rects, base)

	generate(base, 1, 6, &rects)

	ctx := gg.NewContext(size, size)
	ctx.SetHexColor("#000000")
	ctx.FillPreserve()

	for _, r := range rects {
		fmt.Printf("width: %v, height: %v\n", r.width(), r.height())

		x := float64(r.bottomleft.x)
		y := float64(r.bottomleft.y)
		w := float64(r.width())
		h := float64(r.height())

		ctx.DrawRectangle(x, y, w, h)
		ctx.SetHexColor(colors[rand.Intn(len(colors))])
		ctx.FillPreserve()

		weight := float64(randIntInRange(8, 10))
		ctx.SetHexColor("#000000")
		ctx.SetLineWidth(weight)
		ctx.Stroke()
	}

	ctx.SavePNG("mondrian.png")
}
