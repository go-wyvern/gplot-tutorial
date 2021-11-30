package main

import (
	classfile "github.com/go-wyvern/gplot-tutorial/ch3/classfile"
	math "math"
)

type index struct {
	classfile.Figure
}

func (this *index) MainEntry() {
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot-tutorial/ch3/source/index.plot:1
	x := classfile.Linspace(0, 2*math.Pi, 20)
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot-tutorial/ch3/source/index.plot:2
	this.Plot__1(x, x.Gop_Mul(x))
}
func main() {
	classfile.Gopt_Figure_Main(new(index))
}
