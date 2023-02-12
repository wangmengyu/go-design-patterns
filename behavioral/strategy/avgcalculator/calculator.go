package avgcalculator

type AvgCalculator interface {
	// 平均数计算器接口
	Calculate(numbers []float64) float64
}

type SampleAvg struct {
	// 普通平均数计算
}

func (sa *SampleAvg) Calculate(numbers []float64) float64 {
	sum := float64(0)
	for _, num := range numbers {
		sum += num
	}

	return sum / float64(len(numbers))
}

type WeightAvg struct {
	// 加权平均数计算

}

func (wa *WeightAvg) Calculate(numbers []float64) float64 {
	var sum float64
	var weightSum float64
	for i, num := range numbers {
		sum += num * float64(i+1)
		weightSum += float64(i + 1)
	}
	return sum / weightSum
}
