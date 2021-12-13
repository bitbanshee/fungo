package primitives

type Monad interface {
	Functor
	Unit() interface{}
	Bind(MonadicFunc) Monad
}

type MonadicFunc func(interface{}) Monad
