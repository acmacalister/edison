package main

type Freedom struct {
	X  int
	Y  int
	Z  int
	a1 int
	a2 int
	a3 int
	m1 int
	m2 int
	m3 int
}

func parseFreedomData(data []byte) *Freedom {
	return &Freedom{X: int(data[0]), Y: int(data[1]), Z: int(data[2]),
		a1: int(data[3]), a2: int(data[4]), a3: int(data[5]),
		m1: int(data[6]), m2: int(data[7]), m3: int(data[8])}
}
