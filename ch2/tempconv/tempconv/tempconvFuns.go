package tempconv

func FToC(f Fehrenheit) (c Celsius) {
	return Celsius((f - 32) * 5 / 9)
}

func CToF(c Celsius) (f Fehrenheit) {
	return Fehrenheit(c*9/5 + 32)
}
