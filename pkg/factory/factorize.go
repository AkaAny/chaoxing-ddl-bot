package factory

import (
	"errors"
	"fmt"
)

type Config interface {
	GetType() string
}

type Creator interface {
	Config
	Create(conf Config) (Object, error)
}

type Object interface{}

type Factorize struct {
	typeCreatorMap map[string]Creator
}

func (f *Factorize) Put(creators ...Creator) error {
	if f.typeCreatorMap == nil {
		f.typeCreatorMap = make(map[string]Creator)
	}
	for _, creator := range creators {
		if creator == nil {
			return errors.New("creator cannot be nil")
		}
		f.typeCreatorMap[creator.GetType()] = creator
	}
	return nil
}

func (f Factorize) Get(t string) Creator {
	creator, valid := f.typeCreatorMap[t]
	if !valid {
		return nil
	}
	return creator
}

func (f Factorize) DoCreate(conf Config) (interface{}, error) {
	if conf == nil {
		return nil, errors.New("conf cannot be nil")
	}
	var t = conf.GetType()
	var creator = f.Get(t)
	if creator == nil {
		return nil, fmt.Errorf("type:%s is not found", t)
	}
	return creator.Create(conf)
}
