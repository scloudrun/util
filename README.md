# util

`util` golang Universal tool kit 


## Installation

Use [`go get`](https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies) to install and update:

```sh
$ go get -u github.com/scloudrun/util
```

## Quick start
 
```sh
# assume the following codes in example.go file
$ cat example.go
```

```go
package main

import "github.com/scloudrun/util"

func main() {
    fmt.Println(util.Md5("util"))
    fmt.Println(util.Int64("1234"))
}
```

```
# run example.go
$ go run example.go
```

## Todo
- go test
- extend
