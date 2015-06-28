# mpk
Small Go library that helps to fetch data about public transport in [Wrocław](https://pl.wikipedia.org/wiki/Wroc%C5%82aw). Full list of transportation types is available [here](http://pasazer.mpk.wroc.pl/jak-jezdzimy/mapa-pozycji-pojazdow).

## Usage

```go
package main

import (
    "github.com/piotrkowalczuk/mpk"
)

func main() {
    service, err := mpk.New(&http.Client{})
    if err != nil {
        log.Fatal(err)
    }
    
    locations, err := service.GPS.Fetch(map[mpk.TransportationType][]string{
        mpk.TransportationTypeBus: []string{
            "a",
            "602",
            "710",
        },
        mpk.TransportationTypeTram: []string{
            "0L",
            "8",
        },
    })
}
```
