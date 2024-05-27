package main

import (
	"fmt"

	"github.com/katnich/cinema/movie"
	"github.com/katnich/cinema/ticket"
)

// init จะถูกเรียกทุกครั้งที่ run จะเรียกตามที่ import เรียงจาก a-z
func init() {
	fmt.Println("init: main")
}

func main() {
	movieName := movie.FindName("tt4154796")
	fmt.Println(movieName)
	movie.Review(movieName, 8.4)
	ticket.Buy(movieName)
}
