package storage

import (
	"encoding/json"
	"sync"

	"github.com/ilker-raimov/cca/common/storage/model/model_user"
)

var (
	instance *BadgerObjectify
	once     sync.Once
)

func GetInstance() *BadgerObjectify {
	once.Do(func() {
		instance = &BadgerObjectify{badger: GetBadgerInstance()}

		setup(instance)
	})

	return instance
}

func setup(bo *BadgerObjectify) {
	competitor := model_user.Competitor()
	organizer := model_user.Organizer()
	admin := model_user.Admin()

	bo.Save().Entity(competitor).Now()
	bo.Save().Entity(organizer).Now()
	bo.Save().Entity(admin).Now()
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

type BadgerLoadOrCreateObjectify struct {
	bo     *BadgerObjectify
	action func() error
}

func (blco *BadgerLoadOrCreateObjectify) Entity(object Objectifiable[string], id string, fallback Objectifiable[string]) *BadgerLoadOrCreateObjectify {
	return &BadgerLoadOrCreateObjectify{
		bo: blco.bo,
		action: func() error {
			exists, exist_err := blco.bo.badger.Exists(id)

			if exist_err != nil {
				return exist_err
			}

			var data []byte
			var err error

			if !exists {
				data, err = json.Marshal(fallback)
			} else {
				data, err = blco.bo.badger.Load(id)
			}

			if err != nil {
				return err
			}

			return json.Unmarshal(data, object)
		},
	}
}

func (blco *BadgerLoadOrCreateObjectify) Now() error {
	return blco.action()
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

func (bo *BadgerObjectify) LoadOrCreate() *BadgerLoadOrCreateObjectify {
	return &BadgerLoadOrCreateObjectify{
		bo: bo,
	}
}
