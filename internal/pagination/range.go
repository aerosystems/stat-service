package pagination

import (
	"errors"
	"strconv"
)

const (
	DefaultLimit  = 10
	DefaultOffset = 0
)

type Range struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

func NewRange(limit, offset int) *Range {
	return &Range{
		Limit:  limit,
		Offset: offset,
	}
}

func GetFromQuery(query map[string][]string) (*Range, error) {
	limit := DefaultLimit
	offset := DefaultOffset
	var err error

	if len(query["limit"]) != 0 {
		limit, err = strconv.Atoi(query["limit"][0])
		if err != nil {
			return nil, errors.New("limit parameter must be integer")
		}
	}

	if len(query["offset"]) != 0 {
		offset, err = strconv.Atoi(query["offset"][0])
		if err != nil {
			return nil, errors.New("offset parameter must be integer")
		}
	}

	return NewRange(limit, offset), nil
}
