# go-i18n

## Install

```
go get github.com/laipz8200/i18n
```

## Usage

```
cat i18n/zh-cn.yaml

welcome: 欢迎
```

```go
package main

import (
	"fmt"

	"github.com/laipz8200/i18n"
)

func main() {
	t := i18n.Lang("zh-cn").Sprintf("welcome")
	fmt.Println(t)
}
```

for more example see test case.
