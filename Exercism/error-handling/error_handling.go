package erratum

import (
	"errors"
)

func Use(f ResourceOpener, s string) (err error) {

	var resource Resource

	exitLoop := false

	for !exitLoop {
		resource, err = f()
		switch {
		case errors.As(err, &TransientError{}):
			continue
		case err != nil:
			return err
		default:
			exitLoop = true
		}
	}

	defer resource.Close()

	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case FrobError:
				resource.Defrob(e.defrobTag)
				err = e.inner
			case error:
				err = e
			}
		}
	}()

	resource.Frob(s)

	return err
}
