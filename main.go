package main

import (
	"fmt"
)

type Profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

func MakeProfile(name string) Profile {
	if name == "Sasuke" {
		return Profile{
			Name:   name,
			Health: 200,
			Power:  100,
			Exp:    0,
		}
	}
	if name == "Goku" {
		return Profile{
			Name:   name,
			Health: 400,
			Power:  300,
			Exp:    100,
		}
	}
	if name == "Naruto" {
		return Profile{
			Name:   name,
			Health: 150,
			Power:  200,
			Exp:    50,
		}
	}
	return Profile{}
}

func PowerUp(profile Profile, multiplayer int) Profile {
	return Profile{
		Name:   profile.Name,
		Health: profile.Health * multiplayer,
		Power:  profile.Power * multiplayer,
		Exp:    profile.Exp * multiplayer,
	}
}

func main() {
	ProfileSasuke := MakeProfile("Sasuke")
	fmt.Println("===HEROES NAME===")
	fmt.Println("Name : ", ProfileSasuke.Name)
	fmt.Println("Health : ", ProfileSasuke.Health)
	fmt.Println("Power : ", ProfileSasuke.Power)
	fmt.Println("Exp : ", ProfileSasuke.Exp)

	PowerSasuke := PowerUp(ProfileSasuke, 3)
	fmt.Println("====HEROES POWER UP===")
	fmt.Println("Name : ", PowerSasuke.Name)
	fmt.Println("Health : ", PowerSasuke.Health)
	fmt.Println("Power : ", PowerSasuke.Power)
	fmt.Println("Exp : ", PowerSasuke.Exp)

	ProfileGoku := MakeProfile("Goku")
	fmt.Println("===HEROES NAME===")
	fmt.Println("Name : ", ProfileGoku.Name)
	fmt.Println("Health : ", ProfileGoku.Health)
	fmt.Println("Power : ", ProfileGoku.Power)
	fmt.Println("Exp : ", ProfileGoku.Exp)
	PowerGoku := PowerUp(ProfileGoku, 2)
	fmt.Println("====HEROES POWER UP===")
	fmt.Println("Name : ", PowerGoku.Name)
	fmt.Println("Health : ", PowerGoku.Health)
	fmt.Println("Power : ", PowerGoku.Power)
	fmt.Println("Exp : ", PowerGoku.Exp)

	ProfileNaruto := MakeProfile("Naruto")
	fmt.Println("===HEROES NAME===")
	fmt.Println("Name : ", ProfileNaruto.Name)
	fmt.Println("Health : ", ProfileNaruto.Health)
	fmt.Println("Power : ", ProfileNaruto.Power)
	fmt.Println("Exp : ", ProfileNaruto.Exp)
	PowerNaruto := PowerUp(ProfileNaruto, 4)
	fmt.Println("====HEROES POWER UP===")
	fmt.Println("Name : ", PowerNaruto.Name)
	fmt.Println("Health : ", PowerNaruto.Health)
	fmt.Println("Power : ", PowerNaruto.Power)
	fmt.Println("Exp : ", PowerNaruto.Exp)
}
