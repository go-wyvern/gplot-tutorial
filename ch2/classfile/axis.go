package classfile

import (
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

type Axis struct {
	p *plot.Plot
}

func (c *Axis) Plot(x, y []float64) {
	if len(x) != len(y) {
		panic("The length of the slice x is not equal to the slice y")
	}
	xys := plotter.XYs{}

	for i := 0; i < len(x); i++ {
		xys = append(xys, plotter.XY{
			X: x[i],
			Y: y[i],
		})
	}
	line, err := plotter.NewLine(xys)
	if err != nil {
		log.Fatal(err)
	}
	c.p.Add(line)
}
