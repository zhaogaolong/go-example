## 原生的单元测试

#### 1 首先创建一个 `calc.go` 的文件。

假如这就是我们的业务代码，增加了 Sum 方法

```go
// calc.go
package native_uniitest

func Sum(firstNum, secondNum int) int {
	return firstNum + secondNum
}

```

#### 2 创建单测文件 `calc_test.go`

单侧文件创建规则：与源文件同名后加 `_test.go`的文件。

```go
// calc_test.go
package native_uniitest

import "testing"

func TestSum(t *testing.T) {
	rightResult := 8
	result := Sum(3, 5)
	if result != rightResult {
		t.Errorf("error, require %d, got %d", rightResult, result)
	}
}

```

#### 3 运行单侧

    $ go test -v calc_test.go calc.go
    === RUN   TestSum
    --- PASS: TestSum (0.00s)
    PASS
    ok      command-line-arguments  0.005s

#### 4 单侧覆盖率

##### -cover 显示单侧覆盖率

    $ go test -v calc_test.go calc.go -cover
    === RUN   TestSum
    --- PASS: TestSum (0.00s)
    PASS
    coverage: 100.0% of statements
    ok      command-line-arguments  0.004s  coverage: 100.0% of statements

##### -coverprofile 保存成文件

    $ go test -v calc_test.go calc.go -coverprofile=size_coverage.out
    === RUN   TestSum
    --- PASS: TestSum (0.00s)
    PASS
    coverage: 100.0% of statements
    ok      command-line-arguments  0.004s  coverage: 100.0% of statements

##### cover -func 查看每一个 function 的代码覆盖率

    $ go tool cover -func=size_coverage.out
    /Users/zhaogaolong/go/src/github.com/zhaogaolong/go-example/go_test/unittest/native_uniitest/calc.go:3: Sum             100.0%
    total:                                                                                                  (statements)    100.0%

#### cover -html 浏览器查看代码覆盖率

    $ go tool cover -html=size_coverage.out
