package main

import (
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/nexidian/gocliselect"
)

type Kind int

const (
	HUMAN Kind = iota + 1
	BOT
)

type Player struct {
	Id   int
	Kind Kind
	Name string
	Hero Hero
}

type LogType int

const (
	INFO LogType = iota + 1
	ATTACK
	HEAL
)

type Log struct {
	Value   string
	LogType LogType
}

type Game struct {
	Players      []Player
	ActivePlayer *Player
	CombatLog    []Log
}

func (g *Game) appendCombatLog(log string, logType LogType) {
	g.CombatLog = append(g.CombatLog, Log{Value: log, LogType: logType})
	for len(g.CombatLog) > 3+numBots*2 {
		g.CombatLog = g.CombatLog[1:]
	}
}

func NewGame() *Game {
	clearScreen()
	myFigure := figure.NewFigure("Fisht", "rectangles", true)
	myFigure.Print()

	g := new(Game)

	p := new(Player)
	p.Kind = HUMAN
	println()
	p.Name = getInput("What's your name (between 4 and 8 characters): ")

	if len(p.Name) < 4 {
		suffix := strings.Repeat("_", 4-len(p.Name))
		p.Name = fmt.Sprintf("%v%v", p.Name, suffix)
	}

	if len(p.Name) > 8 {
		p.Name = p.Name[:8]
	}
	fmt.Println()
	menu := gocliselect.NewMenu("Pick a class ")
	menu.AddItem("Mage", "mage")
	menu.AddItem("Warrior", "warrior")
	menu.AddItem("Exit the game", "exit")
	choice := menu.Display()

	switch choice {
	case "mage":
		p.Hero = *NewHero(MAGE)
	case "warrior":
		p.Hero = *NewHero(WARRIOR)
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("Wrong input, dumbass")
		os.Exit(1)
	}
	g.Players = append(g.Players, *p)
	return g
}

func screenLayout(g *Game) {

	s := "+-------------------------------"
	separator := strings.Repeat(s, len(g.Players))
	fmt.Print(separator)
	fmt.Println("+")
	for i := range g.Players {
		fmt.Printf("|\t\t %v  \t\t", i+1)
	}
	fmt.Print("|")
	fmt.Println()
	fmt.Print(separator)
	fmt.Println("+")
	for _, player := range g.Players {
		fmt.Printf("| %v (%v) : \t\t", player.Name, player.Hero.Class.ClassName.ToString())
	}
	fmt.Print("|")
	fmt.Println()
	for _, player := range g.Players {
		fmt.Printf("| => %v hp \t\t\t", player.Hero.HealthPoint)
	}
	fmt.Print("|")
	fmt.Println()
	for _, player := range g.Players {
		fmt.Printf("| => %v mp \t\t\t", player.Hero.MagicPoint)
	}
	fmt.Print("|")
	fmt.Println()
	fmt.Println(separator)

	for _, log := range g.CombatLog {
		printLog(log)
	}

	fmt.Println(strings.Replace(separator, "+", "-", -1))
}

func printLog(log Log) {
	text := fmt.Sprintf("| %v    					\n", log.Value)
	// possible values are red, greed, yellow, blue, magenta, cyan and white
	switch log.LogType {
	case INFO:
		color.Cyan(text)
	case ATTACK:
		color.Red(text)
	case HEAL:
		color.Green(text)
	default:
		color.Cyan(text)
	}
}

func whoStarts(g *Game) {

	// randomizing the players list
	rand.Shuffle(len(g.Players), func(i, j int) { g.Players[i], g.Players[j] = g.Players[j], g.Players[i] })
	for i := range g.Players {
		g.Players[i].Id = i + 1
	}
	g.ActivePlayer = &g.Players[0]
}

func isSomeoneDead(g *Game) {
	for i := range g.Players {
		if g.Players[i].Hero.HealthPoint <= 0 {
			g.appendCombatLog(fmt.Sprintf("%v is dead.", g.Players[i].Name), ATTACK)
			g.Players = append(g.Players[:i], g.Players[i+1:]...)
			break
		}
	}

}

