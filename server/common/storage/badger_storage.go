package storage

import (
	"sync"

	"github.com/dgraph-io/badger/v4"
	"github.com/ilker-raimov/cca/common/util/file"
)

var (
	badger_instance *Badger
	badger_once     sync.Once
)

const (
	STORAGE_PATH = ".storage_db"
)

func GetBadgerInstance() *Badger {
	badger_once.Do(func() {
		cleanup()

		opts := badger.DefaultOptions(STORAGE_PATH).WithLoggingLevel(badger.INFO)
		db_temp, err := badger.Open(opts)

		if err != nil {
			panic(err)
		}

		badger_instance = &Badger{db: db_temp}
	})

	return badger_instance
}

func cleanup() {
	storage_exists := file.Exists(STORAGE_PATH)

	if !storage_exists {
		return
	}

	err := file.DeleteAll(STORAGE_PATH)

	if err == nil {
		return
	}

	panic(err)
}

type Badger struct {
	db *badger.DB
}

func (b *Badger) Save(key string, data []byte) error {
	return b.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), data)
	})
}

func (b *Badger) Load(key string) ([]byte, error) {
	var value []byte

	err := b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))

		if err != nil {
			return err
		}

		value, err = item.ValueCopy(nil)

		return err
	})

	if err != nil {
		return nil, err
	}

	return value, nil
}

func (b *Badger) Exists(key string) (bool, error) {
	var exists bool

	err := b.db.View(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte(key))
		if err == nil {
			exists = true

			return nil
		}

		if err == badger.ErrKeyNotFound {
			exists = false

			return nil
		}

		return err
	})

	return exists, err
}

func (b *Badger) Delete(key string) error {
	return b.db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})
}

func (b *Badger) Close() error {
	return b.db.Close()
}
