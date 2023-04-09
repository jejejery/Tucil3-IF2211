package structs

type MapValue struct {
	name string;
	coordinate Point;
}

func CreateMapValue(mapValue *MapValue, name string, point Point) {
	mapValue.name = name;
	mapValue.coordinate = point;
}

func GetName(mapValue MapValue) string {
	return mapValue.name;
}

func GetCoordinate(mapValue MapValue) Point {
	return mapValue.coordinate;
}
