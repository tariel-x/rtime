package main

import (
	"fmt"
	"time"

	"github.com/tariel-x/rtime"
)

func main() {
	nowTime := rtime.Now()
	specificTime := rtime.Date(2023, 03, 1, 02, 48, 05, 0, time.UTC)
	inheritedTime := rtime.RTime{Time: time.Now()}

	fmt.Println(nowTime.Format(rtime.GOST2003Word))
	fmt.Println(nowTime.Format("январь/January, 2-е"))
	fmt.Println(specificTime.Format(rtime.GOST2016Numeric))
	fmt.Println(inheritedTime.Format(rtime.GOST2003NumericReverse))
}
