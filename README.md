golang-chinese-to-pinyin
========================

中文转拼音golang实现


example
```golang
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

    py.Split = "-"
    py.Upper = false

    p, _ = py.Convert(s)
    fmt.Println(p)
}
```
