package rectangle

import "errors"

//FindRectancularIntersection ...
func FindRectancularIntersection(rec1, rec2 map[string]int) (map[string]int,
	error) {

	leftX, width, err := findRangeOverlap(rec1["leftX"], rec1["width"],
		rec2["leftX"], rec2["width"])
	if err != nil {
		return map[string]int{}, err
	}
	bottomY, height, err := findRangeOverlap(rec1["bottomY"], rec1["height"],
		rec2["bottomY"], rec2["height"])
	if err != nil {
		return map[string]int{}, err
	}
	return map[string]int{"leftX": leftX, "bottomY": bottomY, "width": width,
		"height": height}, nil

}

func findRangeOverlap(point1, length1, point2, length2 int) (int, int,
	error) {
	//highest point = left/lower side overlap = x to the farthest right or
	//uppermost y
	highest := point1
	if point2 > highest {
		highest = point2
	}
	//lowest point = right/upper side overlap
	lowest := point1 + length1
	if point2+length2 < lowest {
		lowest = point2 + length2
	}

	//if highest > lowest, there is no overlap
	if highest >= lowest {
		return 0, 0, errors.New("No overlap")
	}

	width := lowest - highest
	return highest, width, nil

}
