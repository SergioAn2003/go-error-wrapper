# Error Wrapper

Библиотека для обертки ошибок.


# Установка

```bash
  go get -u github.com/SergioAn2003/go-error-wrapper
```
<br>

# Пример
```go
package main

import (
	"errors"
	"fmt"

	"github.com/SergioAn2003/go-error-wrapper"
)

func someFn() (err error) {
	defer errorWrapper.PtrWithOP(&err, "path.someFn")
	return errors.New("some error")
}

func main() {
	err := errorWrapper.WithOP(someFn(), "main")

	// main: path.someFn: some error
	fmt.Println(err)
}
```
