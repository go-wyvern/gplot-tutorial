package main

import classfile "github.com/go-wyvern/gplot-tutorial/ch1/classfile"

type index struct {
	classfile.Speak
}

func (this *index) MainEntry() {
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot-tutorial/ch1/source/index.p:1
	classfile.Say("hello world")
}
func main() {
	classfile.Gopt_Speak_Main(new(index))
}
