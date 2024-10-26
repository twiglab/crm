package loc

/*
百度地图API	百度坐标
腾讯搜搜地图API	火星坐标
搜狐搜狗地图API	搜狗坐标
阿里云地图API	火星坐标
图吧MapBar地图API	图吧坐标
高德MapABC地图API	火星坐标
灵图51ditu地图API	火星坐标
*/

import (
	"math"
)

// GPSUtil is a utility class for GPS calculations.
// 小写方法是私有方法，大写方法是公有方法 可根据需要调整
type GPSUtil struct {
}

const (
	pi   = math.Pi                  // 3.1415926535897932384626 // 圆周率
	x_pi = math.Pi * 3000.0 / 180.0 // 圆周率对应的经纬度偏移
	a    = 6378245.0                // 长半轴
	ee   = 0.00669342162296594323   // 扁率
)

func transformLat(x, y float64) float64 {
	ret := -100.0 + 2.0*x + 3.0*y + 0.2*y*y + 0.1*x*y + 0.2*math.Sqrt(math.Abs(x))
	ret += (20.0*math.Sin(6.0*x*pi) + 20.0*math.Sin(2.0*x*pi)) * 2.0 / 3.0
	ret += (20.0*math.Sin(y*pi) + 40.0*math.Sin(y/3.0*pi)) * 2.0 / 3.0
	ret += (160.0*math.Sin(y/12.0*pi) + 320*math.Sin(y*pi/30.0)) * 2.0 / 3.0
	return ret
}

func (receiver *GPSUtil) transformLat(x, y float64) float64 {
	ret := -100.0 + 2.0*x + 3.0*y + 0.2*y*y + 0.1*x*y + 0.2*math.Sqrt(math.Abs(x))
	ret += (20.0*math.Sin(6.0*x*pi) + 20.0*math.Sin(2.0*x*pi)) * 2.0 / 3.0
	ret += (20.0*math.Sin(y*pi) + 40.0*math.Sin(y/3.0*pi)) * 2.0 / 3.0
	ret += (160.0*math.Sin(y/12.0*pi) + 320*math.Sin(y*pi/30.0)) * 2.0 / 3.0
	return ret
}

func transformlng(x, y float64) float64 {
	ret := 300.0 + x + 2.0*y + 0.1*x*x + 0.1*x*y + 0.1*math.Sqrt(math.Abs(x))
	ret += (20.0*math.Sin(6.0*x*pi) + 20.0*math.Sin(2.0*x*pi)) * 2.0 / 3.0
	ret += (20.0*math.Sin(x*pi) + 40.0*math.Sin(x/3.0*pi)) * 2.0 / 3.0
	ret += (150.0*math.Sin(x/12.0*pi) + 300.0*math.Sin(x/30.0*pi)) * 2.0 / 3.0
	return ret
}

func (receiver *GPSUtil) transformlng(x, y float64) float64 {
	ret := 300.0 + x + 2.0*y + 0.1*x*x + 0.1*x*y + 0.1*math.Sqrt(math.Abs(x))
	ret += (20.0*math.Sin(6.0*x*pi) + 20.0*math.Sin(2.0*x*pi)) * 2.0 / 3.0
	ret += (20.0*math.Sin(x*pi) + 40.0*math.Sin(x/3.0*pi)) * 2.0 / 3.0
	ret += (150.0*math.Sin(x/12.0*pi) + 300.0*math.Sin(x/30.0*pi)) * 2.0 / 3.0
	return ret
}

func outOfChina(lat, lng float64) bool {
	if lng < 72.004 || lng > 137.8347 {
		return true
	}
	if lat < 0.8293 || lat > 55.8271 {
		return true
	}
	return false
}

func (receiver *GPSUtil) outOfChina(lat, lng float64) bool {
	if lng < 72.004 || lng > 137.8347 {
		return true
	}
	if lat < 0.8293 || lat > 55.8271 {
		return true
	}
	return false
}

func transform(lat, lng float64) (float64, float64) {
	if outOfChina(lat, lng) {
		return 0.0, 0.0
	}
	dLat := transformLat(lng-105.0, lat-35.0)
	dlng := transformlng(lng-105.0, lat-35.0)
	radLat := lat / 180.0 * pi
	magic := math.Sin(radLat)
	magic = 1 - ee*magic*magic
	SqrtMagic := math.Sqrt(magic)
	dLat = (dLat * 180.0) / ((a * (1 - ee)) / (magic * SqrtMagic) * pi)
	dlng = (dlng * 180.0) / (a / SqrtMagic * math.Cos(radLat) * pi)
	mgLat := lat + dLat
	mglng := lng + dlng
	return mgLat, mglng
}

