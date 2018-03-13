package main

import (
	"math"
)

type Vector struct {
	x, y, z float64
}

func (v Vector) length() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func (v Vector) Dot(o Vector) float64 {
	return v.x*o.x + v.y*o.y + v.z*o.z
}

func (v Vector) Normalize() {
	l := v.length()
	v = Vector{v.x / l, v.y / l, v.z / l}
}
