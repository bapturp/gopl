package main

import "fmt"

const (
	KB = 1000
	MB = KB * KB
	GB = KB * KB
	TB = GB * KB
	PB = TB * KB
	EB = PB * KB
	ZB = EB * KB
	YB = ZB * KB
)

func main() {
	fmt.Println(ZB)
}
