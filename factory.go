package addreality

import (
	"github.com/pkg/errors"
)

const errorParams = `count of lines and params must be larger than 0 or 0 if unlimited`
const errorCountFields = `count of fields must be larger than 0`
const errorCountLines = `count of lines must be larger than 0`

type InsertBuilderFactory interface {
	CreateInsertBuilder(table string, fields ...string) (InsertBuilder, error)
}

type factory struct {
	maxCountLines int
	maxCountArgs  int
}

func NewInsertBuilderFactory(lines, params int) (InsertBuilderFactory, error) {
	if lines < 0 || params < 0 {
		return nil, errors.New(errorParams)
	}
	return &factory{maxCountLines: lines, maxCountArgs: params}, nil
}

func (r *factory) CreateInsertBuilder(table string, fields ...string) (InsertBuilder, error) {
	if len(fields) == 0 {
		return nil, errors.New(errorCountFields)
	}
	if r.maxCountArgs > 0 && r.maxCountArgs/len(fields) == 0 {
		return nil, errors.New(errorCountLines)
	}
	maxCountLines := r.maxCountLines
	if r.maxCountArgs > 0 && r.maxCountArgs/len(fields) < maxCountLines || maxCountLines == 0 {
		maxCountLines = r.maxCountArgs / len(fields)
	}
	return &insertBuilder{maxCountLines: maxCountLines, table: table, fields: fields}, nil
}
