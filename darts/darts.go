package darts

//Score func formula (x - x0)^2 + (y - y0)^2 <= R^2
func Score(x float64, y float64) int {
	point := x * x + y * y
	if point <= 1 {
		return 10
	} else if point <= 25 {
		return 5
	} else if point <= 100 {
		return 1
	}
	return 0
}