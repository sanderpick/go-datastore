package sync

import (
	"testing"

	ds "github.com/textileio/go-datastore"
	dstest "github.com/textileio/go-datastore/test"
)

func TestSync(t *testing.T) {
	dstest.SubtestAll(t, MutexWrap(ds.NewMapDatastore()))
}
