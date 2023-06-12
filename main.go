package main

import "fmt"

type ByteSize float64

const (
	KB ByteSize = 1 << (10 * (iota + 1))
	MB
	GB
)Go Notebook Kernel: Open Interactive.

func (b ByteSize) String() string {
	switch {
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}

func main() {
	fmt.Println(1001*KB, 2.5*MB, 3.5*GB)
	fmt.Println(ByteSize(121000000))
}

