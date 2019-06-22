package date

import (
	"testing"
)

func Test_DateDiffWithSameDayShouldBe0Days(t *testing.T) {
	expectedResult := "0 Days"
	dateFrom := Date{Day: 22, Month: 6, Year: 2019}
	dateTo := dateFrom

	actualResult := dateFrom.DiffWith(dateTo)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}
