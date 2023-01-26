# gormlog

[![Go Reference](https://pkg.go.dev/badge/github.com/petoem/gormlog.svg)](https://pkg.go.dev/github.com/petoem/gormlog)

Zerolog logging for gorm

## Installation

```sh
go get -u github.com/petoem/gormlog
```

## Usage

```go
package main

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/petoem/gormlog"
)

func main() {
	logger := gormlog.NewLogger(log.Logger)
	db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{Logger: logger})
	// ...
}
```

## Contributing

If you wish to contribute to the code or documentation, feel free to fork the repository and submit a pull request.

## License

Licensed under the [MIT license](LICENSE).
