package models

import "github.com/gsinekliev/eval-service/service/eval"

type Error struct {
	Expression string      `json:"expression"`
	Endpoint   string      `json:"endpoint"`
	Frequency  int         `json:"frequency"`
	ErrorType  eval.Status `json:"type"`
}

type ErrorStore map[string]Error

func InitErrorStore() ErrorStore {
	return make(ErrorStore)
}

func (store *ErrorStore) AddError(err Error) {
	key := err.Endpoint + err.Expression
	value, ok := (*store)[key]
	if ok {
		value.Frequency += err.Frequency
		(*store)[key] = value
	} else {
		(*store)[key] = err
	}
}

type Request struct {
	Expression string `json:"expression"`
}
