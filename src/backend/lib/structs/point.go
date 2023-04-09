package structs
import (
	"math"
)

type Point struct {
	x float64;
	y float64;
}

func createPoint(point *Point, x float64, y float64) {
	point.x = x;
	point.y = y;
}

func getX (point Point) float64 {
	return point.x;
}

func getY (point Point) float64 {
	return point.y;
}

func getDistance (p1 Point, p2 Point) float64 {
	return math.Sqrt(math.Pow(p1.x - p2.x, 2) + math.Pow(p1.y - p2.y, 2))
}
