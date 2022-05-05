package spec

import "math"

// ReferenceGamut is the standard gamut used to represent colour on a full-spectrum light source.
var ReferenceGamut = Gamut{
	R: Point{1, 0},
	G: Point{0, 1},
	B: Point{0, 0},
}

// HueGamut is the reference gamut suggested by Phillips for the Hue light bulbs.
var HueGamut = Gamut{
	R: Point{0.675, 0.322},
	G: Point{0.409, 0.518},
	B: Point{0.167, 0.04},
}

// Hue2Gamut is a custom gamut which attempts to approximate the human eye's perception
// of colour on the Hue light bulbs by reducing green slightly.
var Hue2Gamut = Gamut{
	R: Point{1, 0.0},
	G: Point{0.0, 0.95},
	B: Point{0.01, 0.02},
}

type Point struct {
	X, Y float64
}

func (p Point) Unpack() (float64, float64) {
	return p.X, p.Y
}

func (p Point) Plus(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Minus(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

func (p Point) Scale(s float64) Point {
	return Point{p.X * s, p.Y * s}
}

func (p Point) Dot(q Point) float64 {
	return p.X*q.X + p.Y*q.Y
}

func (p Point) Len() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

type Gamut struct {
	R, G, B Point
}

// Project a point from a reference gamut onto this gamut by
// performing a series of linear transforms on it.
func (g *Gamut) Project(p Point) Point {
	// We calculate our relative position along the Red and Green vectors
	// (from the Blue origin) and then scale the resulting position based
	// on the same ratio from the new Blue origin.
	rr, rg := p.X/1.0, p.Y/1.0

	nrr, nrg := g.R.Minus(g.B).Scale(rr), g.G.Minus(g.B).Scale(rg)

	return g.B.Plus(nrr).Plus(nrg)
}

func (g *Gamut) Clamp(x, y float64) (float64, float64) {
	point := triangleClosest(Point{x, y}, g.R, g.G, g.B)
	return point.X, point.Y
}

func (g *Gamut) InGamut(x, y float64) bool {
	d1 := triangleSign(Point{x, y}, g.R, g.B)
	d2 := triangleSign(Point{x, y}, g.B, g.G)
	d3 := triangleSign(Point{x, y}, g.G, g.R)

	hasNeg := d1 < 0 || d2 < 0 || d3 < 0
	hasPos := d1 > 0 || d2 > 0 || d3 > 0

	return !(hasNeg && hasPos)
}

func triangleSign(a, b, c Point) float64 {
	return (a.X-c.X)*(b.Y-c.Y) - (b.X-c.X)*(a.Y-c.Y)
}

func triangleClosest(p, a, b, c Point) Point {
	d1 := p.Minus(a).Len()
	d2 := p.Minus(b).Len()
	d3 := p.Minus(c).Len()

	dmin := math.Min(d1, math.Min(d2, d3))
	switch dmin {
	case d1:
		return a
	case d2:
		return b
	default:
		return c
	}
}
