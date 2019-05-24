package tempconv

// CToF converts a Celcius temperature to Fahrenheit
func CToF(c Celcius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celcius
func FToC(f Fahrenheit) Celcius { return Celcius((f - 32) * 5 / 9) }
