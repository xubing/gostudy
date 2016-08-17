package main

import (
	"github.com/hydra13142/cube"
	"github.com/hydra13142/geom"
	"github.com/hydra13142/paint"
	"image"
	"image/color/palette"
	"image/gif"
	"math"
	"os"
)

var (
	pln *cube.Plain
	unx cube.Vector
	uny cube.Vector
)

const (
	H = 18
	W = 36
)

func init() {
	pln, _ = cube.NewPlain(cube.Point{}, cube.Vector{4, 4, 3})
	uny = cube.FromTo(cube.Point{}, pln.VerticalPoint(cube.Point{0, 0, 10})).Unit()
	unx = cube.OuterProduct(uny, cube.Vector{4, 4, 3}).Unit()
}
func main() {
	var x [H + 1][W]cube.Point
	var y [H + 1][W]geom.Point
	dz := math.Pi / H
	dxy := math.Pi * 2 / W
	for i := 0; i <= H; i++ {
		az := float64(i)*dz - math.Pi/2
		r := 140 * math.Cos(az)
		z := 140 * math.Sin(az)
		for j := 0; j < W; j++ {
			axy := float64(j) * dxy
			x[i][j] = cube.Point{math.Cos(axy) * r, math.Sin(axy) * r, z}
		}
	}
	pics := make([]*image.Paletted, 0, 20)
	img := paint.Image{
		FR: paint.Green,
		BG: paint.White,
	}
	stp := dxy / 20
	delay := make([]int, 0, 20)
	for t := 0; t < 20; t++ {
		img.Image = image.NewPaletted(image.Rect(0, 0, 300, 300), palette.Plan9)
		for i := 0; i <= H; i++ {
			for j := 0; j < W; j++ {
				ox := cube.FromTo(cube.Point{}, x[i][j])
				y[i][j] = geom.Point{cube.InnerProduct(ox, unx), cube.InnerProduct(ox, uny)}
				a, b := x[i][j].X, x[i][j].Y
				x[i][j].X = a*math.Cos(stp) - b*math.Sin(stp)
				x[i][j].Y = b*math.Cos(stp) + a*math.Sin(stp)
			}
		}
		for i := 0; i < H; i++ {
			for j := 0; j < W; j++ {
				img.Line(
					150+int(y[i][j].X),
					150-int(y[i][j].Y),
					150+int(y[i][(j+1)%W].X),
					150-int(y[i][(j+1)%W].Y),
				)
				img.Line(
					150+int(y[i][j].X),
					150-int(y[i][j].Y),
					150+int(y[i+1][j].X),
					150-int(y[i+1][j].Y),
				)
			}
		}
		pics = append(pics, img.Image.(*image.Paletted))
		delay = append(delay, 5)
	}
	file, _ := os.Create("ball.gif")
	defer file.Close()
	gif.EncodeAll(file, &gif.GIF{
		Image:     pics,
		Delay:     delay,
		LoopCount: 5 * len(delay),
	})
}
