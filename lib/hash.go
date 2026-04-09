package lib

import (
	"math"
)

func Hash(inputStr string) (hex64 string, err error) {
	// Pradėti nuo koordinačių (0,0)
	var x, y float64 = 0, 0

	// Kiekvienas simbolis prideda 1 ilgio "švytuoklę" (pendulum) tam tikru kampu, kur kampas = a (radianais).
	for _, v := range inputStr {
		x += math.Cos(float64(v))
		y += math.Sin(float64(v))
	}

	// log.Printf("x = %v, y = %v", x, y)

	// placeholder
	hex64 = "d716c0c1cc5ebf520fb000b6b274c576f9a5984a22fa7d6f3f314f902357c061"
	return hex64, err
}
