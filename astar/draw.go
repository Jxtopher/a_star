package astar

import (
	"image"
	"image/color"

	"github.com/jxtopher/a_star/worldgen"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
)

// Plot the world with the path
func Plot(file string, w worldgen.World, path []worldgen.Coordinate) {
	resolution := uint64(50)
	dest := image.NewRGBA(image.Rect(0, 0, int(w.Xsize*resolution), int(w.Ysize*resolution)))
	gc := draw2dimg.NewGraphicContext(dest)

	// Background
	white := color.RGBA{0xff, 0xff, 0xff, 0xff}
	gc.SetFillColor(white)
	draw2dkit.Rectangle(gc, float64(0), float64(0), float64(w.Xsize*resolution), float64(w.Ysize*resolution))
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
					gc,
					float64(j*resolution),
					float64(i*resolution),
					float64(j*resolution+resolution),
					float64(i*resolution+resolution),
				)
			}
		}
	}
	gc.Fill()

	// Path
	gc.SetFillColor(color.RGBA{0x66, 0x99, 0x33, 0xff})
	gc.SetStrokeColor(color.RGBA{0x66, 0x99, 0x33, 0xff})
	gc.SetLineWidth(5)

	gc.BeginPath()
	for _, coor := range path {
		gc.LineTo(float64(coor.X*resolution+resolution/2),
			float64(coor.Y*resolution+resolution/2))
	}
	gc.Stroke()

	if len(path) > 1 {
		// start
		royalblue := color.RGBA{0x41, 0x69, 0xE1, 0xff}
		gc.SetFillColor(royalblue)
		gc.SetStrokeColor(royalblue)
		draw2dkit.Rectangle(gc,
			float64(path[0].X*resolution),
			float64(path[0].Y*resolution),
			float64(path[0].X*resolution+resolution),
			float64(path[0].Y*resolution+resolution))
		gc.Fill()

		// end
		red := color.RGBA{0xFF, 0x0, 0x0, 0xff}
		gc.SetFillColor(red)
		gc.SetStrokeColor(red)
		draw2dkit.Rectangle(gc,
			float64(path[len(path)-1].X*resolution),
			float64(path[len(path)-1].Y*resolution),
			float64(path[len(path)-1].X*resolution+resolution),
			float64(path[len(path)-1].Y*resolution+resolution))
		gc.Fill()
	}

	// save to file
	draw2dimg.SaveToPngFile(file+".png", dest) //"./results/"+
}
