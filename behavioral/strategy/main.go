package main

import (
	"design_patterns/behavioral/strategy/avgcalculator"
	"fmt"
)

func main() {
	numbers := []float64{1, 2, 3, 4, 5}

	// 创建简单平均值计算器
	calculator := avgcalculator.SampleAvg{}
	fmt.Println("Simple average:", calculator.Calculate(numbers))

	// 创建加权平均值计算器
	calculator2 := avgcalculator.WeightAvg{}
	fmt.Println("Weighted average:", calculator2.Calculate(numbers))
}