func changeActivePlayer(g *Game) {

	// returns index of the ActivePlayer in Players slice
	i := slices.IndexFunc(g.Players, func(p Player) bool {
		return p.Id == g.ActivePlayer.Id
	})

	if i == len(g.Players)-1 {
		g.ActivePlayer = &g.Players[0]
	} else {
		g.ActivePlayer = &g.Players[i+1]
	}

	g.appendCombatLog(fmt.Sprintf("%v is playing...", g.ActivePlayer.Name), INFO)

}

func pickTarget(g *Game) *Player {
	clearScreen()
	screenLayout(g)
	menu := gocliselect.NewMenu("Choose a target to attack ")
	fmt.Println("")
	for _, p := range g.Players {
		if p.Id != g.ActivePlayer.Id {
			menu.AddItem(p.Name, strconv.Itoa(p.Id))
		}
	}

	choice := menu.Display()

	var target *Player
	for i := range g.Players {
		if strconv.Itoa(g.Players[i].Id) == choice {
			target = &g.Players[i]
		}
	}
	return target
}

func pickPotion(g *Game) {

}

func pickAction(g *Game) {
	clearScreen()
	screenLayout(g)
	menu := gocliselect.NewMenu("Pick an action ")
	menu.AddItem("Attack", "attack")
	// menu.AddItem("Potions", "potions")
	menu.AddItem("Quit game", "exit")
	choice := menu.Display()

	switch choice {
	case "attack":
		s := pickSpell(g)
		t := pickTarget(g)
		g.ActivePlayer.attack(t, s, g)
	case "potions":
		pickPotion(g)
	case "exit":
		os.Exit(0)
	}
}

func pickSpell(g *Game) Spell {

	clearScreen()
	screenLayout(g)
	menu := gocliselect.NewMenu("Choose a spell to cast ")
	for _, s := range g.ActivePlayer.Hero.Class.Spells {
		menu.AddItem(fmt.Sprintf("%v [%v-%v dmg] - %v mp", s.ToString(), spellInfo[s].Min, spellInfo[s].Max, spellInfo[s].Cost), s.ToString())
	}
	choice := menu.Display()
	var spell Spell
	switch choice {
	case "fireball":
		spell = FIREBALL
	case "lightning":
		spell = LIGHTNING
	case "swordstrike":
		spell = SWORDSTRIKE
	case "headbut":
		spell = HEADBUT
	}
	return spell
}

func botPlay(g *Game) {
	clearScreen()
	screenLayout(g)

	potentialTargets := make([]*Player, 0, len(g.Players)-1)

	for _, p := range g.Players {
		if p.Id != g.ActivePlayer.Id {
			potentialTargets = append(potentialTargets, &p)
		}
	}

	if len(potentialTargets) == 0 {
		return
	}

	randomTarget := potentialTargets[rand.Intn(len(potentialTargets))]
	var target *Player
	for i := range g.Players {
		if g.Players[i].Id == randomTarget.Id {
			target = &g.Players[i]
		}

	}
	randomSpell := g.ActivePlayer.Hero.Class.Spells[rand.Intn(len(g.ActivePlayer.Hero.Class.Spells))]

	g.ActivePlayer.attack(target, randomSpell, g)
}

func isThereAWinner(g *Game) bool {
	if len(g.Players) == 1 {
		return true
	} else {
		return false
	}
}

func gameOver(g *Game) {
	g.appendCombatLog(fmt.Sprintf("%v win that Fisht !", g.Players[0].Name), HEAL)
	clearScreen()
	screenLayout(g)
	os.Exit(0)
}

func (g *Game) Fight() {

	whoStarts(g)

	fmt.Println(g.ActivePlayer.Kind)
	for !isThereAWinner(g) {
		changeActivePlayer(g)
		switch g.ActivePlayer.Kind {
		// this is a BOT
		case 0:
			botPlay(g)
		// this is a HUMAN
		case 1:
			pickAction(g)
		}
		isSomeoneDead(g)
	}

	gameOver(g)

	// for g.Hero.HealthPoint > 0 && g.Bot.HealthPoint > 0 {

}
