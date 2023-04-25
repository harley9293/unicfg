# unicfg
[!https://github.com/harley9293/unicfg/workflows/Go/badge.svg]

unicfg is an open-source library written in Go that aims to parse various configuration file formats and convert them to a unified method linkage interface.

currently supported configuration formats:

- json
- ini

## Installation

Install unicfg using the go get command:

```shell
go get -u github.com/harley9293/unicfg
```

## Usage

First, import the unicfg library:

```go
import "github.com/harley9293/unicfg"
```

Then, parse the configuration file using `New`:

```go
elem, err := unicfg.New("path/to/your/config/file")
if err != nil {
    log.Fatalf("Error parsing config: %v", err)
}

```

Retrieve configuration information using the unified interface:

```go
// Get basic data types
s := elem.Key("test1").String()
i := elem.Key("test2.test3").MustInt(100)
b := elem.Key("test4").Key("test5").Bool()

// Get array data types
for i := elem.Key("test6").Next(); i != nil; i = i.Next() {
 // ...
}

// Get map data types
for k, v := range elem.Children() {
 // ...
}
```
