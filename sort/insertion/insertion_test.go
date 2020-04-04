package insertion

import (
	"testing"

	"github.com/regemdeu/algorithms/sort/testutil"
)

func TestSortIntSlice_basic(t *testing.T) {
	t.Parallel()
	testutil.BasicSort(t, SortIntSlice)
}

func TestSortIntSliceFromLeft_basic(t *testing.T) {
	t.Parallel()
	testutil.BasicSort(t, SortIntSliceFromLeft)
}

func Benchmark_SortIntSlice(b *testing.B) {
	testutil.Bench(b, SortIntSlice)
}

func Benchmark_SortIntSliceFromLeft(b *testing.B) {
	testutil.Bench(b, SortIntSliceFromLeft)
}
