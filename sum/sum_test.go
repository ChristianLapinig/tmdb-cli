package sum

import (
	"testing"
)

func TestSum(t *testing.T) {
	for _, test := range []struct {
		Name     string
		Nums     []int
		Expected int
	}{
		{
			Name:     "2+2+2=6",
			Nums:     []int{2, 2, 2},
			Expected: 6,
		},
		{
			Name:     "100+50+400=550",
			Nums:     []int{100, 50, 400},
			Expected: 550,
		},
		{
			Name:     "20+30=50",
			Nums:     []int{20, 30},
			Expected: 50,
		},
		{
			Name:     "10=10",
			Nums:     []int{10},
			Expected: 10,
		},
		{
			Name:     "Empty input should return 0",
			Nums:     []int{},
			Expected: 0,
		},
	} {
		t.Run(test.Name, func(t *testing.T) {
			actual := Sum(test.Nums...)
			if actual != test.Expected {
				t.Errorf("FAILED - expected %d, got %d", test.Expected, actual)
			}
		})
	}
}
