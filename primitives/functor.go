package primitives

type Functor interface {
	Map(MapFunc) Functor
}

type MapFunc func(in interface{}) interface{}
