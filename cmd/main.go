package main

import (
	"flag"
	"fmt"
	"log"
	"nas-benchmark/pkg/cpu"
	"nas-benchmark/pkg/memory"
	"nas-benchmark/pkg/network"
	"nas-benchmark/pkg/storage"
)

func main() {
	// 添加命令行参数解析
	help := flag.Bool("h", false, "显示帮助信息")
	flag.Parse()

	if *help {
		fmt.Println("使用说明：")
		fmt.Println("  -h: 显示帮助信息")
		return
	}
	fmt.Println("NAS Benchmark Tool")

	// CPU Test
	cpuScore, err := cpu.Test()
	if err != nil {
		log.Fatalf("Failed to test CPU: %v", err)
	}
	fmt.Printf("CPU Score: %d\n", cpuScore)

	// Memory Test
	memoryScore, err := memory.Test()
	if err != nil {
		log.Fatalf("Failed to test Memory: %v", err)
	}
	fmt.Printf("Memory Score: %d\n", memoryScore)

	// Storage Test
	storageScore, err := storage.Test()
	if err != nil {
		log.Fatalf("Failed to test Storage: %v", err)
	}
	fmt.Printf("Storage Score: %d\n", storageScore)

	// Network Test
	networkScore, err := network.Test()
	if err != nil {
		log.Fatalf("Failed to test Network: %v", err)
	}
	fmt.Printf("Network Score: %d\n", networkScore)
}
