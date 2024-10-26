package latlng

import (
	"fmt"

	"github.com/uber/h3-go/v4"
)

type Index struct {
	Index6 string
	Index7 string
	Index8 string

	Lat, Lng float64
}

func ToIndex(lat, lng float64) Index {
	latLng := h3.NewLatLng(lat, lng)

	index6 := fmt.Sprint(latLng.Cell(6))
	index7 := fmt.Sprint(latLng.Cell(7))
	index8 := fmt.Sprint(latLng.Cell(8))

	return Index{
		Index6: index6,
		Index7: index7,
		Index8: index8,

		Lat: lat,
		Lng: lng,
	}
}
