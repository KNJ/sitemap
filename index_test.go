package sitemap

import "testing"

func TestEmptyIndex(t *testing.T) {
	idx := NewIndex()
	if idx.urlsets == nil {
		t.Fatalf("New Index is expected to have non-nil urlsets.")
	}

	t.Run("output", func(t *testing.T) {
		err := idx.output(&dummyDriver{})
		if err != nil {
			t.Fatalf("Empty Index is expected to output something with no error.")
		}
	})
}
