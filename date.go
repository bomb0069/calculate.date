package date

type Date struct {
	Day   int
	Month int
	Year  int
}

func (d Date) DiffWith(diffDay Date) string {
	return "0 Days"
}
