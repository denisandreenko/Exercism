package robotname

import (
	"errors"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//Robot 
type Robot struct {
	name string
}

const nameLength = 5
const lettersCount = 2

var globalNameSpace = make(map[string]struct{})

const max = 26 * 26 * 10 * 10 * 10

//Name - give robot a new name
func (r *Robot) Name() (string, error) {
	if len(r.name) > 0 {
		return r.name, nil
	}
	if len(globalNameSpace) == max {
		return "", errors.New("namespace exhausted")
	}
	r.name = newName()
	return r.name, nil
}

//Reset - reset robot name
func (r *Robot) Reset() {
	r.name = ""
}

func newName() string {
	b := make([]rune, nameLength)
	for i := range b {
		if i+1 <= lettersCount {
			b[i] = rune('A' + rand.Intn(26))
		} else {
			b[i] = rune('0' + rand.Intn(10))
		}
	}
	name := string(b)
	if _, ok := globalNameSpace[name]; ok {
		return newName()
	}
	globalNameSpace[name] = struct{}{}
	return name
}
