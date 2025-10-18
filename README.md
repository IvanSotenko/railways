# Railways

Go package for building railway systems.

## Usage example

```go
package main

import (
	"fmt"

	"github.com/eevandeya/railways/station"
    "github.com/shopspring/decimal"
)

func main() {
	Station := station.NewStation("Station_1", decimal.NewFromFloat(98.0), 100)
	Station.TicketsSold = 20
	fmt.Println(Station.TicketsAvailable())
}
```