func (receiver *GPSUtil) transform(lat, lng float64) []float64 {
	if receiver.outOfChina(lat, lng) {
		return []float64{lat, lng}
	}
	dLat := receiver.transformLat(lng-105.0, lat-35.0)
	dlng := receiver.transformlng(lng-105.0, lat-35.0)
	radLat := lat / 180.0 * pi
	magic := math.Sin(radLat)
	magic = 1 - ee*magic*magic
	SqrtMagic := math.Sqrt(magic)
	dLat = (dLat * 180.0) / ((a * (1 - ee)) / (magic * SqrtMagic) * pi)
	dlng = (dlng * 180.0) / (a / SqrtMagic * math.Cos(radLat) * pi)
	mgLat := lat + dLat
	mglng := lng + dlng
	return []float64{mgLat, mglng}
}

// WGS84_To_Gcj02 84 to 火星坐标系 (GCJ-02) World Geodetic System ==> Mars Geodetic System
// @param lat
// @param lng
// @return
func WGS84_To_Gcj02(lat, lng float64) (float64, float64) {
	if outOfChina(lat, lng) {
		return 0.0, 0.0
	}
	dLat := transformLat(lng-105.0, lat-35.0)
	dlng := transformlng(lng-105.0, lat-35.0)
	radLat := lat / 180.0 * pi
	magic := math.Sin(radLat)
	magic = 1 - ee*magic*magic
	SqrtMagic := math.Sqrt(magic)
	dLat = (dLat * 180.0) / ((a * (1 - ee)) / (magic * SqrtMagic) * pi)
	dlng = (dlng * 180.0) / (a / SqrtMagic * math.Cos(radLat) * pi)
	mgLat := lat + dLat
	mglng := lng + dlng
	return mgLat, mglng
}

func (receiver *GPSUtil) WGS84_To_Gcj02(lat, lng float64) []float64 {
	if receiver.outOfChina(lat, lng) {
		return []float64{lat, lng}
	}
	dLat := receiver.transformLat(lng-105.0, lat-35.0)
	dlng := receiver.transformlng(lng-105.0, lat-35.0)
	radLat := lat / 180.0 * pi
	magic := math.Sin(radLat)
	magic = 1 - ee*magic*magic
	SqrtMagic := math.Sqrt(magic)
	dLat = (dLat * 180.0) / ((a * (1 - ee)) / (magic * SqrtMagic) * pi)
	dlng = (dlng * 180.0) / (a / SqrtMagic * math.Cos(radLat) * pi)
	mgLat := lat + dLat
	mglng := lng + dlng
	return []float64{mgLat, mglng}
}

// GCJ02_To_WGS84
// 火星坐标系 (GCJ-02) to WGS84
// @param lng
// @param lat
// @return
func (receiver *GPSUtil) GCJ02_To_WGS84(lat, lng float64) []float64 {
	gps := receiver.transform(lat, lng)
	lngtitude := lng*2 - gps[1]
	latitude := lat*2 - gps[0]
	return []float64{latitude, lngtitude}
}

/**
 * 火星坐标系 (GCJ-02) 与百度坐标系 (BD-09) 的转换算法 将 GCJ-02 坐标转换成 BD-09 坐标
 *
 * @param lat
 * @param lng
 */
func (receiver *GPSUtil) gcj02_To_Bd09(lat, lng float64) []float64 {
	x := lng
	y := lat
	z := math.Sqrt(x*x+y*y) + 0.00002*math.Sin(y*x_pi)
	theta := math.Atan2(y, x) + 0.000003*math.Cos(x*x_pi)
	templng := z*math.Cos(theta) + 0.0065
	tempLat := z*math.Sin(theta) + 0.006
	gps := []float64{tempLat, templng}
	return gps
}

/**
 * * 火星坐标系 (GCJ-02) 与百度坐标系 (BD-09) 的转换算法 * * 将 BD-09 坐标转换成GCJ-02 坐标 * * @param
 * bd_lat * @param bd_lng * @return
 */
func (receiver *GPSUtil) bd09_To_Gcj02(lat, lng float64) []float64 {
	x := lng - 0.0065
	y := lat - 0.006
	z := math.Sqrt(x*x+y*y) - 0.00002*math.Sin(y*x_pi)
	theta := math.Atan2(y, x) - 0.000003*math.Cos(x*x_pi)
	templng := z * math.Cos(theta)
	tempLat := z * math.Sin(theta)
	gps := []float64{tempLat, templng}
	return gps
}

/**将WGS84转为bd09
 * @param lat
 * @param lng
 * @return
 */
func (receiver *GPSUtil) WGS84_To_bd09(lat, lng float64) []float64 {
	gcj02 := receiver.WGS84_To_Gcj02(lat, lng)
	bd09 := receiver.gcj02_To_Bd09(gcj02[0], gcj02[1])
	return bd09
}

func (receiver *GPSUtil) bd09_To_WGS84(lat, lng float64) []float64 {
	gcj02 := receiver.bd09_To_Gcj02(lat, lng)
	WGS84 := receiver.GCJ02_To_WGS84(gcj02[0], gcj02[1])
	return WGS84
}
