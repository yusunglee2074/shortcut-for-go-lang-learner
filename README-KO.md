# Shortcut-for-go-lang-learner
I will make shortcut for someone who wants to reach a level where they can say "I've dealt with the language"


[한국어](./README-ko.md)

## Our journey

- Install(Drawing Go charactor during installation)
- What is ```package main```?
- Variable
- Run script
- Function
- Loop
- Stdout, Stdin
- Make simplest cli game.
- Where do I go next

### Install(Drawing Go charactor during installation)
!['내 그림'](static/drawing.jpg "고 캐릭터")

### What is ```package main```
**Package** keyword is using describe script name and namespace, It can be imported from another package script.  
**Any code written in Go belongs to a package.**  
You can Import another package with ```import "{packageName}"```   
```main``` package has special function called the ```main``` function. It is entry point to excute your program. And main function has zero arguments nor not return any value.


### Variable  
Go is strongly, statically **typed language**.  
Variables are used by the **complier**. 

```go
package main

func main() {
  var n int = 3
  const s string = "Power"
  const b bool = false
}

```

### Run script
You can make script, how to run this script?
```bash
go run {filename}
```

### Function
You can declare function with ```func``` keyword.

```go
package main

import ("math/rand")

func main() {
  const doIDateOnWeekend bool = canIDateOnThisWeekend()
  const doIDateOnThisMonth bool = canIDate(23)
}

func canIDateOnThisWeekend() bool {
  return false
}

func canIDate(luckyNumber int) bool {
  var rullet []bool = {false, false, false, false, false, false, true}
  rand.Seed(luckyNumber)

  return rullet[rand.Intn(len(rullet) - 1)]
}
```

### Loop
Go has only one looping construct, the ```for``` loop.

```go
package main

func main() {
	var sum int = 0
	for var i = 0; i < 10; i++ {
		sum += i
	}
}
```
### Stdout, Stdin
Go has built-in library called ```standard library```.  
It has bunch of extended function for language. We use ```fmt package``` for Stdin, Stdout, this package implements formatted I/O with functions is analogous to C's printf and scanf.

```go
package main

import (
  "fmt"
)

func main() {
  var i interface{} = 23
  fmt.Printf("%v\n", i)
}
```



### Make simplest cli game.


### Where do I go next









