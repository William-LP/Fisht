package main

import (
	"fmt"
	"math/rand"
)

type Spell int

const (
	FIREBALL Spell = iota + 1
	LIGHTNING
	SWORDSTRIKE
	HEADBUT
)

func (s Spell) ToString() string {
	switch s {
	case FIREBALL:
		return "fireball"
	case LIGHTNING:
		return "lightning"
	case SWORDSTRIKE:
		return "swordstrike"
	case HEADBUT:
		return "headbut"
	}
	return "unknown"
}

type Range struct {
	Min, Max int
}

type SpellInfo struct {
	Range
	Cost int
}

var spellInfo = map[Spell]SpellInfo{
	FIREBALL:    {Range: Range{Min: 10, Max: 30}, Cost: 10},
	LIGHTNING:   {Range: Range{Min: 5, Max: 40}, Cost: 20},
	SWORDSTRIKE: {Range: Range{Min: 15, Max: 25}, Cost: 10},
	HEADBUT:     {Range: Range{Min: 10, Max: 35}, Cost: 15},
}

func (p *Player) attack(target *Player, spell Spell, game *Game) {

	info, ok := spellInfo[spell]
	if !ok {
		game.appendCombatLog(fmt.Sprintf("%v tried to cast an unknown spell on %v. This had no effect.", p.Name, target.Name), INFO)

	}

	if p.Hero.MagicPoint >= info.Cost {
		p.Hero.reduceMagicPoint(info.Cost)
		damage := rand.Intn(info.Max-info.Min+1) + info.Min
		target.Hero.reduceHealthPoint(damage)
		game.appendCombatLog(fmt.Sprintf("%v cast %v on %v for %v damages", p.Name, spell.ToString(), target.Name, damage), ATTACK)
	} else {
		game.appendCombatLog(fmt.Sprintf("%v tried to cast %v on %v but had not enought magic point. This had no effect.", p.Name, spell.ToString(), target.Name), INFO)
	}

}
