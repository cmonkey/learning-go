package formulas

import "math"

func Sqrt(x float64) float64{
    z := 0.0
    for i := 0; i < 1000; i++{
        z -= (z * z - x) / (2 * x)
    }

    return z
}

func Circum(d float64) float64{
    return math.Pi * d
}

func Surface(d float64) float64{
    return math.Pow(d, 2) * math.Pi / 2
}
