// Surface computes an SVG rendering of a 3-d function.
package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320            //canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

type svgFunc func(x, y float64) float64

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Did not enter the necessary arguments: surface|eggbox")
	}
	var f svgFunc
	switch os.Args[1] {
	case "surface":
		f = surface
	case "eggbox":
		f = eggbox
	default:
		fmt.Fprintln(os.Stderr, "Did not provide valid argument: surface|eggbox")
	}

	fmt.Printf("<svg smlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, f)
			bx, by := corner(i, j, f)
			cx, cy := corner(i, j+1, f)
			dx, dy := corner(i+1, j+1, f)
			if hasNaN(ax, ay, bx, by, cx, cy, dx, dy) {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' />\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int, f svgFunc) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D canvas (sx,sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func surface(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func eggbox(x, y float64) float64 {
	return math.Sin(x) * math.Cos(y) // Egg box
}

func hasNaN(values ...float64) bool {
	for _, v := range values {
		if math.IsNaN(v) {
			return true
		}
	}
	return false
}
