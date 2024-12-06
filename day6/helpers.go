package main

type Pair struct {
	X int
	Y int
}

func out_of_bounds(a Pair, b Pair) bool {
	return a.X < 0 || a.Y < 0 || a.X > b.X || a.Y > b.Y
}

func add(a Pair, b Pair) Pair {
	return Pair{X: a.X + b.X, Y: a.Y + b.Y}
}

func hash(a Pair) string {
	return string(a.X) + "-" + string(a.Y)
}

func turn(a Pair) Pair {
	if a.X == 1 {
		return Pair{X: 0, Y: 1}
	}
	if a.Y == 1 {
		return Pair{X: -1, Y: 0}
	}
	if a.X == -1 {
		return Pair{X: 0, Y: -1}
	}
	if a.Y == -1 {
		return Pair{X: 1, Y: 0}
	}
	panic("Wrong direction pair")
}
