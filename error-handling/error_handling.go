package erratum

import "errors"

// Use open a resource, call Frob(input) on the result resource and then close that resource (in all cases)
func Use(opener ResourceOpener, h string) (err error) {
	var res Resource
	var openErr error

	res, openErr = opener()
	for openErr != nil {
		if _, ok := openErr.(TransientError); !ok {
			return openErr
		}
		res, openErr = opener()
	}
	defer res.Close()
	defer func() {
		if r := recover(); r != nil {
			if recoveredErr, ok := r.(FrobError); ok {
				res.Defrob(recoveredErr.defrobTag)
				err = recoveredErr.inner
			} else {
				var ok bool
				if err, ok = r.(error); !ok {
					err = errors.New("panic without error")
				}
			}
		}
	}()
	res.Frob(h)
	return nil
}
