package main

import "fmt"

func main(){
	countdown := 10
	for countdown>0{
		fmt.Printf("countdown: %d\n",countdown)
		countdown--
	}
}