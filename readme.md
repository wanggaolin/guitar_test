# 吉他练习工具
### 吉他在线问答
- 支持空弦音问答测试
- 支持和弦级问答测试
- 支持和弦构成音测试

## 选项
```shell
guitar  -h
guitar tools

Usage:
  guitar-tools [flags]

Flags:
  -a, --all       题库随机问答测试 (default true)
  -e, --empty     吉他空弦音测试
  -h, --help      help for guitar-tools
  -l, --level     和弦级数问答测试
  -s, --scale     和弦够成音问答测试
  -v, --version   查询版本号
```

```text
请输入 C调 5级和弦是______: g
回答正确
请输入 G调 C和弦是_____级和弦: 4
回答正确
请输入 G调 D和弦是_____级和弦: 
错误,正确答案是: 5
G 和弦构成音(多个音之间逗号间隔)______: 5,7,2
回答正确
DM 和弦构成音(多个音之间逗号间隔)______: 2,4,6
回答正确
F 和弦构成音(多个音之间逗号间隔)______: 4,6,1
回答正确
第 5 弦空弦音是: 6
错误,正确答案是: A
第 4 弦空弦音是: D
回答正确
第 1 弦空弦音是: E
回答正确
```
