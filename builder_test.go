package addreality

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestInsertBuilder_Append(t *testing.T) {
	i := getTestInsertBuilder()
	err := i.Append(15, 25, 28)
	if err == nil || err.Error() != errorCountArgs {
		t.Errorf(`missed error, received: %#v\n`, err)
	}

	err = i.Append(15, 25)
	if err != nil {
		t.Errorf(`received error, received: %#v\n`, err)
	}
}

func TestInsertBuilder_ToSQL(t *testing.T) {
	i := getTestInsertBuilder()
	i.Append(`text1`, 1)
	i.Append(`text2`, 2)
	i.Append(`text3`, 3)

	expected := []BatchQuery{
		{
			Query: `INSERT INTO "table" ("field1", "field2") VALUES (?, ?), (?, ?)`,
			Args:  []interface{}{`text1`, 1, `text2`, 2},
		},
		{
			Query: `INSERT INTO "table" ("field1", "field2") VALUES (?, ?)`,
			Args:  []interface{}{`text3`, 3},
		},
	}
	for j, row := range i.ToSQL() {
		if expected[j].Query != row.Query || !cmp.Equal(expected[j].Args, row.Args) {
			t.Errorf("row: %d\nexpected query: %s\nreceived query: %s\nexpected args: %v\nreceived args: %v\n", j, expected[j].Query, row.Query, expected[j].Args, row.Args)
		}
	}

	i.Append(`text4`, 4)
	i.Append(`text5`, 5)
	expected = []BatchQuery{
		{
			Query: `INSERT INTO "table" ("field1", "field2") VALUES (?, ?), (?, ?)`,
			Args:  []interface{}{`text4`, 4, `text5`, 5},
		},
	}
	for j, row := range i.ToSQL() {
		if expected[j].Query != row.Query || !cmp.Equal(expected[j].Args, row.Args) {
			t.Errorf("row: %d\nexpected query: %s\nreceived query: %s\nexpected args: %v\nreceived args: %v\n", j, expected[j].Query, row.Query, expected[j].Args, row.Args)
		}
	}

	rows := i.ToSQL()
	if !cmp.Equal(rows, []BatchQuery{}) {
		t.Errorf(`received data, expected <nil>\nreceived: %+v\n`, rows)
	}
}

func getTestInsertBuilder() InsertBuilder {
	return &insertBuilder{
		maxCountLines: 2,
		table:         `table`,
		fields:        []string{`field1`, `field2`},
	}
}
