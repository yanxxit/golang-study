# struct 结构体

## 参考
- https://juejin.im/post/5ca2f37ce51d4502a27f0539

## struct 复用
```go
type CommonConfig struct {
	Port int32
	Host string
}

//特殊配置
type configuration struct {
	CommonConfig
	Enabled bool
	Path    string
	Name    string
	Title   string
}

var Config configuration
Config.Host 可以直接获取CommonConfig的字段
```