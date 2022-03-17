package listfns_test

import (
	"testing"

	"github.com/bpar476/listfns"
)

func TestMapSameType(t *testing.T) {
	source := []int{1, 2, 3}

	res := listfns.Map(source, func(in int) int { return in * 2 })

	for i := range source {
		val := source[i]
		result := res[i]
		expected := val * 2
		if expected != result {
			t.Fail()
			t.Logf("Result[%d] was %d. Expected: %d\n", i, result, expected)
		}
	}
}
func TestMapDifferentType(t *testing.T) {
	source := []int{1, 2, 3, 4}
	res := listfns.Map(source, func(in int) bool { return in%2 == 0 })

	for i := range source {
		val := source[i]
		result := res[i]
		expected := val%2 == 0
		if expected != result {
			t.Fail()
			t.Logf("Result[%d] was %t. Expected: %t\n", i, result, expected)
		}
	}
}

func TestFold(t *testing.T) {
	source := []int{1, 2, 3, 4}
	res := listfns.Fold(source, false, func(acc bool, current int) bool {
		return acc || current == 10
	})

	if res {
		t.Fail()
		t.Logf("Result was %t, expected %t\n", res, false)
	}
}

func TestReduce(t *testing.T) {
	source := []int{1, 2, 3, 4}
	sum := 10

	res := listfns.Reduce(source, func(acc, current int) int {
		return acc + current
	})

	if res != sum {
		t.Fail()
		t.Logf("Result was %d, expected %d\n", res, sum)
	}
}

func TestContains(t *testing.T) {
	source := []int{1, 2, 3}
	res := listfns.Contains(source, 2)

	if !res {
		t.Fail()
		t.Logf("Result was false. Expected true\n")
	}
}

func TestDoesNotContain(t *testing.T) {
	source := []int{1, 2, 3}
	res := listfns.Contains(source, 4)

	if res {
		t.Fail()
		t.Logf("Result was true. Expected false\n")
	}
}
