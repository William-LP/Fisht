package main

type ClassName int

const (
	MAGE ClassName = iota + 1
	WARRIOR
)

type Class struct {
	ClassName ClassName
	Spells    []Spell
}

func (c ClassName) ToString() string {
	switch c {
	case MAGE:
		return "Mage"
	case WARRIOR:
		return "Warrior"
	}
	return "unknown"
}

type Hero struct {
	HealthPoint,
	MagicPoint int
	Class Class
}

func NewBot(class ClassName, name string, game *Game) {
	h := NewHero(class)
	p := new(Player)
	p.Name = name
	p.Hero = *h
	game.Players = append(game.Players, *p)
}

func NewHero(class ClassName) *Hero {
	h := new(Hero)
	var Mage = Class{ClassName: MAGE, Spells: []Spell{FIREBALL, LIGHTNING}}
	var Warrior = Class{ClassName: WARRIOR, Spells: []Spell{SWORDSTRIKE, HEADBUT}}

	bonuses := modifyPoints(class)
	h.HealthPoint = bonuses.HealthPoint
	h.MagicPoint = bonuses.MagicPoint
	switch class.ToString() {
	case "Mage":
		h.Class = Mage

	case "Warrior":
		h.Class = Warrior
	}

	return h
}

func (h *Hero) reduceHealthPoint(value int) {
	h.HealthPoint = h.HealthPoint - value
}

func (h *Hero) reduceMagicPoint(value int) {
	h.MagicPoint = h.MagicPoint - value

}
