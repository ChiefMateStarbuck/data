package main

import (
	"log"
	"math/rand"

	"github.com/thanhpk/randstr"
)

type Ant struct {
	name      string
	color     string
	isWarrior bool
	age       int
}

// AoS = Array of Structs
func AoS(size int) []Ant {
	ants := []Ant{}
	ant := Ant{}

	for i := 0; i < size; i++ {
		ant.name = randstr.Hex(16)
		ant.color = randstr.Hex(16)
		ant.isWarrior = randBool()
		ant.age = i
		ants = append(ants, ant)
	}

	return ants
}

// Counting all the warriors
func CountAoS(ants []Ant) {
	warriorsCount := 0

	for _, ant := range ants {
		if ant.isWarrior {
			warriorsCount++
		}
	}
}

type AntColony struct {
	size     int
	names    []string
	colors   []string
	ages     []int
	warriors []bool
}

// SoA = Stuct of Arrays
func SoA(size int) AntColony {
	ants := AntColony{}

	for i := 0; i < size; i++ {
		ants.names = append(ants.names, randstr.Hex(16))
		ants.colors = append(ants.colors, randstr.Hex(16))
		ants.ages = append(ants.ages, i)
		ants.warriors = append(ants.warriors, randBool())
	}
	ants.size = size

	return ants
}

// Counting all the warriors
func CountSoA(ants AntColony) {
	warriorsCount := 0

	for i := 0; i < ants.size; i++ {
		if ants.warriors[i] {
			warriorsCount++
		}
	}
}

func main() {
	log.Println("hahaha")
}

func randBool() bool {
	return rand.Intn(2) == 0
}
