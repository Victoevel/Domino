package Domino

type Domino struct {
	Dots [2]int
}

func NewDomino(dotOne, dotTwo int) *Domino {
	dom := &Domino{Dots: [2]int{dotOne, dotTwo}}
	return dom
}

func GenerateDominoes() []*Domino {
	var doms [28]*Domino
	count := 0
	for i := 0; i <= 6; i++ {
		for j := 0; j <= i; j++ {
			doms[count] = NewDomino(i, j)
			count++
		}
	}
	return doms[:]
}

func (d *Domino) rotateDomino() {
	d.Dots[0], d.Dots[1] = d.Dots[1], d.Dots[0]
}
