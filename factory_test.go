package addreality

import "testing"

func TestNewInsertBuilderFactory(t *testing.T) {
	_, err := NewInsertBuilderFactory(-1, 1)
	if err == nil || err.Error() != errorParams {
		t.Errorf(`missed error, received: %#v\n`, err)
	}

	_, err = NewInsertBuilderFactory(1, -1)
	if err == nil || err.Error() != errorParams {
		t.Errorf(`missed error, received: %#v\n`, err)
	}

	_, err = NewInsertBuilderFactory(1, 1)
	if err != nil {
		t.Errorf(`received error: %#v\n`, err)
	}
}

func TestFactory_CreateInsertBuilder(t *testing.T) {
	f := &factory{maxCountLines: 2, maxCountArgs: 3}

	_, err := f.CreateInsertBuilder(`table`)
	if err == nil || err.Error() != errorCountFields {
		t.Errorf(`missed error, received: %#v\n`, err)
	}

	_, err = f.CreateInsertBuilder(`table`, `field1`, `field2`, `field3`, `field4`)
	if err == nil || err.Error() != errorCountLines {
		t.Errorf(`missed error, received: %#v\n`, err)
	}

	_, err = f.CreateInsertBuilder(`table`, `field1`, `field2`)
	if err != nil {
		t.Errorf(`received error, received: %#v\n`, err)
	}
}
