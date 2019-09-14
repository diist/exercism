package erratum

import (
	"fmt"
)

func Use(o ResourceOpener, input string) error {
	r, e := open(o)

	if e != nil {
		return e
	}

	defer r.Close()
	r.Frob(input)

	return nil
}

func open(o ResourceOpener) (r Resource, e error) {
	for {
		r, e := o()

		if e == nil {
			return r, e
		}

		_, isTransient := e.(TransientError)
		if !isTransient {
			fmt.Println("Transient Error detected")
			return r, e
		}
	}
}
