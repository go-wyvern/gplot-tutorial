package classfile

import (
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
)

type Axis struct {
	p *plot.Plot
}

func (c *Axis) Plot__0(x, y []float64) {
	c.addLine(newVec(x), newVec(y))
}

func (c *Axis) Plot__1(x, y Vector) {
	c.addLine(x, y)
}

func (c *Axis) addLine(x, y Vector) {
	if x.Len() != y.Len() {
		panic("The length of the slice x is not equal to the slice y")
	}
	xys := plotter.XYs{}

	for i := 0; i < x.Len(); i++ {
		xys = append(xys, plotter.XY{
			X: x.AtVec(i),
			Y: y.AtVec(i),
		})
	}
	line, err := plotter.NewLine(xys)
	if err != nil {
		log.Fatal(err)
	}
	line.Color = plotutil.Color(0)
	c.p.Add(line)
}
