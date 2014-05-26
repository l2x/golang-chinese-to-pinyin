golang-chinese-to-pinyin
========================

中文转拼音golang实现

install
--------
```
go get github.com/l2x/golang-chinese-to-pinyin
```


example
--------

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
    //Tian Qi Bu Cuo


    //设置分隔符
    //首字母是否大写
    py.Split = "-"
    py.Upper = false
    p, _ = py.Convert(s)
    fmt.Println(p)
    //tian-qi-bu-cuo
}
```
