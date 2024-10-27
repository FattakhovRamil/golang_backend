package sum

type Calculator interface {
	SumNumbers(numbers []int) int
}

type SimpleSumCalculator struct{}

func NewSimpleSumCalculator() *SimpleSumCalculator {
	return &SimpleSumCalculator{}
}

func (s *SimpleSumCalculator) SumNumbers(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
