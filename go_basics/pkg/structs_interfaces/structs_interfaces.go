package structs_interfaces

type GasEngine struct {
	Mpg     uint8
	Gallons uint8
}

func (e GasEngine) milesLeft() uint8 {
	return e.Gallons * e.Mpg
}

type ElectricEngine struct {
	Mpkwh uint8
	Kwh   uint8
}

func (e ElectricEngine) milesLeft() uint8 {
	return e.Kwh * e.Mpkwh
}

type engine interface {
	milesLeft() uint8
}

func CanMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		println("You can make it there!")
	} else {
		println("Needs to fuel it up first!")
	}
}
