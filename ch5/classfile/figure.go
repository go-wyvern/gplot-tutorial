package classfile

import (
	"image"
	"image/color"
	"log"
	"reflect"

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
	Gop_sprite = "Axis"
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

// Gopt_Figure_Run is required by Go+ compiler as the entry of a .plot project.
func Gopt_Figure_Run(plot Ploter, x, y int) {
	v := reflect.ValueOf(plot).Elem()
	t := reflect.TypeOf(plot).Elem()
	p := instance(v)
	pos := 0
	for i, n := 0, v.NumField(); i < n; i++ {
		typ := t.Field(i).Type
		m, ok := reflect.PtrTo(typ).MethodByName("Main")
		if ok {
			parent, axis := instanceAxis(t.Field(i))
			pos += 1
			m.Func.Call([]reflect.Value{parent})
			p.Subplot(x, y, pos)
			p.Align[p.pos.row][p.pos.col] = axis
		}
	}
}

type Figure struct {
	Align      [][]*Axis
	pos        Positon
	w, h       float64
	rows, cols int
	app        fyne.App
}

type Positon struct {
	row, col int
}

func (f *Figure) initPlot() {
	f.w = float64(4 * vg.Inch)
	f.h = float64(3 * vg.Inch)
	f.pos = Positon{
		row: 0,
		col: 0,
	}
	f.rows = 1
	f.cols = 1
	f.Align = [][]*Axis{
		{
			&Axis{
				p: plot.New(),
			},
		},
	}
	f.app = app.New()
}
func (f *Figure) finishPlot() {
	if f.pos.col+1 == f.cols && f.pos.row+1 == f.rows {
		f.show(f.draw())
	}
}

func (f *Figure) Plot__0(x, y []float64) {
	f.Align[f.pos.row][f.pos.col].Plot__0(x, y)
}

func (f *Figure) Plot__1(x, y Vector) {
	f.Align[f.pos.row][f.pos.col].Plot__1(x, y)
}

func (f *Figure) Subplot(x, y, pos int) {
	f.rows = x
	f.cols = y
	var col int
	row := pos / y
	if pos%y == 0 {
		col = y - 1
		row = row - 1
	} else {
		col = pos%y - 1
	}

	f.pos.row = row
	f.pos.col = col

	rows := len(f.Align)
	if rows != x {
		f.Align = make([][]*Axis, x)
	}
	for j := 0; j < rows; j++ {
		if len(f.Align[j]) != y {
			f.Align[j] = make([]*Axis, y)
		}
	}

	if f.Align[row][col] == nil {
		f.Align[row][col] = &Axis{
			p: plot.New(),
		}
	}
}

func (f *Figure) draw() image.Image {
	w, h := vg.Length(f.w*float64(f.cols)), vg.Length(f.h*float64(f.rows))
	convas := vgimg.NewWith(vgimg.UseWH(w, h), vgimg.UseBackgroundColor(color.Gray16{0xaaaa}))
	dc := draw.New(convas)
	t := draw.Tiles{
		Rows:      f.rows,
		Cols:      f.cols,
		PadTop:    vg.Length(20),
		PadBottom: vg.Length(20),
		PadRight:  vg.Length(20),
		PadLeft:   vg.Length(20),
		PadX:      vg.Length(20),
		PadY:      vg.Length(20),
	}
	plots := make([][]*plot.Plot, f.rows)
	for j := 0; j < f.rows; j++ {
		plots[j] = make([]*plot.Plot, f.cols)
	}
	canvases := plot.Align(plots, t, dc)
	for j := 0; j < f.rows; j++ {
		for i := 0; i < f.cols; i++ {
			axis := f.Align[j][i]
			if axis != nil {
				plots[j][i] = axis.p
				if plots[j][i] != nil {
					plots[j][i].Draw(canvases[j][i])
				}
			}
		}
	}
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

func Linspace(l, r float64, n int) Vector {
	s := make([]float64, n)
	dst := floats.Span(s, l, r)
	return newVec(dst)
}

func instance(plotter reflect.Value) *Figure {
	fld := plotter.FieldByName("Figure")
	if !fld.IsValid() {
		log.Panicf("type %v doesn't has field gplot.Figure", plotter.Type())
	}
	return fld.Addr().Interface().(*Figure)
}

func instanceAxis(field reflect.StructField) (reflect.Value, *Axis) {
	typ := field.Type
	parent := reflect.New(typ)
	axis := parent.Elem().FieldByName("Axis")
	axis.Set(reflect.ValueOf(&Axis{
		p: plot.New(),
	}).Elem())
	return parent, axis.Addr().Interface().(*Axis)
}
