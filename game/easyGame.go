package main

import (
	"fmt"
	"math/rand"
)

type Hero interface {
	Attack(hero Hero)
	Defend()
	FirstSkill(hero Hero)
	SecondSkill(hero Hero)
	Health() int
	setHealth(int)
	setDefense(int)
}

type Warrior struct {
	Healths int
	Power   int
	Defense int
	Rage    int
}
type Mage struct {
	Healths int
	Power   int
	Defense int
	Mana    int
}
type Hunter struct {
	Healths int
	Power   int
	Defense int
	Energy  int
}

func (war *Warrior) Attack(hero Hero) {
	war.Rage += 10
	damage := war.Power + (10 + rand.Intn(11)) - war.Defense
	fmt.Println("You have dealt ", damage, " damage")
	hero.setHealth(damage)
	hero.setDefense(damage)
}
func (war *Warrior) Defend() {
	war.Rage += 5
	war.Defense = war.Defense + (8 + rand.Intn(9))
}
func (war *Warrior) FirstSkill(hero Hero) {
	if war.Rage < 10 {
		fmt.Println("You're not angry enough! Be angrier!!!Arrrrrr")
		return
	}
	fmt.Println("YOU want to be like your hero Dovakin, so you scream at the enemy. You're stunning him!")
	war.Power += 1
	war.Rage -= 10
	damage := war.Power + 5 + rand.Intn(6) - war.Defense
	fmt.Println("You have dealt ", damage, " damage")
	hero.setHealth(damage)
	hero.setDefense(damage)
}
func (war *Warrior) SecondSkill(hero Hero) {
	if war.Rage < 20 {
		fmt.Println("You're not angry enough! Be angrier!!!Arrrrrr")
		return
	}
	fmt.Println("You take out the peasant pitchfork that you stole from your father on the farm and rush with it to the enemy, imagining that you are a Warsong")
	war.Power += 3
	war.Rage -= 20
	damage := war.Power + 20 + rand.Intn(21) - war.Defense
	fmt.Println("You have dealt ", damage, " damage")
	hero.setHealth(damage)
	hero.setDefense(damage)

}
func (war *Warrior) Health() int {
	return war.Healths
}
func (war *Warrior) setHealth(points int) {
	war.Healths = war.Healths - points
}
func (war *Warrior) setDefense(point int) {
	war.Defense = war.Defense - point
}

func (mag *Mage) Attack(hero Hero) {
	mag.Mana += 3
	damage := mag.Power + 7 + rand.Intn(8) - mag.Defense
	fmt.Println("You have dealt ", damage, " damage")
	hero.setHealth(damage)
	hero.setDefense(damage)

}
func (mag *Mage) Defend() {
	mag.Mana += 12
	mag.Defense = mag.Defense + (13 + rand.Intn(14))
}
func (mag *Mage) FirstSkill(hero Hero) {
	if mag.Mana < 10 {
		fmt.Println("You have too little mana. Draw it harder!")
		return
	}
	fmt.Println("You turn the enemy into the animal that he is most afraid of. And what is your surprise when the enemy became a sheep.The enemy is stunned.")
	mag.Power += 4
	mag.Mana -= 10
	damage := mag.Power + (4 + rand.Intn(5)) - mag.Defense
	fmt.Println("You have dealt ", damage, " damage")
	hero.setHealth(damage)
	hero.setDefense(damage)

}
func (mag *Mage) SecondSkill(hero Hero) {
	if mag.Mana < 20 {
		fmt.Println("You have too little mana. Draw it harder!")
		return
	}
	fmt.Println("YYou summon a huge fireball as if you are from the Uchiha clan and throw it at the enemy")
	mag.Mana -= 20
	mag.Power += 3
	damage := mag.Power + (22 + rand.Intn(23)) - mag.Defense
	fmt.Println("You have dealt ", damage, " damage")
	hero.setHealth(damage)
	hero.setDefense(damage)

}
func (mag *Mage) Health() int {
	return mag.Healths
}
func (mag *Mage) setHealth(points int) {
	mag.Healths = mag.Healths - points
}
func (mag *Mage) setDefense(point int) {
	mag.Defense = mag.Defense - point
}

