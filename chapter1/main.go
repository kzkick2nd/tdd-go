package main

type Doller struct {
	amount int
}

func (d *Doller) times(t int) {
	d.amount *= t
}

func main() {}
