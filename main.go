package main

import (
	"Domino/Domino"
	"Domino/Player"
	"Domino/Tree"
	"fmt"
	"os"
	"os/exec"
)

type Game struct {
	tree        Tree.Tree
	players     []*Player.Player
	shufflePile []*Domino.Domino
	firstTurn   int
}

func Init() *Game {
	game := &Game{firstTurn: 0}
	game.shufflePile = Tree.ShuffleArray(Domino.GenerateDominoes())

	var n int

	for {
		fmt.Printf("Ingresar numero de Jugadores (Entre 2 y 4): ")
		fmt.Scan(&n)
		if n > 4 || n < 2 {
			fmt.Printf("Demasiados o muy pocos jugadores!\n\n")
			continue
		}
		fmt.Printf("Iniciando con %d jugadores\n", n)
		break
	}

	for i := 0; i < n; i++ {
		game.players = append(game.players, &Player.Player{})
	}

	for i, _ := range game.players {
		for j := 0; j < 7; j++ {
			var v *Domino.Domino
			v, game.shufflePile = Tree.Pop(game.shufflePile)
			game.players[i].Hand = append(game.players[i].Hand, v)
		}
	}

	largestIndex := 0
	largestN := 0
	var largestHand *Player.Player

	for i, v := range game.players {
		for j, _ := range v.Hand {
			if v.Hand[j].Dots[0] == v.Hand[j].Dots[1] {
				if v.Hand[j].Dots[0] > largestN {
					largestN = v.Hand[j].Dots[0]
					largestIndex = j
					largestHand = v
					game.firstTurn = wrap(0, len(game.players)-1, i+1)
				}
			}
		}
	}
	if largestN == 0 {
		for i, v := range game.players {
			for j, _ := range v.Hand {
				if v.Hand[j].Dots[0]+v.Hand[j].Dots[1] > largestN {
					largestN = v.Hand[j].Dots[0] + v.Hand[j].Dots[1]
					largestIndex = j
					largestHand = v
					game.firstTurn = wrap(0, len(game.players)-1, i+1)
				}
			}
		}
	}
	game.tree.Original, largestHand.Hand = Tree.PopIndex(largestHand.Hand, largestIndex)

	return game
}

func (g *Game) Print() {
	fmt.Printf("Tree: \n\tRightEnd: %d\n\tLeftEnd: %d\n\tOriginal: %v\n\tRightNode: %v\n\tLeftNode: %v\n", g.tree.RightEnd, g.tree.LeftEnd, g.tree.Original, g.tree.RightNode, g.tree.LeftNode)
	fmt.Printf("Players:\n\tHands:\n")
	for _, p := range g.players {
		for _, v := range p.Hand {
			fmt.Printf("\t%v", *v)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\nShuffle Pile:")
	for _, v := range g.shufflePile {
		fmt.Printf("\t%v", *v)
	}
	fmt.Printf("\n\nFirst Turn: %d\n", g.firstTurn)
}

func (g *Game) isEmpty() bool {
	for _, v := range g.players {
		if len(v.Hand) == 0 {
			return true
		}
	}
	return false
}

func wrap(min, max, n int) int {
	if n < min {
		return n + max + 1
	} else if n > max {
		return n - max - 1
	} else {
		return n
	}
}

func main() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	game := Init()
	game.Print()
	for !game.isEmpty() {
		var cmd string

		flag := true
		for flag {
			fmt.Printf("Es turno del jugador numero %d!\n", game.firstTurn)
			fmt.Printf("Ingrese su comando (Escribir 'ayuda' para ver lista de comandos): ")
			fmt.Scan(&cmd)
			switch cmd {
			case "ayuda":
				fmt.Printf(`
Lista de Comandos:
	"ayuda"     : Muestra los comandos del juego.
	"mano"      : Muestra la mano actual.
	"tablero"   : Muestra el estado actual del tablero.
	"pasar"     : Si no tienes nada que jugar, puedes  robar una ficha y pasar tu turno.
	"jugar [i]" : Puedes intentar jugar una ficha de tu mano. El juego determinara donde poner la ficha automaticamente.
	"salir"     : Salir del juego.
`)
			case "mano":
				for _, v := range game.players[game.firstTurn].Hand {
					fmt.Printf("\t%v", *v)
				}
				fmt.Printf("\n")
			case "tablero":
				fmt.Printf("Funcionamiento Pendiente\n")
			case "pasar":
				var newFicha *Domino.Domino
				if !(len(game.shufflePile) == 0) {
					newFicha, game.shufflePile = Tree.Pop(game.shufflePile)
					game.players[game.firstTurn].Hand = append(game.players[game.firstTurn].Hand, newFicha)
				}
				game.firstTurn = wrap(0, len(game.players)-1, game.firstTurn+1)
				flag = false
				c := exec.Command("clear")
				c.Stdout = os.Stdout
				c.Run()
			case "salir":
				return
			default:
				fmt.Printf("Comando no reconocido, intente de nuevo!\n")
			}
		}
	}
}
