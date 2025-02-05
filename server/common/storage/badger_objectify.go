package storage

import (
	"encoding/json"
	"sync"
)

var (
	instance *BadgerObjectify
	once     sync.Once
)

func GetInstance() *BadgerObjectify {
	once.Do(func() {
		instance = &BadgerObjectify{badger: GetBadgerInstance()}
	})

	return instance
}

type BadgerSaveObjectify struct {
	bo     *BadgerObjectify
	action func() error
}

func (bso *BadgerSaveObjectify) Entity(object Objectifiable[string]) *BadgerSaveObjectify {
	return &BadgerSaveObjectify{
		bo: bso.bo,
		action: func() error {
			key := object.Key()
			data, json_err := json.Marshal(object)

			if json_err != nil {
				return json_err
			}

			return bso.bo.badger.Save(key, data)
		},
	}
}

func (bso *BadgerSaveObjectify) Now() error {
	return bso.action()
}

type BadgerLoadObjectify struct {
	bo     *BadgerObjectify
	action func() error
}

func (blo *BadgerLoadObjectify) Entity(object Objectifiable[string], id string) *BadgerLoadObjectify {
	return &BadgerLoadObjectify{
		bo: blo.bo,
		action: func() error {
			data, err := blo.bo.badger.Load(id)

			if err != nil {
				return err
			}

			return json.Unmarshal(data, object)
		},
	}
}

func (blo *BadgerLoadObjectify) Now() error {
	return blo.action()
}

type BadgerExistObjectify struct {
	bo     *BadgerObjectify
	action func() (bool, error)
}

func (beo *BadgerExistObjectify) Entity(id string) *BadgerExistObjectify {
	return &BadgerExistObjectify{
		bo: beo.bo,
		action: func() (bool, error) {
			return beo.bo.badger.Exists(id)
		},
	}
}

func (beo *BadgerExistObjectify) NowT() (bool, error) {
	return beo.action()
}

type BadgerObjectify struct {
	badger *Badger
}

func (bo *BadgerObjectify) Save() *BadgerSaveObjectify {
	return &BadgerSaveObjectify{
		bo: bo,
	}
}

func (bo *BadgerObjectify) Load() *BadgerLoadObjectify {
	return &BadgerLoadObjectify{
		bo: bo,
	}
}

func (bo *BadgerObjectify) Exist() *BadgerExistObjectify {
	return &BadgerExistObjectify{
		bo: bo,
	}
}
