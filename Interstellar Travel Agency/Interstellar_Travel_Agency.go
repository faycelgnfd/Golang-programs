package main

import "fmt"

// function to print the amount of fuel left
func fuelGauge(fuel int) {
	fmt.Println("Amount of fuel left:", fuel, "liters")
}

// fucntion that calculates the amount of fuel needed to travel to a planet
func calculateFuel(planet string) int {

	if planet == "Venus" || planet == "venus" {
		return 300000
	} else if planet == "Mercury" || planet == "mercury" {
		return 500000
	} else if planet == "Mars" || planet == "mars" {
		return 700000
	} else {
		return 0
	}
}

// function that prints a greeting statement to a planet
func greetPlanet(planet string) {
	fmt.Println("Welcome dear traveler to", planet)
}

// fucntion that prints error when fuel is not available
func cantFly() {
	fmt.Println("Sorry we don't have the required amount of fuel to fly there")
}

// function that make the fly to a planet and return the fuel remained from the flight
func flyToPlanet(planet string, fuel int) int {
	var fuelRemaining, fuelCost int
	fuelRemaining = fuel

	fuelCost = calculateFuel(planet)

	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	} else {
		cantFly()
	}

	return fuelRemaining
}

func main() {
	var fuel = 1000000
	var planetChoice = "Venus"
	fuel = flyToPlanet(planetChoice, fuel)
	fuelGauge(fuel)
}
