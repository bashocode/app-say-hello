# APP Say Hello

This app use the dependency of go from [https://github.com/bashocode/go-say-hello](https://github.com/bashocode/go-say-hello)

- initial module of go
```
go mod init github.com/{username}/{repository} // let say github.com/bashocode/app-say-hello
```
- add file called main.go
```go
package main

import (
	"fmt"

	go_say_hello "github.com/bashocode/go-say-hello"
)

func main() {
	fmt.Println(go_say_hello.SayHello())
}

```
- push to git
```
git add .
git commit -m "initial learn to use module"
git remote add origin git@github.com:{username}/{repository}.git
git push -u origin master
```

How to run app
```
cd app-say-hello
go mod tidy
go run main.go
```
