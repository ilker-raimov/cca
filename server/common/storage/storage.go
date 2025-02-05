package storage

type Storage[T any, R any] interface {
	Save(key T, data R) error
	Load(key T) (R, error)
	Exists(key T) (bool, error)
	Delete(key T) error
	Close() error
}
