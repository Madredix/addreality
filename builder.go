package addreality

import (
	"errors"
	"strings"
	"sync"
)

const errorCountArgs = `number of arguments must equal the number of fields`

type InsertBuilder interface {
	Append(args ...interface{}) error
	ToSQL() []BatchQuery
}

type insertBuilder struct {
	maxCountLines int
	table         string
	fields        []string
	args          []interface{}
	mu            sync.Mutex
}

type BatchQuery struct {
	Query string
	Args  []interface{}
}

func (i *insertBuilder) Append(args ...interface{}) error {
	if len(i.fields) != len(args) {
		return errors.New(errorCountArgs)
	}
	i.mu.Lock()
	defer i.mu.Unlock()
	i.args = append(i.args, args...)
	return nil
}

func (i *insertBuilder) ToSQL() []BatchQuery {
	i.mu.Lock()
	defer i.mu.Unlock()
	countFields := len(i.fields)
	countArgs := len(i.args)
	maxCountLines := i.maxCountLines
	if maxCountLines == 0 {
		maxCountLines = countArgs / countFields
	}

	countBatchs := countArgs / countFields / maxCountLines
	if countArgs/countFields%maxCountLines > 0 {
		countBatchs++
	}
	countLines := countArgs / countFields

	batchQueries := make([]BatchQuery, countBatchs)
	fieldsCondition := `, (` + strings.Repeat(`, ?`, countFields)[2:] + `)`
	for j := range batchQueries {
		currentCountLines := maxCountLines
		if countLines-j*maxCountLines < currentCountLines {
			currentCountLines = countLines - j*maxCountLines
		}
		batchQueries[j].Query = `INSERT INTO "` + i.table + `" ("` + strings.Join(i.fields, `", "`) + `") VALUES ` + strings.Repeat(fieldsCondition, currentCountLines)[2:]
		batchQueries[j].Args = i.args[maxCountLines*countFields*j : maxCountLines*countFields*j+currentCountLines*countFields]
	}
	i.args = make([]interface{}, 0)

	return batchQueries
}
