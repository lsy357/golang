package collection

import (
	"sync"
	"testing"
)

func TestSyncMap(t *testing.T) {
	test := sync.Map{}

	test.Store("test_key", "test_value")
}