func (hun *Hunter) Attack(hero Hero) {
	hun.Energy += 12
	damage := hun.Power + (11 + rand.Intn(12)) - hun.Defense
	fmt.Println("You have dealt ", damage, " damage")
	hero.setHealth(damage)
	hero.setDefense(damage)
}
func (hun *Hunter) Defend() {
	hun.Energy += 7
	hun.Defense += 10 + rand.Intn(11)
}
func (hun *Hunter) FirstSkill(hero Hero) {
	if hun.Energy < 10 {
		fmt.Println("You have too little Energy, you need to pet some animal urgently")
		return
	}
	fmt.Println("You send your pet to walk between the enemy's legs and rub against them.You stunned him.")
	hun.Power += 2
	hun.Energy -= 10
	damage := hun.Power + (7 + rand.Intn(8)) - hun.Defense
	fmt.Println("You have dealt ", damage, " damage")
	hero.setHealth(damage)
	hero.setDefense(damage)
}
func (hun *Hunter) SecondSkill(hero Hero) {
	if hun.Energy < 20 {
		fmt.Println("You have too little Energy, you need to pet some animal urgently")
		return
	}
	fmt.Println("You imagine that you are a Rexar, and your cat is a bear. And gather all your strength in this shot!")
	hun.Energy -= 20
	hun.Power += 3
	damage := hun.Power + (25 + rand.Intn(26)) - hun.Defense
	fmt.Println("You have dealt ", damage, " damage")
	hero.setHealth(damage)
	hero.setDefense(damage)
}
func (hun *Hunter) Health() int {
	return hun.Healths
}
func (hun *Hunter) setHealth(points int) {
	hun.Healths = hun.Healths - points
}
func (hun *Hunter) setDefense(point int) {
	hun.Defense = hun.Defense - point
}

func main() {
	var choiceOne, choiceTwo string
	fmt.Println("Choose your first hero: W - warrior, M - mage, H - hunter")
	fmt.Scan(&choiceOne)
	var firstHero, secondHero Hero
	switch choiceOne {
	case "W":
		firstHero = &Warrior{Healths: 100, Power: 8, Defense: 6, Rage: 0}
	case "M":
		firstHero = &Mage{Healths: 100, Power: 6, Defense: 8, Mana: 0}
	case "H":
		firstHero = &Hunter{Healths: 100, Power: 7, Defense: 7, Energy: 0}
	}
	fmt.Println("Choose your second hero: W - warrior, M - mage, H - hunter")
	fmt.Scan(&choiceTwo)
	switch choiceTwo {
	case "W":
		secondHero = &Warrior{Healths: 100, Power: 8, Defense: 6, Rage: 0}
	case "M":
		secondHero = &Mage{Healths: 100, Power: 6, Defense: 8, Mana: 0}
	case "H":
		secondHero = &Hunter{Healths: 100, Power: 7, Defense: 7, Energy: 0}
	}
	var choice int
	for firstHero.Health() != 0 && secondHero.Health() != 0 {
		fmt.Println("Choose your action: 1 - attack, 2 - defense, 3 - first skill (cost 10), 4 - second skill (cost 20), 9 - exit")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			firstHero.Attack(secondHero)
		case 2:
			firstHero.Defend()
		case 3:
			firstHero.FirstSkill(secondHero)
		case 4:
			firstHero.SecondSkill(secondHero)
		case 9:
			return
		}

		if secondHero.Health() <= 0 {
			fmt.Println("The first one won")
			return
		}
		fmt.Println("Choose your action: 1 - attack, 2 - defense, 3 - first skill (cost 10), 4 - second skill (cost 20), 9 - exit")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			secondHero.Attack(firstHero)
		case 2:
			secondHero.Defend()
		case 3:
			secondHero.FirstSkill(firstHero)
		case 4:
			secondHero.SecondSkill(firstHero)
		case 9:
			return
		}
		if firstHero.Health() <= 0 {
			fmt.Println("The second one won")
			return
		}
	}
}
