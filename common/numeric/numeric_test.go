package numeric

import "testing"

func TestGcd2(t *testing.T) {
	table := []struct {
		a, b, g int
	}{
		{1, 1, 1},
		{1, 2, 1},
		{2, 1, 1},
		{2, 2, 2},
		{2, 3, 1},
		{3, 2, 1},
		{3, 3, 3},
		{3, 4, 1},
		{4, 3, 1},
	}

	for _, test := range table {
		actual := Gcd2(test.a, test.b)
		if actual != test.g {
			t.Errorf("Gcd2(%d, %d) = %d, expected %d", test.a, test.b, actual, test.g)
		}
	}

}

func TestGcd(t *testing.T) {
	table := []struct {
		values []int
		gcd    int
	}{
		{[]int{1, 1}, 1},
		{[]int{1, 2}, 1},
		{[]int{2, 1}, 1},
		{[]int{2, 2}, 2},
		{[]int{2, 3, 4}, 1},
		{[]int{2, 4, 6, 8}, 2},
		{[]int{3, 6, 6, 9, 12}, 3},
	}
	for _, test := range table {
		actual := Gcd(test.values)
		if actual != test.gcd {
			t.Errorf("Gcd(%v) = %d, expected %d", test.values, actual, test.gcd)
		}
	}
}

func TestLcm2(t *testing.T) {
	table := []struct {
		a, b, l int
	}{
		{1, 1, 1},
		{1, 2, 2},
		{2, 1, 2},
		{2, 2, 2},
		{2, 3, 6},
		{3, 2, 6},
		{3, 3, 3},
		{3, 4, 12},
		{4, 3, 12},
	}

	for _, test := range table {
		actual := Lcm2(test.a, test.b)
		if actual != test.l {
			t.Errorf("Lcm2(%d, %d) = %d, expected %d", test.a, test.b, actual, test.l)
		}
	}
}

func TestLcm(t *testing.T) {
	table := []struct {
		values []int
		lcm    int
	}{
		{[]int{1, 1}, 1},
		{[]int{1, 2}, 2},
		{[]int{2, 1}, 2},
		{[]int{2, 2}, 2},
		{[]int{2, 3, 4}, 12},
		{[]int{2, 4, 6, 8}, 24},
		{[]int{3, 6, 6, 9, 12}, 36},
	}
	for _, test := range table {
		actual := Lcm(test.values)
		if actual != test.lcm {
			t.Errorf("Lcm(%v) = %d, expected %d", test.values, actual, test.lcm)
		}
	}
}
