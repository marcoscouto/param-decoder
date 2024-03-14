# Param Decoder

![Go](https://img.shields.io/badge/1.21-100000?style=flat&logo=go&logoColor=FFFFFF&labelColor=76E1FE&color=FFFFFF) ![Version](https://img.shields.io/badge/v0.0.3-100000?style=flat&logo=git&logoColor=FFFFFF&labelColor=76E1FE&color=FFFFFF)

This lib is built to get request query params and deserialize in structs in the easiest way.

## How to install

Add at your project with the command

```
go get github.com/marcoscouto/param-decoder
```

## How to use

To use the lib you need to create a struct with public attributes and with the supported types.

```
type Params struct {
	ID        string    `query:"id"`
	Name      string    `query:"name"`
	CreatedAt time.Time `query:"created_at"`
	Age       int32     `query:"age"`
	Balance   float32   `query:"balance"`
	Verified  bool      `query:"verified"`
	Names     []string  `query:"names"`
	Ages      []int32   `query:"ages"`
	Balances  []float32 `query:"balances"`
	Verifieds []bool    `query:"verifieds"`
}
```

Import in the file that will use

```
decoder "github.com/marcoscouto/param-decoder"
```

Call the method and insert the struct type in generics field.

```
params := decoder.DecodeQueryParams[Params](r.URL.Query())
```

The method will return a `Params` struct with the values from query. If the value is empty in query parameter, the struct will return the default value of type.

## Supported types

- string
- bool
- time.Time
- int
`int, int8, int16, int32, int64`
- uint
`uint, uint8, uint16, uint32, uint64`
- float
`float32, float64`
- slices [string, int, uint, float, bool]

## Full Example

main.go

```
package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	decoder "github.com/marcoscouto/param-decoder"
)

const (
	ContentType     = "Content-Type"
	ApplicationJson = "application/json"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		params := decoder.DecodeQueryParams[Params](r.URL.Query())
		marshal, _ := json.Marshal(params)
		w.Header().Set(ContentType, ApplicationJson)
		w.Write(marshal)
	})
	http.ListenAndServe(":3000", r)
}

type Params struct {
	ID        string    `query:"id"`
	Name      string    `query:"name"`
	CreatedAt time.Time `query:"created_at"`
	Age       int32     `query:"age"`
	Balance   float32   `query:"balance"`
	Verified  bool      `query:"verified"`
	Names     []string  `query:"names"`
	Ages      []int32   `query:"ages"`
	Balances  []float32 `query:"balances"`
	Verifieds []bool    `query:"verifieds"`
}
```

cURL
```
curl --location 'http://localhost:3000?name=some_name&id=some_id&age=99&created_at=2024-01-05T12%3A13%3A14Z&balance=10.11&verified=true&ages=1%2C2%2C3&balances=1.1%2C2.2%2C3.3&verifieds=true%2Cfalse%2Ctrue&names=a%2Cb%2Cc'
```