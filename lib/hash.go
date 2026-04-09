package lib

import (
	"fmt"
	"math"
	"math/big"
)

// floatToUint returns an 128-bit big int representation of float x.
func floatToUint(x float64) big.Int {
	// Atsikratome minuso ženklo
	x = math.Abs(x)

	// Normalizuojame modulį į [1, 10), kad pirmasis išgaunamas skaitmuo būtų vienetų vietoje
	// Pvz., 12.3 -> 1.23, 0.123 -> 1.23
	x = x / math.Pow10(int(math.Floor(math.Log10(x))))

	// Konvertuojame float skaiciu i hex reiksme
	// Maziausia ir didziausia leidziamos isvesties reiksmes yra 0x10000... ir 0xFFFFF...
	// Desimtaineje sistemoje tai lygu 2.12 x 10^37 ir 34.02 x 10^37
	// Tokiu veiksmu isitikiname, kad skaicius kris isvesties ribose
	x = x * 3 * math.Pow10(37-16)
	fmt.Printf("float64 value of x is: %v\n", x)

	s := fmt.Sprintf("%.0f", x)
	var output big.Int
	output.SetString(s, 10)

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

	hashedX := floatToUint(x)
	hashedY := floatToUint(y)

	return fmt.Sprintf("%x %x", &hashedX, &hashedY)
}
