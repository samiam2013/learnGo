package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// Car has a mpg, a top speed and fuel capacity
type Car struct {
	MPG         float64
	TopMPH      float64
	FuelGal     float64
	FuelCap     float64
	PosMi       float64
	timeElapsed time.Duration
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

func NewRandomCar() Car {
	// random mpg 20-40, random top speed 110 - 160, random fuelGal 12-40
	//nolint this isn't a cryptographic use of rand
	var (mpg =  float64( 20 + rand.Int63n(20))
		top = float64(110 + rand.Int63n(50))
		fuel = float64(12 + rand.Int63n(28)))
	return NewCar(mpg, top, fuel)
}

func (c *Car) Refuel() {
	log.Println("refueling")
	if RaceLength-c.PosMi < c.MPG*1.0 {
		c.timeElapsed += 3 * time.Minute
		c.FuelGal += 1.2
		return
	}
	c.timeElapsed += 20 * time.Minute
	c.FuelGal = c.FuelCap
}

func (c *Car) Travel(d time.Duration) {
	gallonsPerMile := 1 / c.MPG
	c.timeElapsed += d
	hrNS := float64(time.Hour.Nanoseconds())
	dt := float64(d.Nanoseconds())
	fh := dt / hrNS // fraction in hours
	dis := c.TopMPH * fh
	c.PosMi += dis
	c.FuelGal -= gallonsPerMile * dis

	log.Println(fh, " fraction of an hour travelled")
	log.Println("travelled", d, "miles")
}

func (c *Car) String() string {
	return fmt.Sprintf("Car with top speed %.2f, MPG %.2f, fuel %.2f, "+
		"finished in %s", c.TopMPH, c.MPG, c.FuelGal, c.timeElapsed)
}

func race(c Car, length float64) {
	loopTime := time.Millisecond * 250
	for {
		// drive for one minute at top speed
		c.Travel(loopTime)
		// if fuel < 1 gallon
		if c.FuelGal < 1.00 {
			c.Refuel()
		}

		// if you've driven past the finish line send finish signal and exit
		if c.PosMi >= length {
			fmt.Printf("Car %s finished\n", c.String())
			return
		}
	}
}

const RaceLength = 2906 // miles

func main() {
	a := NewRandomCar()
	b := NewRandomCar()

	race(a, RaceLength)
	race(b, RaceLength)

}
