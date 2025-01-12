package storage

import (
	"fmt"
	"math/rand/v2"
	"os"
	"time"

	"github.com/shirou/gopsutil/disk"
)

const defaultTestSizeGB = 1 // 默认测试大小为1GB

func Test() (int, error) {
	fmt.Println("Storage Test")

	// 获取磁盘总空间大小
	totalSpace, err := getDiskTotalSpace()
	if err != nil {
		return 0, err
	}
	fmt.Printf("Total disk space: %f MB\n", float64(totalSpace)/1024/1024) // 打印总磁盘空间大小（MB）
	// 计算测试大小，即总空间的400分之一
	// 如果总大小是0，使用默认测试大小
	var testSize int64
	if totalSpace == 0 {
		fmt.Println("Total disk space is 0, using default test size of 1GB")
		testSize = int64(defaultTestSizeGB * 1024 * 1024 * 1024)
	} else {
		fmt.Printf("Total disk space: %f MB\n", float64(totalSpace)/1024/1024) // 打印总磁盘空间大小（MB）
	}
	// 创建测试文件
	file, err := os.Create("testfile")
	if err != nil {
		return 0, err
	}
	defer os.Remove("testfile")
	defer file.Close()

	// 写入测试数据
	data := make([]byte, testSize)
	startTime := time.Now() // 开始计时
	_, err = file.Write(data)
	if err != nil {
		return 0, err
	}
	writeTime := time.Since(startTime) // 计算写入耗时
	fmt.Printf("Write time: %v\n", writeTime)
	// 读取测试数据
	_, err = file.Seek(0, 0)
	if err != nil {
		return 0, err
	}
	startTime = time.Now() // 重新开始计时
	buf := make([]byte, testSize)
	_, err = file.Read(buf)
	if err != nil {
		return 0, err
	}
	readTime := time.Since(startTime) // 计算读取耗时
	// 评分逻辑需要根据实际情况设计
	fmt.Printf("Read time: %v\n", readTime)
	score := int(1000 / (writeTime.Seconds() + readTime.Seconds()))

	return score, nil
}

func getDiskTotalSpace() (uint64, error) {
	parts, err := disk.Partitions(true)
	if err != nil {
		return 0, err
	}
	if len(parts) == 0 {
		return 0, fmt.Errorf("no partitions found")
	}

	// 通常我们测试系统盘，这里假设是第一个分区
	diskInfo, err := disk.Usage(parts[0].Mountpoint)
	if err != nil {
		return 0, err
	}

	return diskInfo.Total, nil
}

func randScore() int {
	return int(100 + rand.Float64()*(900)) // 生成100到1000之间的随机分数
}
