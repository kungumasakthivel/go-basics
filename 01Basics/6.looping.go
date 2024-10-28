package main
import "fmt"

func main() {
	// go only has for loop

	for i:=0; i<11; i++ {
		fmt.Println(i)
	}

	// infinity for loop
	for {
		fmt.Println("infinity")
		break
	}

	i:=0
	for i < 3 {
		i += 2
	}
	fmt.Println(i)

	// array
	rvar:= []string{"zero", "one", "two", "three"}
	for i, j:= range rvar {
		fmt.Println(i, j)
	}

	// iter srting
	for i, j:= range "abcd" {
		fmt.Println(i, j)
	}

	// map
	mmap := map[int] string {
		01: "day",
		02: "month",
		03: "year",
	}
	for key, val:= range mmap {
		fmt.Println(key, val)
	}

	// channel
	chnl:= make(chan int) 
	go func() {
		chnl <- 100
		chnl <- 1000
		chnl <- 10000
		close(chnl)
	}()
	for i:= range chnl {
		fmt.Println(i)
	}
}