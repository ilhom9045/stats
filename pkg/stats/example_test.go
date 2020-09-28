package stats

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ilhom9045/bank/v2/pkg/types"
)

func TestCategoriesAvgUser(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 1_000_00},
		{ID: 2, Category: "auto", Amount: 1_000_00},
		{ID: 3, Category: "food", Amount: 3_000_00},
		{ID: 4, Category: "auto", Amount: 4_000_00},
		{ID: 5, Category: "fun", Amount: 5_000_00},
	}
	expected := map[types.Category]types.Money{
		"auto": 2_000_00,
		"food": 3_000_00,
		"fun":  5_000_00,
	}

	result := CategoriesAvg(payments)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestPeriodsDynamicUser(t *testing.T) {
	first := map[types.Category]types.Money{
		"auto": 1_000_00,
		"food": 2_000_00,
	}
	second := map[types.Category]types.Money{
		"auto": 2_000_00,
		"food": 3_000_00,
		"m":    1_000_00,
	}
	expected := map[types.Category]types.Money{
		"auto": 1_000_00,
		"food": 1_000_00,
		"m":    1_000_00,
	}

	result := PeriodsDynamic(first, second)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}
