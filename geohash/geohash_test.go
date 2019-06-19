package geohash

import (
	"testing"
)

func TestGeoHash(t *testing.T) {
	latitude, longitude := 31.1932993, 121.43960190000007
	hash, area := Encode(latitude, longitude, 8)
	t.Logf("this area geohash values is %s, and area is %v", hash, area)
	areas := GetNeighbors(latitude, longitude, 7)
	t.Logf("this area neighbors are %v", areas)
}

