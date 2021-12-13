package main

import (
	"fmt"

	"github.com/bitbanshee/fungo/collections"
	"github.com/bitbanshee/fungo/primitives"
)

func main() {
	mten := primitives.Just(10)
	mtwenty := mten.Bind(addTen)
	if _, ok := mtwenty.(*primitives.Nothing); ok {
		fmt.Printf("no result for Just(10) + 10")
	} else {
		fmt.Printf("result 10 + 10: %v", mtwenty.Unit())
	}

	mnothing := primitives.Nothing{}
	mtwentyFromNothing := mnothing.Bind(addTen)
	if _, ok := mtwentyFromNothing.(*primitives.Nothing); ok {
		fmt.Printf("\nno result for Nothing + 10")
	} else {
		fmt.Printf("\nresult Nothing + 10: %v", mtwentyFromNothing.Unit())
	}

	data := []interface{}{"c", "a", "b", "b", "c"}
	set := collections.NewSet(data...)
	iterator := set.Iterate()
	fmt.Println("\niterating over set")
	for {
		maybeNext := iterator.Next()
		if _, ok := maybeNext.(*primitives.Nothing); ok {
			break
		}
		fmt.Println(maybeNext.Unit())
	}

	fset := set.
		Map(addY).
		Map(addA).(collections.Set)
	fiterator := fset.Iterate()
	fmt.Println("iterating over mapped set")
	for {
		maybeNext := fiterator.Next()
		if _, ok := maybeNext.(*primitives.Nothing); ok {
			break
		}
		fmt.Println(maybeNext.Unit())
	}
}

func addTen(cur interface{}) primitives.Monad {
	if i, ok := cur.(int); ok {
		return primitives.Just(i + 10)
	}
	return &primitives.Nothing{}
}

func addY(cur interface{}) interface{} {
	if i, ok := cur.(string); ok {
		return i + "y"
	}
	// I could return nil here too
	return &primitives.Nothing{}
}

func addA(cur interface{}) interface{} {
	if i, ok := cur.(string); ok {
		return i + "a"
	}
	// I could return nil here too
	return &primitives.Nothing{}
}
