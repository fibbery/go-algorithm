package geohash

import "bytes"

const (
	BASE32               = "0123456789bcdefghjkmnpqrstuvwxyz"
	MaxLatitude  float64 = 90
	MinLatitude  float64 = -90
	MaxLongitude float64 = 180
	MinLongitude float64 = -180
)

var (
	base32 = []byte(BASE32)
	bits   = []int{16, 8, 4, 2, 1,}
)

type Box struct {
	MinLat, MaxLat float64
	MinLng, MaxLng float64
}

func (b *Box) width() float64 {
	return b.MaxLng - b.MinLng
}

func (b *Box) height() float64 {
	return b.MaxLat - b.MinLat
}

func Encode(latitude, longitude float64, precision int) (string, *Box) {
	var geohash bytes.Buffer
	var mid float64
	minLati, maxLati, minLng, maxLng := MinLatitude, MaxLatitude, MinLongitude, MaxLongitude
	chunk, index, length, isEven := 0, 0, 0, true
	for length < precision {
		// set longitude when position is even , set latitude when position is odd, the position begin from 0
		if isEven {
			if mid = (minLng + maxLng) / 2; mid < longitude {
				chunk |= bits[index]
				minLng = mid
			} else {
				maxLng = mid
			}
		} else {
			if mid = (minLati + maxLati) / 2; mid < latitude {
				chunk |= bits[index]
				minLati = mid
			} else {
				maxLati = mid
			}
		}

		isEven = !isEven
		if index < 4 {
			index++
		} else {
			geohash.WriteByte(base32[chunk])
			length, chunk, index = length+1, 0, 0
		}
	}

	area := &Box{
		MinLat: minLati,
		MaxLat: maxLati,
		MinLng: minLng,
		MaxLng: maxLng,
	}

	return geohash.String(), area

}

func GetNeighbors(latitude, longitude float64, precision int) []string {
	center, area := Encode(latitude, longitude, precision)

	latUnit := (area.MaxLat - area.MinLat) / 2
	lngUnit := (area.MaxLng - area.MinLng) / 2

	// get neighbors area on up/down/left/right
	north, _ := Encode(area.MaxLat+latUnit, area.MinLng+lngUnit, precision)
	south, _ := Encode(area.MinLat-latUnit, area.MinLng+lngUnit, precision)
	west, _ := Encode(area.MinLat+latUnit, area.MinLng-lngUnit, precision)
	east, _ := Encode(area.MinLat+latUnit, area.MaxLng+lngUnit, precision)

	// get corner
	northWest, _ := Encode(area.MaxLat+latUnit, area.MinLng-lngUnit, precision)
	southWest, _ := Encode(area.MinLat-latUnit, area.MinLng-lngUnit, precision)
	northEast, _ := Encode(area.MaxLat+latUnit, area.MaxLng+lngUnit, precision)
	southEast, _ := Encode(area.MinLat-latUnit, area.MaxLng+lngUnit, precision)

	return []string{northWest, north, northEast, west, center, east, southWest, south, southEast,}
}
