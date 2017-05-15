package libs

import (
	"math"
)

const EarthRadius float64 = 6378137 // 6378137

// EarthDistance 返回值的单位为米
func EarthDistance(lat1, lng1, lat2, lng2 float64) float64 {
	lat1 = rad(lat1)
	lng1 = rad(lng1)
	lat2 = rad(lat1)
	lng2 = rad(lng2)
	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return dist * EarthRadius
}

// rad 装换器
func rad(param float64) (result float64) {
	result = param * math.Pi / 180.0
	return
}
