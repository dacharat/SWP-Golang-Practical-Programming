package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	gas_pedal      uint16 //min: 0,      max: 65535
	brake_pedal    uint16 //min: 0,      max: 65535
	steering_wheel int16  //min: -32768  max: 32768
	top_speed_kmh  float64
}

func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax)
}

func (c *car) mph() float64 {
	c.top_speed_kmh = 500
	return float64(c.gas_pedal) * (c.top_speed_kmh / kmh_multiple / usixteenbitmax)
}

//modify the struct itself via pointer type
func (c *car) new_top_speed(newspeed float64) {
	c.top_speed_kmh = newspeed
}

func newer_top_speed(c car, speed float64) car {
	c.top_speed_kmh = speed
	return c
}

func main() {
	a_car := car{gas_pedal: 65000, brake_pedal: 0, steering_wheel: 12562, top_speed_kmh: 225.0}
	// a_car := car{22314, 0, 12562, 225.0}
	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
	// a_car.new_top_speed(500)
	// fmt.Println(a_car.kmh())
	// fmt.Println(a_car.mph())

	//a_car.new_top_speed(500)
	a_car = newer_top_speed(a_car, 500)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
}
