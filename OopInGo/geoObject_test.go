package OopInGo

import "testing"

func TestGeoObject(t *testing.T) {
	var geoObject Painter
	geoObject = GeoObject{1, Position{2, 3}}
	geoObject.Paint()
	geoObject = Circle{GeoObject{3, Position{4, 5}}, 6}
	geoObject.Paint()
	geoObject = Rectangle{GeoObject{6, Position{7, 8}}, 9, 10}
	geoObject.Paint()
	geoObject = Triangle{GeoObject{10, Position{11, 12}},
		Position{13, 14}, Position{15, 16}, Position{17, 18}}
	geoObject.Paint()
}
