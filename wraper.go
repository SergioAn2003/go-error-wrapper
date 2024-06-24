package errorWrapper

import (
	"fmt"
)

// PtrWithOP обернет имеющуюся ошибку в строку формата "<op>: <err>".
//
// Пример использования:
//
//	 package main
//
//	 import (
//	   "errors"
//	   "fmt"
//
//	   errorWrapper "gitlab.app.cube/medcore/go-utils/error-wrapper"
//	 )
//
//	 func someFn() (err error) {
//		defer errorWrapper.PtrWithOP(&err, "path.someFn")
//		return errors.New("some error")
//	}
//
//	func main() {
//		// path.someFn: some error
//		fmt.Println(someFn())
//	}
func PtrWithOP(err *error, op string) { //nolint:gocritic
	if *err != nil {
		*err = fmt.Errorf("%s: %w", op, *err)
	}
}

// WithOP вернет обернутую ошибку в формате "<op>: <err>".
//
// Пример использования:
//
//	package main
//
//	import (
//		"errors"
//		"fmt"
//
//		errorWrapper "gitlab.app.cube/medcore/go-utils/error-wrapper"
//	)
//
//	func main() {
//		err := errors.New("some error")
//
//		// wrap message: some error
//		fmt.Println(errorWrapper.WithOP(err, "wrap message"))
//	}
func WithOP(err error, op string) error {
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
