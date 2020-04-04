package merge

import (
	"testing"

	"github.com/regemdeu/algorithms/sort/testutil"
)

func TestSortIntSlice_basic(t *testing.T) {
	t.Parallel()
	testutil.BasicSort(t, SortIntSlice)
}

func Benchmark_SortIntSlice(b *testing.B) {
	testutil.Bench(b, SortIntSlice)
}
