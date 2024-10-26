package main

import (
	"fmt"

	"github.com/twiglab/crm/adv/latlng"
)

func main() {
	l := latlng.ToIndex(37.775938728915946, -122.41795063018799)
	fmt.Println(l)
}
