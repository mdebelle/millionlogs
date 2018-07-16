Millionlogs Server
===

### requirement
go version go1.10.2 or greater

### get started
```
$>go get github.com/mdebelle/millionlogs
$>cd $GOPATH/src/github.com/mdebelle/millionlogs
$>go install
$>$GOPATH/bin/millionlogs --help
Usage of $GOPATH/bin/millionlogs:
  -port string
        select your favorite port to expose the api (default ":8080")
  -preload
        preloading all data for faster requests but longer init
  -prerank
        preranking all data for faster requests but longer init
  -sample
        set your sample (.tsv) file path location (default "sample/small.tsv")
$>$GOPATH/bin/millionlogs -port :8080 -preload -prerank -sample "sample/hn_logs.tsv"
server init...
selected file sample/hn_logs.tsv
initialisation took 57.151083374s
server ready on port :8080
[GET] /1/queries/popular/2015?size=2000 took 5.672263ms
^C
$>$GOPATH/bin/millionlogs -sample "sample/hn_logs.tsv"
server init...
selected file sample/hn_logs.tsv
initialisation took 28.775µs
server ready on port :8080
[GET] /1/queries/popular/2015?size=2000 took 7.013133889s
[GET] /1/queries/popular/2015?size=2000 took 1.653882ms

```

The server package handles requests on `http://localhost:8080/` *by default*

### examples

GET `http://localhost:8080/1/queries/count/2015`

response `200 OK`
```json
{
    "count": 573697
}
```

GET `http://localhost:8080/1/queries/popular/2015?size=3`

response `200 OK`
```json
[
    {
        "query": "http%3A%2F%2Fwww.getsidekick.com%2Fblog%2Fbody-language-advice",
        "count": 6675
    },
    {
        "query": "http%3A%2F%2Fwebboard.yenta4.com%2Ftopic%2F568045",
        "count": 4652
    },
    {
        "query": "http%3A%2F%2Fwebboard.yenta4.com%2Ftopic%2F379035%3Fsort%3D1",
        "count": 3100
    }
]
```

### note

**date** format should be one of:
- YYYY-MM-DD hh:mm:ss
- YYYY-MM-DD hh:mm
- YYYY-MM-DD hh
- YYYY-MM-DD
- YYYY-MM
- YYYY

**size** format should be a decimal number bigger than 0 (roman number are not tolerate)

### tests
```
$>go get github.com/mdebelle/millionlogs
$>cd $GOPATH/src/github.com/mdebelle/millionlogs
$>go test
```


