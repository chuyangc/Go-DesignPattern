package StrategyPattern

import "testing"

func TestStrategy_Do(t *testing.T) {
	context := &Context{}

	// 执行策略1
	strategy1 := &Strategy1{}
	context.Strategy = strategy1
	context.Do()

	// 执行策略2
	strategy2 := &Strategy2{}
	context.Strategy = strategy2
	context.Do()
}