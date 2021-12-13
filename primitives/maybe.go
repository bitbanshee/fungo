package primitives

// Maybe is a representation of and additiv monad
type Maybe Monad

var none = &Nothing{}

func Just(v interface{}) Maybe {
	if v == nil {
		return none
	}
	if cv, ok := v.(Nothing); ok {
		return &cv
	}
	if cv, ok := v.(just); ok {
		return &cv
	}
	return &just{value: v}
}

type Nothing struct{}

func (m *Nothing) Unit() interface{} {
	return m
}

func (m *Nothing) Bind(f MonadicFunc) Monad {
	return m
}

func (m *Nothing) Map(mf MapFunc) Functor {
	return m
}

type just struct {
	value   interface{}
	closure func() interface{}
}

func (m *just) Unit() interface{} {
	if m.value == nil {
		m.value = m.closure()
		// free space used by closure
		m.closure = nil
	}
	return m.value
}

func (m *just) Bind(f MonadicFunc) Monad {
	binded := func() interface{} {
		rm := f(m.Unit())
		if _, ok := rm.(*Nothing); ok {
			return rm
		}
		return rm.Unit()
	}
	return &just{closure: binded}
}

func (m *just) Map(mf MapFunc) Functor {
	binded := func() interface{} {
		r := mf(m.Unit())
		if r == nil {
			return none
		}
		if n, ok := r.(*Nothing); ok {
			return n
		}
		return r
	}
	return &just{closure: binded}
}
