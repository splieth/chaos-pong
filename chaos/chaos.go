package chaos

import (
	"math/rand"
	"time"
)

type Chaos interface {
	Terminate() Result
}

type Result struct {
	Success bool
	Message string
}

var functions = []Chaos{
	EC2InstanceTerminateChaos{},
}

func init() {
	rand.Seed(time.Now().Unix())
}

func Random() {
	fn := functions[rand.Intn(len(functions))]
	fn.Terminate()
}
