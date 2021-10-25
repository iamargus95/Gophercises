package robotname

import (
	"errors"
	"time"
)

const maxNumNames = 26 * 26 * 10 * 10 * 10

type Robot struct {
	name string
}

var (
	issuedNames = 0
	generator   = NewRobotLCG()
)

func (r *Robot) Name() (string, error) {
	if issuedNames == maxNumNames {
		return "", errors.New("exhausted all names")
	}
	if r.name != "" {
		return r.name, nil
	}
	r.name = calculateName(generator.Next())
	issuedNames++
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}

// calculateName maps value to its corresponding string name.
// value must be >= 0 and < maxNumNames
func calculateName(value int64) string {
	if value < 0 || value >= maxNumNames {
		panic("receieved out of bounds value")
	}
	let := rune(value / 1000)
	num := rune(value % 1000)
	first, second := let/26, let%26
	third := num / 100
	fourth := (num / 10) % 10
	fifth := num % 10
	return string([]rune{
		'A' + first,
		'A' + second,
		'0' + third,
		'0' + fourth,
		'0' + fifth,
	})
}

type lcg struct {
	a int64
	c int64
	m int64
	x int64
}

func NewLCG(a, c, m int64) lcg {
	return lcg{
		a: a,
		c: c,
		m: m,
		x: 0,
	}
}

func (g *lcg) Seed(seed int64) {
	g.x = seed % g.m
}

func (g *lcg) Next() int64 {
	g.x = (g.a*g.x + g.c) % g.m
	return g.x
}

// NewRobotLCG creates a linear congruential generator with period equal
// to 676K. This function returns a PRNG that will produce 676K unique
// integers before repeating.
func NewRobotLCG() lcg {
	// choose a, c, m such that
	// m = 676K
	// c != 0
	// c and m are coprime
	// a-1 is divisible by all prime factors of m (2, 5, 13)
	// a-1 is divisible by 4
	// a = 261
	// c = 399757
	// m = 676000
	g := NewLCG(261, 399757, 676000)
	g.Seed(time.Now().UnixNano())
	return g
}
