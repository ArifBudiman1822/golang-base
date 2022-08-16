package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	utc := time.Date(2022, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)

	loc, _ := time.LoadLocation("Asia/Jakarta")
	layout := "2006-01-02"
	parse, _ := time.ParseInLocation(layout, "2022-08-15", loc)
	fmt.Println(parse)
}
