package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

import "math"

type shapeSidesQuantity int

const SidesSquare shapeSidesQuantity = 4
const SidesTriangle shapeSidesQuantity = 3
const SidesCircle shapeSidesQuantity = 0

func CalcSquare(sideLen float64, sidesNum shapeSidesQuantity) float64 {
	var result float64

	if sidesNum == SidesSquare {
		result = math.Pow(sideLen, 2)
	}

	return result
}
