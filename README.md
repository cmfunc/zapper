# zapper

based in go.uber.org/zap

zapper is implement of io.Writer and simply config.

## 性能对比

```shell
go test -run=BenchmarkZaperInfo -bench=BenchmarkZaperInfo -benchmem  -count=40 > zaper_benchmark_channel_count40.txt

```

## bug

1. 日志打印错乱
2. 日志未按时切割

## TODO

- [ ] 日志着色
- [ ] 结构体字段内存对齐
- [ ] Writer + Closer 实现拆开
- [ ] FileRotator 的滚动时间以整点为分割
- [ ] README.md 添加使用教程
