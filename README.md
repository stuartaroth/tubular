# tubular

*tubular* is a small abstraction library for writing XLSX files and separated values files using the same data structure. For the XLSX internally it uses [github.com/tealeg/xlsx](https://github.com/tealeg/xlsx).

### Installation

Install the library using `go get`

```sh
$ go get github.com/stuartaroth/tubular
```

### Usage

[GoDoc](https://godoc.org/github.com/stuartaroth/tubular)

*tubular* requires clients to create a slice of Page:

[Page](https://godoc.org/github.com/stuartaroth/tubular#Page)

```go
type Page struct {
  Name string
  Rows [][]interface{}
}
```

For an example of use, see below and the examples directory:

### Example Usage

You can find this code at https://github.com/stuartaroth/tubular/blob/master/examples/xlsx.go

```go
package main

import (
  "fmt"
  "github.com/stuartaroth/tubular"
)

func main() {
  pages := []tubular.Page{
    tubular.Page{
      Name: "Horror",
      Rows: [][]interface{}{
        []interface{}{"Title", "Director", "Year Released"},
        []interface{}{"A Nightmare on Elm Street", "Wes Craven", 1984},
        []interface{}{"Black Christmas", "Bob Clark", 1974},
        []interface{}{"Friday the 13th", "Sean S. Cunningham", 1980},
        []interface{}{"Halloween", "John Carpenter", 1978},
        []interface{}{"The Texas Chain Saw Massacre", "Tobe Hooper", 1974},
      },
    },
    tubular.Page{
      Name: "Sci-Fi",
      Rows: [][]interface{}{
        []interface{}{"Title", "Director", "Year Released"},
        []interface{}{"2001: A Spacey Odyssey", "Stanley Kubrick", 1968},
        []interface{}{"Aliens", "James Cameron", 1986},
        []interface{}{"Back to the Future", "Robert Zemeckis", 1985},
        []interface{}{"Blade Runner", "Ridley Scott", 1982},
        []interface{}{"Star Wars", "George Lucas", 1977},
      },
    },
  }

  filename, err := tubular.WriteXLSXFile("Movies", pages)
  if err != nil {
    fmt.Println("Error writing xlsx file:", err)
    return
  }

  fmt.Println("Successfully wrote xlsx file:", filename)
}
```

### License

BSD-2 clause - see LICENSE for more details