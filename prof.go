// +build ignore

package main

import (
    "fmt"
    "time"
    "mmsego"
    "log"
    )

func main() {
    log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)
    var s = new(mmsego.Segmenter)
    //s.Init("../darts/darts.lib")
    //s.LoadText("../data/words.dic")
    s.Init("tmp.lib")
    t := time.Now()
    s.Mmseg("我爱天安门,哈哈")
    s.Mmseg("南京市长江大桥欢迎您?错误")
    //s.Split([]rune("南京市长江大桥欢迎您?错误"))
    s.Mmseg("营销系统开发部首届游戏大赛")
    //s.Split([]rune("营销系统开发部首届游戏大赛"))
    fmt.Printf("Duration: %v\n", time.Since(t))
}
