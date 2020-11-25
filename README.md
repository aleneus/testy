# testy

Helpers for reducing the amount of code in tests

## Example

```
import (
    "github.com/aleneus/testy"
    "testing"
)

func TestSomething(t *testing.T) {
    testy.AssertEqual(t, 2*2, 4)
}
```
