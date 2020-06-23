# Golang 한글 자모 분리

- 입력된 문장에서 한글을 자모로 분리해줍니다.

## Example
```go
package main

import (
	"fmt"

	"github.com/sucream/koreanutil"
)

func main() {
    fmt.Println(koreanutil.GetJamo("안hi녕하세世界요"))
    // [[ㅇ ㅏ ㄴ] [h] [i] [ㄴ ㅕ ㅇ] [ㅎ ㅏ ] [ㅅ ㅔ ] [世] [界] [ㅇ ㅛ ]]
}

```

## Install
- `go get github.com/sucream/koreanutil`

## License
- MIT License