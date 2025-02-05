package storage

type Objectifiable[T any] interface {
	Key() T
}

type Now interface {
	Now() error
}

type NowT[T any] interface {
	Now() T
}

type SaveObjectify[T any] interface {
	Entity(object Objectifiable[T]) *SaveObjectify[T]
}

type LoadObjectify[T any] interface {
	Entity(object Objectifiable[T], id T) *LoadObjectify[T]
}

type ExistObjectify[T any] interface {
	Entity(id T) *ExistObjectify[T]
}

type Objectify[T any, L LoadObjectify[T], S SaveObjectify[T], E ExistObjectify[T]] interface {
	Save() *S
	Load() *L
	Exist() *E
	// Delete(id T) error
}
