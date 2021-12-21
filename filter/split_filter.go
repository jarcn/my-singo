package filter

import (
	"errors"
	"strings"
)

var SplitFilterWrongFormatError = errors.New("input data should be string")

type SplitFilter struct {
	delimiter string
}

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter}
}

func (sf *SplitFilter) Process(data Request) (Responese, error) {
	if str, ok := data.(string); !ok {
		return nil, SplitFilterWrongFormatError
	} else {
		parts := strings.Split(str, sf.delimiter)
		return parts, nil
	}
}
