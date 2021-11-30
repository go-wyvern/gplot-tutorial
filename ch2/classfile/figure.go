package classfile

import (
	"image"
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

const (
	GopPackage = true
	Gop_game   = "Figure"
)

type Ploter interface {
	initPlot()
	finishPlot()
}

// Gopt_Figure_Main is required by Go+ compiler as the entry of a .plot project.
func Gopt_Figure_Main(plot Ploter) {
	plot.initPlot()
	defer plot.finishPlot()
	plot.(interface{ MainEntry() }).MainEntry()
}

type Figure struct {
	axis *Axis
	app  fyne.App
}

func (f *Figure) initPlot() {
	f.axis = &Axis{
		p: plot.New(),
	}
	f.app = app.New()
}
func (f *Figure) finishPlot() {
	// f.axis.p.Save(4*vg.Inch, 4*vg.Inch, "figure.png")
	f.show(f.draw())
}

func (f *Figure) Plot(x, y []float64) {
	f.axis.Plot(x, y)
}

func (f *Figure) draw() image.Image {
	w, h := 4*vg.Inch, 3*vg.Inch
	convas := vgimg.NewWith(vgimg.UseWH(w, h), vgimg.UseBackgroundColor(color.Gray16{0xaaaa}))
	dc := draw.New(convas)
	f.axis.p.Draw(dc)
	return convas.Image()
}

func (f *Figure) show(img image.Image) {
	image := canvas.NewImageFromImage(img)
	image.FillMode = canvas.ImageFillOriginal
	w := f.app.NewWindow("Figure")
	w.SetContent(image)
	w.SetPadded(false)
	w.ShowAndRun()
}

func Linspace(l, r float64, n int) []float64 {
	s := make([]float64, n)
	dst := floats.Span(s, l, r)
	return dst
}
