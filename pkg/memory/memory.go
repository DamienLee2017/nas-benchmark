package memory

import (
	"fmt"
	"math/rand"
	"runtime"
	"syscall"
	"time"
)

// Test 用于测试内存性能，动态根据系统内存大小分配内存进行测试
func Test() (int, error) {
	// 获取系统内存大小
	var sysInfo syscall.Sysinfo_t
	if err := syscall.Sysinfo(&sysInfo); err != nil {
		return 0, fmt.Errorf("failed to get system memory info: %v", err)
	}

	// 获取系统总内存大小，单位为字节
	totalMemory := sysInfo.Totalram * uint64(syscall.Getpagesize()) // 获取总内存大小

	// 转换为MB单位
	totalMemoryMB := totalMemory / 1024 / 1024

	// 设定测试内存的最大值为 1GB，避免设备内存过小导致分配失败
	maxTestMemoryMB := 1024                // 最大分配内存为 1GB
	testMemoryMB := int(totalMemoryMB / 8) // 按比例分配，最多测试 1GB
	if testMemoryMB > maxTestMemoryMB {
		testMemoryMB = maxTestMemoryMB // 限制最大测试内存为 1GB
	}

	// 为测试分配内存
	memSize := testMemoryMB * 1024 * 1024 // 转换为字节
	memArray := make([]byte, memSize)

	// 输出系统总内存和分配的测试内存大小
	fmt.Printf("Total System Memory: %d MB\n", totalMemoryMB)
	fmt.Printf("Allocated Test Memory: %d MB\n", testMemoryMB)

	// 模拟内存写入操作
	startWrite := time.Now()
	for i := 0; i < len(memArray); i++ {
		memArray[i] = byte(rand.Intn(256)) // 随机写入内存
	}
	writeDuration := time.Since(startWrite)

	// 模拟内存读取操作
	startRead := time.Now()
	for i := 0; i < len(memArray); i++ {
		_ = memArray[i] // 读取内存
	}
	readDuration := time.Since(startRead)

	// 输出内存相关统计信息
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Memory Allocated: %d MB\n", memStats.Alloc>>20) // 已分配内存
	fmt.Printf("Write Duration: %v\n", writeDuration)           // 写入时间
	fmt.Printf("Read Duration: %v\n", readDuration)             // 读取时间

	// 评分逻辑：假设理想的写入时间为 100ms，读取时间为 50ms
	const idealWriteTime = 100 * time.Millisecond * 1000
	const idealReadTime = 1000 * time.Millisecond

	writeScore := int((idealWriteTime - writeDuration) / time.Millisecond)
	readScore := int((idealReadTime - readDuration) / time.Millisecond)
	fmt.Printf("Write Score: %v\n", writeScore) // 写入时间
	fmt.Printf("Read Score: %v\n", readScore)   // 读取时间
	// 防止得分为负数
	if writeScore < 0 {
		writeScore = 0
	}
	if readScore < 0 {
		readScore = 0
	}

	// 对写入和读取得分进行归一化
	normalizedWriteScore := float64(writeScore)
	normalizedReadScore := float64(readScore)

	//
	finalScore := int(normalizedWriteScore + normalizedReadScore)

	// 确保最终得分不为负值
	if finalScore < 0 {
		finalScore = 0
	}

	// 输出最终得分
	fmt.Printf("Final Memory Test Score: %d\n", finalScore)

	return finalScore, nil
}
