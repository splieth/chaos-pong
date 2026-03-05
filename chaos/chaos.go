package chaos

import (
	"math/rand"
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
	EBSDestroyChaos{},
}

func Random() {
	fn := functions[rand.Intn(len(functions))]
	fn.Terminate()
}
