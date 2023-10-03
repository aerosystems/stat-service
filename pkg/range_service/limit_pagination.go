package RangeService

import (
	"errors"
	"strconv"
)

const (
	DefaultLimit  = 10
	DefaultOffset = 0
)

type LimitPagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

func NewLimitPagination(limit, offset int) *LimitPagination {
	return &LimitPagination{
		Limit:  limit,
		Offset: offset,
	}
}

func GetLimitPaginationFromQuery(query map[string][]string) (*LimitPagination, error) {
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

	return NewLimitPagination(limit, offset), nil
}
