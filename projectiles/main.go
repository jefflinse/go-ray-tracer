package main

import (
	"fmt"

	"github.com/jefflinse/go-ray-tracer/rt"
)

type environment struct {
	Gravity *rt.Tuple
	Wind    *rt.Tuple
}

type projectile struct {
	Position *rt.Tuple
	Velocity *rt.Tuple
}

func main() {
	p := &projectile{rt.NewPoint(0, 1, 0), rt.NewVector(1, 1, 0).Normalize()}
	env := &environment{rt.NewVector(0, -0.1, 0), rt.NewVector(-0.01, 0, 0)}
	for p.Position.Y > 0 {
		p = tick(env, p)
		fmt.Println(p.Position)
	}
}

func tick(env *environment, p *projectile) *projectile {
	return &projectile{
		p.Position.Add(p.Velocity),
		p.Velocity.Add(env.Gravity).Add(env.Wind),
	}
}
