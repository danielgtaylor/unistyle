# Unistyle

Unicode text styles for Go:

- 𝗕𝗼𝗹𝗱
- 𝘐𝘵𝘢𝘭𝘪𝘤
- U̲n̲d̲e̲r̲l̲i̲n̲e̲
- S̶t̶r̶i̶k̶e̶t̶h̶r̶o̶u̶g̶h̶
- 𝒞𝓊𝓇𝓈𝒾𝓋ℯ
- 𝔊𝔬𝔱𝔥𝔦𝔠
- And more! See the [docs](https://godoc.org/github.com/danielgtaylor/unistyle)

## How does this work?

These styles are split into two categories:

1. Unicode includes a table of [combining diacritical marks](https://en.wikipedia.org/wiki/Combining_character) that can be used to modify the previous character in a string. This allows for effects like bold and italic as well as underline, overline, and strikethrough.

2. Unicode includes multiple styles for most common characters. Using a character offset replacement mapping you can transform strings into their cursive or other styled equivalents.

## Usage Example

Here is a small example. See the tests for more.

```go
package main

import (
  "fmt"
  "github.com/danielgtaylor/unistyle"
)

func main() {
  // Diacritic examples
  fmt.Println(unistyle.Bold("Bold text"))
  fmt.Println(unistyle.Strikethrough("Strike", unistyle.StrikeStrokeLong))

  // Character mapping examples
  fmt.Println(unistyle.Cursive("Cursive"))
  fmt.Println(unistyle.Fraktur("Gothic"))
}
```

# License

Copyright &copy; 2019 Daniel G. Taylor

https://dgt.mit-license.org/
