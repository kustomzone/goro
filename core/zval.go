package core

import "fmt"

type Val interface {
	GetType() ZType
}

type ZVal struct {
	v Val
}

func (z *ZVal) run(ctx Context) (*ZVal, error) {
	return z, nil
}

func (z *ZVal) IsNull() bool {
	if z == nil {
		return true
	}
	if z.v == nil {
		return true
	}
	return false
}

func (z *ZVal) String() string {
	switch z.v.GetType() {
	default:
		return fmt.Sprintf("%+v", z.v)
	}
}
