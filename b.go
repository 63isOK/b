package b

import (
	"github.com/63isOK/a"
)

func GetSum() int {
	o := a.Obj{
		A: 1,
		B: 2,
		C: 3,
	}

	return o.Sum()
}
