
```markdown
# NAS Benchmark Tool

NAS Benchmark Tool 是一个用于评估网络附加存储（NAS）性能的工具，包括 CPU、内存、存储和网络性能测试。

## 功能

- CPU性能测试
- 内存性能测试
- 存储性能测试
- 网络性能测试

## 编译方法

### 环境要求

- Go 语言环境（版本 1.18 或更高）

### 步骤

1. 克隆项目到本地：

   ```bash
   git clone https://github.com/yourusername/nas-benchmark.git
   cd nas-benchmark
   ```

2. 确保你的工作目录是项目的根目录。

3. 运行构建命令：

   ```bash
   go build -o nas-benchmark ./cmd
   ```

   这将在当前目录下生成一个名为 `nas-benchmark` 的可执行文件。

4. 运行可执行文件：

   ```bash
   ./nas-benchmark
   ```

   或者在Windows系统上：

   ```bash
   nas-benchmark.exe
   ```

## 使用方法

运行 `nas-benchmark` 可执行文件后，程序将自动执行各项性能测试，并输出每个部分的评分结果。

## 贡献

欢迎提交 Pull Request 或者在 Issues 中提出建议和报告问题。

## 许可证

本项目采用 [MIT License](LICENSE)。
```



