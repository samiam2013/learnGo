package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Car has a mpg, a top speed and fuel capacity
type Car struct {
	MPG     float64
	TopMPH  float64
	FuelGal float64
	FuelCap float64
	PosMi   float64
}

func NewCar(mpg, topMPH, fuel float64) Car {
	return Car{
		MPG:     mpg,
		TopMPH:  topMPH,
		FuelGal: fuel,
		FuelCap: fuel,
		PosMi:   0.0,
	}
}

type finish struct {
	car  Car
	time time.Duration
}

const RaceLength = 2906 // miles

func main() {
	a := NewRandomCar()
	b := NewRandomCar()

	finishC := make(chan finish, 1)

	go race(a, finishC)
	go race(b, finishC)

	for i := 0; i < 2; i++ {
		f := <-finishC
		fmt.Printf("Car %+v took %s\n", f.car, f.time)
	}

}

func NewRandomCar() Car {
	// random mpg 20-40, random top speed 110 - 160, random fuelGal 12-40
	mpg := float64(20 + rand.Int63n(20))
	top := float64(110 + rand.Int63n(50))
	fuel := float64(12 + rand.Int63n(28))

	return NewCar(mpg, top, fuel)
}

func race(c Car, fC chan finish) {
	var t time.Duration
	gallonsPerMile := 1 / c.MPG
	loopTime := time.Millisecond * 250
	for {
		// drive for one minute at top speed
		t += loopTime
		fh := float64(loopTime.Nanoseconds()) / float64(time.Hour.Nanoseconds()) // = distance in miles
		// fmt.Println(fh, " fraction of an hour travelled")
		d := c.TopMPH * fh
		// fmt.Println("travelled", d, "miles")
		c.PosMi += d
		c.FuelGal -= (gallonsPerMile * d)

		// TODO calculate distance left and factor into decisions

		// if fuel < 1 gallon
		if c.FuelGal < 1.00 {
			// fmt.Println("refueling")
			t += (20 * time.Minute)
			c.FuelGal = c.FuelCap
		}

		// if you've driven past the finish line send finish signal and exit
		if c.PosMi >= RaceLength {
			fC <- finish{car: c, time: t}
			return
		}
	}

}
