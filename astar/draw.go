package astar

import (
	"fmt"
	"image"
	"image/color"

	"github.com/jxtopher/a_star/worldgen"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
)

// Plot the world with the path
func Plot(file string, w worldgen.World, path []worldgen.Coordinate) {
	dest := image.NewRGBA(image.Rect(0, 0, int(w.Xsize*50), int(w.Ysize*50)))
	gc := draw2dimg.NewGraphicContext(dest)

	// Background
	white := color.RGBA{0xff, 0xff, 0xff, 0xff}
	gc.SetFillColor(white)
	draw2dkit.Rectangle(gc, float64(0), float64(0), float64(w.Xsize*50), float64(w.Ysize*50))
	gc.Fill()

	// sqaure -> blocks
	black := color.RGBA{0x0, 0x0, 0x0, 0xff}
	gc.SetFillColor(black)
	gc.SetStrokeColor(black)

	var i, j uint64
	for i = 0; i < w.Xsize; i++ {
		for j = 0; j < w.Ysize; j++ {
			if w.Ground[i][j] == worldgen.Bloc {
				draw2dkit.Rectangle(
					gc, float64(j*50), float64(i*50), float64(j*50+50), float64(i*50+50),
				)
			}
		}
	}
	gc.Fill()

	// Path
	gc.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
	gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	gc.SetLineWidth(5)

	gc.BeginPath()
	for _, coor := range path {
		fmt.Println(coor)
		gc.LineTo(float64(coor.X*50), float64(coor.Y*50))
	}
	gc.Stroke()

	// start
	bleu := color.RGBA{0x0, 0x0, 0xFF, 0xff}
	gc.SetFillColor(bleu)
	gc.SetStrokeColor(bleu)
	draw2dkit.Rectangle(gc, float64(path[0].X*50), float64(path[0].Y*50), float64(path[0].X*50+50), float64(path[0].Y*50+50))
	gc.Fill()

	// end
	red := color.RGBA{0xFF, 0x0, 0x0, 0xff}
	gc.SetFillColor(red)
	gc.SetStrokeColor(red)
	draw2dkit.Rectangle(gc, float64(path[len(path)-1].X*50), float64(path[len(path)-1].Y*50), float64(path[len(path)-1].X*50+50), float64(path[len(path)-1].Y*50+50))
	gc.Fill()

	// save to file
	draw2dimg.SaveToPngFile(file+".png", dest) //"./results/"+
}
