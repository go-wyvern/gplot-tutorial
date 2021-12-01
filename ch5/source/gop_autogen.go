package main

import classfile "github.com/go-wyvern/gplot-tutorial/ch5/classfile"

type index struct {
	classfile.Figure
	Axis1 Axis1
	Axis2 Axis2
}
type Axis1 struct {
	classfile.Axis
	*index
}
type Axis2 struct {
	classfile.Axis
	*index
}

func (this *index) MainEntry() {
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot-tutorial/ch5/source/index.plot:6
	classfile.Gopt_Figure_Run(this, 1, 2)
}
func main() {
	classfile.Gopt_Figure_Main(new(index))
}
func (this *Axis1) Main() {
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot-tutorial/ch5/source/Axis1.axis:1
	x := classfile.Linspace(0, 2*classfile.Pi, 20)
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot-tutorial/ch5/source/Axis1.axis:2
	this.Plot__1(x, classfile.Sin__1(x))
}
func (this *Axis2) Main() {
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot-tutorial/ch5/source/Axis2.axis:1
	x := classfile.Linspace(0, 2*classfile.Pi, 20)
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot-tutorial/ch5/source/Axis2.axis:2
	this.Plot__1(x, classfile.Cos__1(x))
}
