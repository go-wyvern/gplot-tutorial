package main

import (
	builtin "github.com/goplus/gop/builtin"

	gplot "github.com/go-wyvern/gplot"
)

type index struct {
	gplot.Figure
}

func (this *index) MainEntry() {
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot-tutorial/index.plot:1
	x := func() (_gop_ret []int) {
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot-tutorial/index.plot:1
		for _gop_it := builtin.NewRange__0(0, 6, 1).Gop_Enum(); ; {
			var _gop_ok bool
			x, _gop_ok := _gop_it.Next()
			if !_gop_ok {
				break
			}
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot-tutorial/index.plot:1
			_gop_ret = append(_gop_ret, x)
		}
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot-tutorial/index.plot:1
		return
	}()
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot-tutorial/index.plot:2
	y := func() (_gop_ret []int) {
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot-tutorial/index.plot:2
		for _gop_it := builtin.NewRange__0(0, 6, 1).Gop_Enum(); ; {
			var _gop_ok bool
			x, _gop_ok := _gop_it.Next()
			if !_gop_ok {
				break
			}
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot-tutorial/index.plot:2
			_gop_ret = append(_gop_ret, x*x)
		}
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot-tutorial/index.plot:2
		return
	}()
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot-tutorial/index.plot:3
	this.Plot(x, y)
}

func main() {
	new(index).Main()
}
