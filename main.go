package main

import (
	"fmt"
	"os"
)

func check(e error, s string) {
	if e != nil {
		fmt.Fprintf(os.Stderr, s, e)
		os.Exit(1)
	}
}

func main() {
	// default size of image
	x := 400
	y := 300

	const color = 255.0

	f, err := os.Create("out.ppm")

	defer f.Close()

	check(err, "Error opening file: %v\n")

	_, err = fmt.Fprintf(f, "P3\n%d %d\n255\n", x, y)

	check(err, "Error writing to file: %v\n")

	// Writes to each pixel with r/g/b values
	// from top left to bottom right
	for j := y - 1; j >= 0; j-- {
		for i := 0; i < x; i++ {
			// red and green values range from
			// 0.0 tp 1.0
			v := Vector{
				x: float64(i) / float64(x),
				y: float64(j) / float64(y),
				z: 0.2}

			ir := int(color * v.x)
			ig := int(color * v.y)
			ib := int(color * v.z)

			_, err = fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)

			check(err, "Error writing to file: %v\n")
		}
	}

}
