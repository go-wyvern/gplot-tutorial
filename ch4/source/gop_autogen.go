package main

import classfile "github.com/go-wyvern/gplot-tutorial/ch4/classfile"

type index struct {
	classfile.Figure
}

func (this *index) MainEntry() {
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot-tutorial/ch4/source/index.plot:1
	x := classfile.Linspace(0, 2*classfile.Pi, 20)
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot-tutorial/ch4/source/index.plot:2
	this.Subplot(1, 2, 1)
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot-tutorial/ch4/source/index.plot:3
	this.Plot__1(x, classfile.Sin__1(x))
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot-tutorial/ch4/source/index.plot:4
	this.Subplot(1, 2, 2)
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot-tutorial/ch4/source/index.plot:5
	this.Plot__1(x, classfile.Cos__1(x))
}
func main() {
	classfile.Gopt_Figure_Main(new(index))
}
