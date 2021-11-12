package game

import (
	"fmt"
	"math/rand"
)

func remove(s *[]Hero, i int) []Hero {
	(*s)[i] = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return *s
}

func toFight(heroes *[]Hero, downstream, downstream2 chan Hero) {
	if len(*heroes) == 0 || len(*heroes) == 1 {
		return
	}

	first := 0 + rand.Intn(len(*heroes))
	downstream <- (*heroes)[first]
	remove(heroes, first)

	second := 0 + rand.Intn(len(*heroes))
	downstream2 <- (*heroes)[second]
	remove(heroes, second)
}

func makeFight(heroes *[]Hero, upstream, upstream2 chan Hero) {
	g := 0
	m := 0
	for v := range upstream {
		for p := range upstream2 {
			for {
				if v.amountStamina() > 20 {
					v.SecondSkill(p)
				} else if v.amountStamina() > 10 {
					v.FirstSkill(p)
				} else {
					choose := 1 + rand.Intn(2)
					if choose == 1 {
						fmt.Println(v, "Attack")
						v.Attack(p)
					} else {
						fmt.Println(v, "Defend")
						v.Defend()
					}
				}
				if p.Health() <= 0 {
					fmt.Println("In this battle winner ", v)
					fmt.Println("lose ", p)
					*heroes = append(*heroes, v)
					if g < len(*heroes)/2 {
						g++
					} else if m < len(*heroes)/2 {
						m++
					} else {
						g = 0
						m = 0
						g++
					}
					return
				}
				if p.amountStamina() > 20 {
					p.SecondSkill(v)
				} else if p.amountStamina() > 10 {
					p.FirstSkill(v)
				} else {
					choose := 1 + rand.Intn(2)
					if choose == 1 {
						fmt.Println(p, "Attack")
						p.Attack(v)
					} else {
						fmt.Println(p, "Defend")
						p.Defend()
					}
				}
				if v.Health() <= 0 {
					fmt.Println("In this battle winner ", p)
					fmt.Println("lose ", v)
					*heroes = append(*heroes, p)
					if g < len(*heroes)/2 {
						g++
					} else if m < len(*heroes)/2 {
						m++
					} else {
						g = 0
						m = 0
						g++
					}
					return
				}
			}
		}
	}
}
