package cpu

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

// Test 用于测试综合 CPU 性能，包括整数、浮点运算和多核并行处理
func Test() (int, error) {
	// 记录开始时间
	start := time.Now()

	// 设置运算次数
	const numIterations = 100000000
	// 获取CPU核心数
	numCPU := runtime.NumCPU()
	// 获取CPU线程数
	numThreads := runtime.GOMAXPROCS(0)
	fmt.Printf("CPU核心数：%d，线程数：%d\n", numCPU, numThreads)
	// 使用 WaitGroup 管理并发任务
	var wg sync.WaitGroup
	numGoroutines := numThreads // 假设测试线程数并发任务

	// 使用通道收集每个 goroutine 的处理时长
	resultChannel := make(chan time.Duration, numGoroutines)

	// 并发执行多个 goroutine 测试
	for g := 0; g < numGoroutines; g++ {
		wg.Add(1)
		go func(goroutineIndex int) {
			defer wg.Done()
			duration := runCPUTask(goroutineIndex, numIterations)
			resultChannel <- duration
		}(g)
	}

	// 等待所有 goroutine 执行完毕
	wg.Wait()
	close(resultChannel)

	// 计算总时间
	var totalDuration time.Duration
	for duration := range resultChannel {
		totalDuration += duration
	}

	// 计算平均时间
	averageDuration := totalDuration / time.Duration(numGoroutines)

	// 输出执行时间
	fmt.Printf("CPU Test took %v (average across %d goroutines)\n", averageDuration, numGoroutines)

	// 计算测试的总时长（使用 start 变量）
	totalTestDuration := time.Since(start)
	fmt.Printf("Total Test Duration: %v\n", totalTestDuration)

	// 根据执行时间评估得分（这里的得分可以根据实际需要设计）
	score := int(averageDuration.Milliseconds())
	score = 10000 - score
	return score, nil
}

// runCPUTask 执行具体的 CPU 密集型任务，包括整数运算、浮点运算和多种运算
func runCPUTask(goroutineIndex, iterations int) time.Duration {
	// 测试任务的开始时间
	start := time.Now()

	// 执行加减乘除和浮点数运算的混合计算
	for i := 0; i < iterations; i++ {
		// 整数运算
		_ = i * 2
		_ = i + 1
		_ = i - 1

		// 浮点运算，测试浮点计算能力
		_ = math.Sqrt(float64(i))    // 浮点平方根运算
		_ = math.Pow(float64(i), 2)  // 浮点幂运算
		_ = math.Sin(float64(i))     // 浮点正弦运算
		_ = math.Cos(float64(i))     // 浮点余弦运算
		_ = math.Log(float64(i + 1)) // 浮点对数运算
		_ = float64(i) * 2.71828     // 常见的浮点常数（e）的乘法运算
	}

	// 返回该任务的执行时长
	return time.Since(start)
}
