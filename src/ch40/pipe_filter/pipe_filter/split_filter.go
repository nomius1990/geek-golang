package pipefilter

import (
	"errors"
	"strings"
)

var SpiltFilterWrongFormatError = errors.New("input data should be string")

type SpiltFilter struct {
	delimiter string
}

func NewSpiltFilter(delimiter string) *SpiltFilter{
	return &SpiltFilter{delimiter}
}

func (sf *SpiltFilter) Process(data Request)(Response,error){
	str,ok:= data.(string)
	if !ok{
		return nil,SpiltFilterWrongFormatError
	}
	parts := strings.Split(str,sf.delimiter)
	return parts,nil
}