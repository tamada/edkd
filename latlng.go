package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		x1 = 139.74472
		x2 = 140.09111
		y1 = 35.65500
		y2 = 36.10056
		pi = math.Pi
	)

	x1 = (pi * x1) / 180
	x2 = (pi * x2) / 180
	y1 = (pi * y1) / 180
	y2 = (pi * y2) / 180

	var (
		abx = math.Abs(x1 - x2)
		aby = math.Abs(y1 - y2)
		dis = 0.0
	)

	dis = math.Pow(math.Sin(aby/2), 2) + math.Cos(y1)*math.Cos(y2)*math.Pow(math.Sin(abx/2), 2)
	dis = 2 * math.Asin(math.Sqrt(dis)) * 6371.007177356707

	fmt.Println(dis)

	var (
		p   = (y1 + y2) / 2
		rx  = 6378.137000000
		ry  = 6356.752314245
		e   = math.Sqrt((math.Pow(rx, 2) - math.Pow(ry, 2)) / math.Pow(rx, 2))
		w   = math.Sqrt(1 - math.Pow(e, 2)*math.Pow(math.Sin(p), 2))
		m   = rx * (1 - math.Pow(e, 2)) / math.Pow(w, 3)
		n   = rx / w
		Dis = math.Sqrt(math.Pow(aby*m, 2) + math.Pow(abx*n*math.Cos(p), 2))
	)
	fmt.Println(Dis)
}
