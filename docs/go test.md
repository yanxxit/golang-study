# go test

## 常用 go test
1. go test 将对当前目录下的所有*_test.go文件进行编译并自动运行测试。
2. 测试某个文件使用”-file”参数。go test –file *.go 。
> 例如：go test -file mysql_test.go，"-file"参数不是必须的，可以省略，如果你输入go test b_test.go也会得到一样的效果。
3. 


## 执行某个方法
> go test -v -run TestHome

执行 当前目录下 方法名为 TestHome* 的测试案例

### Go test 测试单个用例
通过 **go help testflag** 可以看出，可以使用 test.run 加正在表达式来指定要测试的用例。
如有以下用例在包 foo/test 中：
```go
func TestFoo(t *testing.T){}
func TestFoo1(t *testing.T){}
func TestFooAdd(t *testing.T){}
func TestFooCdd(t *testing.T){}
func TestAbc(t *testing.T){}
func TestDef(t *testing.T){}

```


* **go test -v foo/test** 将测试所有的用例 
* **go test -v -run TestFoo foo/test** 将测试 TestFoo 和 TestFoo1
* **go test -v -run ^TestFoo$ foo/test** 将只测试 TestFoo

## 备注
- **go test -v** 表示输出详细信息，无论成功失败,不加"-v"表示只显示未通过的用例结果
- 创建benchmark性能测试用例文件mysql_b_test.go(文件名必须是*_b_test.go的类型，*代表要测试的文件名，函数名必须以Benchmark开头如：BenchmarkXxx或Benchmark_xxx)
> 进行所有go文件的benchmark测试 go test -bench=".*" 或 go test . -bench=".*"
> 对某个go文件进行benchmark测试 go test mysql_b_test.go -bench=".*"

## 引用
- [golang test测试使用](https://studygolang.com/articles/2491)