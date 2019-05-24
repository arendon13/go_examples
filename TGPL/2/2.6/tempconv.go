package tempconv

import (
	"fmt"
)

// Celcius represents celcius temperatures
type Celcius float64

// Fahrenheit represents fahrenheit temperatures
type Fahrenheit float64

const (
	// AbsoluteZeroC is the abs zero of Celcius
	AbsoluteZeroC Celcius = -273.15
	// FreezingC is the freezing point of Celcius
	FreezingC Celcius = 0
	// BoilingC is the boiling point of Celcius
	BoilingC Celcius = 100
)

func (c Celcius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
