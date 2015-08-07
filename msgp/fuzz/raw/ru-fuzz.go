package ru

import "github.com/tinylib/msgp/msgp"

func Fuzz(data []byte) int {
	var r msgp.Raw
	_, err := r.UnmarshalMsg(data)
	if err != nil {
		return 0
	}
	left, err := msgp.Skip([]byte(r))
	if err != nil {
		panic("Skip() failed after UnmarshalMsg(): " + err.Error())
	}
	if len(left) != 0 {
		panic("data left over after Skip(): " + string(left))
	}
	// Right now, the output
	// of MarshalMsg is simply
	// uninteresting; it
	// just returns itself.
	return 1
}
