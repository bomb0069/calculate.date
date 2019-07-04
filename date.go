package date

import "fmt"

type Date struct {
	Day   int
	Month int
	Year  int
}

func (d Date) DiffWith(diffDay Date) string {
	return fmt.Sprintf("%d Days", diffDay.Day-d.Day)
}
