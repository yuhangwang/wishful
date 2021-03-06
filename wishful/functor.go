package wishful

type Functor interface {
	Map(f Morphism) Functor
}

type FunctorLaws struct {
	x func(Any) Functor
}

func NewFunctorLaws(fun func(Any) Functor) FunctorLaws {
	return FunctorLaws{
		x: fun,
	}
}

func (o FunctorLaws) Identity(run Morphism) (func(int) Any, func(int) Any) {
	f := func(v int) Any {
		a := o.x(v)
		return run(a.Map(Identity))
	}
	g := func(v int) Any {
		return run(o.x(v))
	}
	return f, g
}

func (o FunctorLaws) Composition(run Morphism) (func(int) Any, func(int) Any) {
	f := func(v int) Any {
		a := o.x(v).(Functor)
		return run(a.Map(Compose(Identity)(Identity)))
	}
	g := func(v int) Any {
		a := o.x(v).(Functor)
		return run(a.Map(Identity).Map(Identity))
	}
	return f, g
}
