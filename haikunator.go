package haikunator

import (
	"fmt"
	"math/rand"
)

var (
	ADJECTIVES = asArray(ADVECTIVES_STR)
	NOUNS      = asArray(NOUNS_STR)
)

type Name interface {
	Haikunate() string
}

type RandomName struct {
	r *rand.Rand
}

func (r RandomName) Haikunate() string {
	return fmt.Sprintf("%v-%v", ADJECTIVES[r.r.Intn(len(ADJECTIVES))], NOUNS[r.r.Intn(len(NOUNS))])
}

func New(seed int64) Name {
	name := RandomName{rand.New(rand.New(rand.NewSource(99)))}
	name.r.Seed(seed)
	return name
}
