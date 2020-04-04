package bubble

import (
	"testing"

	"github.com/regemdeu/algorithms/sort/testutil"
)

func TestSortIntSlice_basic(t *testing.T) {
	testutil.BasicSort(t, SortIntSlice)
}

func Benchmark_Sort(b *testing.B) {
	testutil.Bench(b, SortIntSlice)
}
