package lib

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

// normalize returns the float x shifted to the range [1, 10)
func normalize(x *float64) {
	// Normalizuojame skaiciu į [1, 10), kad pirmasis išgaunamas skaitmuo būtų vienetų vietoje
	// Pvz., 12.3 -> 1.23, 0.123 -> 1.23
	*x /= math.Pow(10, float64(math.Floor(math.Log10(*x))))
}

// floatToUint returns a hex-encoded string of the first 32 digits of x
func floatToUint(x float64) string {
	// Mažiausia ir didžiausia leidžiamos išvesties reikšmės yra 0x10000... ir 0xFFFFF...
	// Desimtaineje sistemoje tai lygu 2.12 x 10^37 ir 34.02 x 10^37

	// Atsikratome minuso ženklo
	x = math.Abs(x)

	// Normalizuojame skaiciu į [1, 10), kad pirmasis išgaunamas skaitmuo būtų vienetų vietoje
	normalize(&x)

	// Italpiname i riba tarp 2.12 ir 34.02
	x *= 3

	// Dedame skaitmenis i isvesties string
	output := ""
	for i := 0; i < 38; i++ {
		// Paimame pirmą skaitmenį (0..9)
		digit := uint64(x)
		// Pridedame skaitmenį prie išvesties string
		output += strconv.FormatUint(digit, 10)
		// Atsikratome dabartinio skaitmens ir perkeliame kitą į vienetų vietą
		x = (x - float64(digit)) * 10
	}

	return output
}

// PHA256 - self-made, performant hash algorithm.
// Returns a 256-bit hash of input in hex.
func PHA256(input []byte) string {
	// Pradėti nuo koordinačių (0,0)
	var x, y float64 = 0, 0

	// Kiekvienas simbolis prideda 1 ilgio "švytuoklę" (pendulum) tam tikru kampu, kur kampas = a (radianais).
	for _, v := range input {
		x += math.Cos(float64(v))
		y += math.Sin(float64(v))
	}

	hashedX := new(big.Int)
	hashedY := new(big.Int)

	hashedX.SetString(floatToUint(x), 10)
	hashedY.SetString(floatToUint(y), 10)

	return fmt.Sprintf("%x%x", hashedX, hashedY)
}
