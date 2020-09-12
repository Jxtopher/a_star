package astar

import (
	"fmt"
	"image"
	"image/color"

	"github.com/jxtopher/a_star/worldgenerator"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
)

// Plot world
func Plot(w worldgenerator.World, path []worldgenerator.Coordinate) {
	// dest := draw2dpdf.NewPdf("L", "mm", "A4")
	// gc := draw2dpdf.NewGraphicContext(dest)
	dest := image.NewRGBA(image.Rect(0, 0, int(w.Xsize*10), int(w.Ysize*10)))
	gc := draw2dimg.NewGraphicContext(dest)

	// Background
	white := color.RGBA{0xff, 0xff, 0xff, 0xff}
	gc.SetFillColor(white)
	draw2dkit.Rectangle(gc, float64(0), float64(0), float64(w.Xsize*10), float64(w.Ysize*10))
	gc.Fill()

	// sqaure -> blocks
	black := color.RGBA{0x0, 0x0, 0x0, 0xff}
	gc.SetFillColor(black)
	gc.SetStrokeColor(black)

	var i, j uint64
	for i = 0; i < w.Xsize; i++ {
		for j = 0; j < w.Ysize; j++ {
			if w.Ground[i][j] == worldgenerator.Bloc {
				draw2dkit.Rectangle(gc, float64(i*10), float64(j*10), float64(i*10+10), float64(j*10+10))
			}
		}
	}
	gc.Fill()

	// //
	gc.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
	gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	gc.SetLineWidth(5)
	// gc.MoveTo(10, 10)
	for _, coor := range path {
		fmt.Println(coor)
		gc.LineTo(float64(coor.X*10), float64(coor.Y*10))
	}
	gc.Stroke()

	// // Set some properties

	// // Draw a closed shape
	// // gc.BeginPath() // Initialize a new path

	// gc.LineTo(30, 30)
	// gc.LineTo(30, 40)
	// gc.Stroke()

	// gc.Stroke()
	// gc.Close()
	// gc.FillStroke()

	// Save to file
	// draw2dpdf.SaveToPdfFile("hello.pdf", dest)
	draw2dimg.SaveToPngFile("hello.png", dest)

}
