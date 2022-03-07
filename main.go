package parent

import (
	"fmt"
	child "github.com/yamamoto-febc/module-child"
)

func Foobar() {
	fmt.Println(child.FooBar() + " from parent")
}
