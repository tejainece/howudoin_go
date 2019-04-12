package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 1000; i++ {
		lat := rand.Float32() * 180
		long := rand.Float32() * 90
		fmt.Print("INSERT INTO ex1 (id, loc) VALUES ('", i, "', st_geographyfromtext('point(", lat, " ", long, ")'));")
		fmt.Println()
	}
}