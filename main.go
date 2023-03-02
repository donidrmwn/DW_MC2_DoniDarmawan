package main

import "fmt"

type Profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

func main() {
	profile := MakeProfile("Naruto")
	fmt.Println("Name : ", profile.Name)
	fmt.Println("Health : ", profile.Health)
	fmt.Println("Power : ", profile.Power)
	fmt.Println("Exp : ", profile.Exp)
	fmt.Println("===HEROES POWER UP===")
	powerUp := PowerUp(profile, 4)
	fmt.Println("Name : ", powerUp.Name)
	fmt.Println("Health : ", powerUp.Health)
	fmt.Println("Power : ", powerUp.Power)
	fmt.Println("Exp : ", powerUp.Exp)

}

func MakeProfile(name string) Profile {
	profile := Profile{}
	profile.Name = name
	if name == "Sasuke" {
		profile.Health = 200
		profile.Power = 100
		profile.Exp = 0
	} else if name == "Goku" {
		profile.Health = 400
		profile.Power = 300
		profile.Exp = 100
	} else if name == "Naruto" {
		profile.Health = 150
		profile.Power = 200
		profile.Exp = 50
	}
	return profile
}

func PowerUp(profile Profile, multiplier int) Profile {
	profile.Health = profile.Health + (profile.Health * multiplier)
	profile.Power = profile.Power + (profile.Power * multiplier)
	profile.Exp = profile.Exp + (profile.Exp * multiplier)
	return profile
}

// var profile = Profile{}

// func main() {
// 	makeProfile("Sasuke")
// 	powerUp(2)
// 	fmt.Println(profile)
// }

// func makeProfile(name string) {

// 	profile.Name = name
// 	if name == "Sasuke" {
// 		profile.Health = 200
// 		profile.Power = 100
// 		profile.Exp = 0
// 	} else if name == "Goku" {
// 		profile.Health = 400
// 		profile.PowePrintln
// 		profile.Exp = 100
// 	} else if name == "Naruto" {
// 		profile.Health = 150
// 		profile.Power = 200
// 		profile.Exp = 50
// 	}

// }

// func powerUp(multiplier int) {
// 	profile.Health = profile.Health * multiplier
// 	profile.Power = profile.Power * multiplier
// 	profile.Exp = profile.Exp * multiplier

// }
