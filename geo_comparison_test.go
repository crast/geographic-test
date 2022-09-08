package geo_comparison

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	g2 "github.com/tidwall/geodesic"
	g1 "github.com/xeonx/geodesic"
	"github.com/xeonx/geographic"
)

var epsilon = 1e-15
var epsilon2 = 1e-14

func compareInverse(t *testing.T, desc string, lat1, lng1, lat2, lng2 float64) {
	t.Run(desc, func(t *testing.T) {
		pt1 := geographic.Point{LatitudeDeg: lat1, LongitudeDeg: lng1}
		pt2 := geographic.Point{LatitudeDeg: lat2, LongitudeDeg: lng2}
		s12a, rawaz1a, rawaz2a := g1.WGS84.Inverse(pt1, pt2)

		var s12b, rawaz1b, rawaz2b float64
		g2.WGS84.Inverse(lat1, lng1, lat2, lng2, &s12b, &rawaz1b, &rawaz2b)

		assert.InEpsilon(t, s12a, s12b, epsilon, "s12(meters)")
		assert.InEpsilon(t, rawaz1a, rawaz1b, epsilon, "rawaz1")
		assert.InEpsilon(t, rawaz2a, rawaz2b, epsilon, "rawaz2")

	})
}

func TestInverse(t *testing.T) {
	/*coordJFK := NewCoordinateFromLatLng(40.64, -73.77)
	coordLHR := NewCoordinateFromLatLng(51.47, 0.46)
	coordLAX := NewCoordinateFromLatLng(34.40, -118.40)
	coordSYD := NewCoordinateFromLatLng(-33.52, 150.90)*/

	compareInverse(t, "JFKLHR", 40.64, -73.77, 51.47, 0.46)
	compareInverse(t, "LHRLAX", 51.47, 0.46, 34.40, -118.40)
	compareInverse(t, "zero-one", 0, 0, 1, 1)
}

var nmToMeters = 1852.00

func TestDirect(t *testing.T) {
	t.Run("JFK", func(t *testing.T) {
		az := 74.39
		for i := 1; i < 2800; i++ {
			t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
				lat1, lng1 := 40.64, -73.77
				pt1 := geographic.Point{LatitudeDeg: lat1, LongitudeDeg: lng1}
				pt2a, az2a := g1.WGS84.Direct(pt1, az, float64(i)*nmToMeters)

				var lat2b, lng2b, az2b float64
				g2.WGS84.Direct(lat1, lng1, az, float64(i)*nmToMeters, &lat2b, &lng2b, &az2b)

				assert.InEpsilon(t, pt2a.LatitudeDeg, lat2b, epsilon2, "lat2")
				assert.InEpsilon(t, pt2a.LongitudeDeg, lng2b, epsilon2, "lng2")
				assert.InEpsilon(t, az2a, az2b, epsilon, "az2")
			})
		}
	})

}
