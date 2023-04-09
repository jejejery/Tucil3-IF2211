package structs

type MapValue struct {
	name string;
	coordinate Point;
}

func createMapValue(mapValue *MapValue, name string, point Point) {
	mapValue.name = name;
	mapValue.coordinate = point;
}

func getName(mapValue MapValue) string {
	return mapValue.name;
}

func getCoordinate(mapValue MapValue) Point {
	return mapValue.coordinate;
}
