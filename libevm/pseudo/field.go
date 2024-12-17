package pseudo

import "fmt"

type FieldIn[C any] interface {
	Get(C) *Type
	Default() *Type
	SetToDefault(C)
}

func NewFieldIn[C any, T any](by Construction, fieldPtr func(C) **Type) (FieldIn[C], error) {
	if !by.IsValid() {
		return field[C]{}, fmt.Errorf("invalid %T(%[1]d)", by)
	}

	ctor := NewConstructor[T]()
	return field[C]{
		makeDefault: func() *Type { return ctor.Construct(by) },
		pointer:     fieldPtr,
	}, nil
}

func MustNewFieldIn[C any, T any](by Construction, fieldPtr func(C) **Type) FieldIn[C] {
	f, err := NewFieldIn[C, T](by, fieldPtr)
	if err != nil {
		panic(err)
	}
	return f
}

type field[C any] struct {
	makeDefault func() *Type
	pointer     func(C) **Type
}

func (f field[C]) Get(c C) *Type {
	p := f.pointer(c)
	if *p == nil {
		*p = f.makeDefault()
	}
	return *p
}

func (f field[C]) Default() *Type {
	return f.makeDefault()
}

func (f field[C]) SetToDefault(c C) {
	*f.pointer(c) = f.makeDefault()
}
