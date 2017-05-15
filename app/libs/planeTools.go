package libs

import (
	"math"
)

var EarthRadius float64 = 6378.137

func GetDistance(startLat, startLng, endLat, endLng float64) {
	radLat1 := rad(startLat)
	radLat2 := rad(endLat)
	redStart := radLat1 - radLat2
	radEnd := rad(startLng) - rad(endLng)
	s := 2 * math.Asin(math.Sqrt(math.Pow(math.Sin(redStart/2), 2)+
		math.Cos(radLat1)*math.Cos(radLat2)*math.Pow(math.Sin(radEnd/2), 2)))
	s = s * EarthRadius
}

// rad 装换器
func rad(param float64) (result float64) {
	result = param * math.Pi / 180.0
	return
}
