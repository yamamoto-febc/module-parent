package parent

import (
	child "github.com/yamamoto-febc/module-child"
)

func Foobar() string {
	return child.FooBar() + " from parent"
}
