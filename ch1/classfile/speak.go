package classfile

import "fmt"

const (
	GopPackage = true
	Gop_game   = "Speak"
)

type Speak struct{}

func Gopt_Speak_Main(i interface{}) {
	i.(interface{ MainEntry() }).MainEntry()
}

func Say(s string) {
	fmt.Println(s)
}
