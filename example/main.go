package main

import (
	"fmt"
	"github.com/l2x/golang-chinese-to-pinyin"
)

func main() {
	s := "天气不错"

	py := Pinyin.New()
	p, _ := py.Convert(s)

	fmt.Println(p)
	//Tian Qi Bu Cuo

	py.Split = "-"
	py.Upper = false

	p, _ = py.Convert(s)
	fmt.Println(p)
	//tian-qi-bu-cuo
}
