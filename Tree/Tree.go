package Tree

import (
	"Domino/Domino"
	"Domino/Player"
	"fmt"
	"math/rand"
	"time"
)

type Tree struct {
	RightEnd, LeftEnd   int
	Original            *Domino.Domino
	RightNode, LeftNode *node
}

type node struct {
	value    *Domino.Domino
	nextNode *node
}

func (t *Tree) findEnds() {
	t.RightEnd = t.RightNode.findLast().value.Dots[0]
	t.LeftEnd = t.LeftNode.findLast().value.Dots[1]
}

func (n *node) findLast() *node {
	if n.nextNode == nil {
		return n
	} else {
		return n.nextNode.findLast()
	}
}

func (n *node) append(d *Domino.Domino) bool {
	newNode := n.findLast()
	if compareDominoes(newNode.value, d) {
		newNode.nextNode = &node{value: d, nextNode: nil}
		return true
	}
	return false
}

func Play(player *Player.Player, index int, tree Tree) bool {
	if index > len(player.Hand) || index < 0 {
		return false
	}
	var nD *Domino.Domino
	nD, player.Hand = PopIndex(player.Hand, index)
	if tree.LeftNode == nil {
		if nD.Dots[0] == 6 || nD.Dots[1] == 6 {
			tree.LeftNode = &node{value: nD}
			return true
		}
		return false
	} else if tree.LeftNode == nil {
		if nD.Dots[0] == 6 || nD.Dots[1] == 6 {
			tree.RightNode = &node{value: nD}
			return true
		}
		return false
	}
	if tree.LeftNode.append(nD) || tree.RightNode.append(nD) {
		return true
	} else {
		player.Hand = append(player.Hand, nD)
		return false
	}
}

func (t *Tree) PrintTree() {
	fmt.Printf("Original: %v\n", t.Original)
	fmt.Printf("Left Nodes: \n")
	currNode := t.LeftNode
	for {
		if currNode == nil {
			break
		}
		fmt.Printf("Node: %v\n", currNode)
		currNode = currNode.nextNode
	}
	fmt.Printf("Right Nodes: \n")
	currNode = t.RightNode
	for {
		if currNode == nil {
			break
		}
		fmt.Printf("Node: %v\n", currNode)
		currNode = currNode.nextNode
	}
}

func compareDominoes(d1, d2 *Domino.Domino) bool {
	return d1.Dots[0] == d2.Dots[0] || d1.Dots[1] == d2.Dots[0] || d1.Dots[0] == d2.Dots[1] || d1.Dots[1] == d2.Dots[1]
}

func Pop(a []*Domino.Domino) (*Domino.Domino, []*Domino.Domino) {
	i := a[len(a)-1]
	a = a[:len(a)-1]
	return i, a
}

func ShuffleArray(arr []*Domino.Domino) []*Domino.Domino {
	var oldArr []*Domino.Domino
	var newArr []*Domino.Domino
	for _, v := range arr {
		oldArr = append(oldArr, v)
	}

	for range arr {
		i := randomInt(len(oldArr))
		newArr = append(newArr, oldArr[i])
		oldArr[i] = oldArr[len(oldArr)-1]
		oldArr[len(oldArr)-1] = nil
		oldArr = oldArr[:len(oldArr)-1]
	}
	return newArr
}

func randomInt(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

func PopIndex(a []*Domino.Domino, index int) (*Domino.Domino, []*Domino.Domino) {
	a[len(a)-1], a[index] = a[index], a[len(a)-1]
	i := a[len(a)-1]
	a = a[:len(a)-1]
	return i, a
}
