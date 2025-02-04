package storage

type Storage interface {
	Save(key string, data []byte) error
	Load(key string) ([]byte, error)
	Exists(key string) (bool, error)
	Delete(key string) error
	Close() error
}
