package loc

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// 坐标-经纬度
type Coord struct {
	Lat float64
	Lng float64
}

func (origin Coord) Distance(q Coord) float64 {
	pos1, pos2 := degPos(origin.Lat, origin.Lng), degPos(q.Lat, q.Lng)
	return hsDist(pos1, pos2)
}

func (origin Coord) InChina() bool {
	return (origin.Lat > 3.85 && origin.Lat < 53.56) && (origin.Lng > 73.65 && origin.Lng < 135.05)
}

func (c *Coord) From(latlng string) (err error) {
	bs := strings.SplitN(latlng, ",", 2)
	if len(bs) != 2 {
		return errors.New("invail coord")
	}
	c.Lat, _ = strconv.ParseFloat(bs[0], 64)
	c.Lng, _ = strconv.ParseFloat(bs[1], 64)
	return nil
}

func (c *Coord) FromWGS84(latlng string) error {
	bs := strings.SplitN(latlng, ",", 2)
	if len(bs) != 2 {
		return errors.New("invail coord")
	}
	lat, err := strconv.ParseFloat(bs[0], 64)
	if err != nil {
		return err
	}
	lng, err := strconv.ParseFloat(bs[1], 64)
	if err != nil {
		return err
	}

	c.Lat, c.Lng = WGS84_To_Gcj02(lat, lng)
	return nil
}

func (origin Coord) String() string {
	return fmt.Sprintf("lat = %f, lng = %f", origin.Lat, origin.Lng)
}

func China() Coord {
	return Coord{Lat: 39.92, Lng: 116.46}
}

// 中国的经纬度范围大约为：纬度3.86~53.55，经度73.66~135.05
// 北京行政中心的纬度为39.92，经度为116.46

var CHINA = China()
