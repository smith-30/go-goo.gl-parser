[![Build Status](https://secure.travis-ci.org/smith-30/go-goo.gl-parser.png?branch=master)](http://travis-ci.org/smith-30/go-goo.gl-parser)
[![Coverage Status](https://coveralls.io/repos/github/smith-30/go-goo.gl-parser/badge.svg?branch=master)](https://coveralls.io/github/smith-30/go-goo.gl-parser?branch=master)
[![GoDoc](https://godoc.org/github.com/smith-30/go-goo.gl-parser?status.svg)](https://godoc.org/github.com/smith-30/go-goo.gl-parser)
[![license](https://img.shields.io/badge/license-MIT-4183c4.svg)](https://github.com/smith-30/go-goo.gl-parser/blob/master/LICENSE)

 ### go-goo.gl-parser

__ex.)__

```go
package main

import (
	"fmt"
	parser "github.com/smith-30/go-goo.gl-parser"
)

func main() {
	p := parser.NewParser("your api token")
	d, err := p.DecodeURL("https://goo.gl/kMxDaw")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	fmt.Println(d) // https://github.com/smith-30/go-goo.gl-parser
}
```

### requirements

- URL Shortener API's token  
Please get it from console.developers.google.com 

### License

go-goo.gl-parser is licensed under the the MIT license. Please see the LICENSE file for details.