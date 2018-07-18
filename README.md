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
  -sample string
        set your sample (.tsv) file path location (default "sample/small.tsv")
$>wc -l sample/hn_logs.tsv
 1623420 sample/hn_logs.tsv
$>$GOPATH/bin/millionlogs -port :8080 -preload -prerank -sample "sample/hn_logs.tsv"
2018/07/17 23:52:48 server init...
2018/07/17 23:52:48 selected file sample/hn_logs.tsv
2018/07/17 23:53:29 initialisation took 40.734273395s
2018/07/17 23:53:29 server ready on port :8080
2018/07/17 23:53:37 [GET] /1/queries/popular/2015?size=2000 took 2.121843ms
2018/07/17 23:53:44 [GET] /1/queries/popular/2015?size=2000 took 2.086858ms
^C
$>$GOPATH/bin/millionlogs -sample "sample/hn_logs.tsv"
2018/07/17 23:47:18 server init...
2018/07/17 23:47:18 selected file sample/hn_logs.tsv
2018/07/17 23:47:18 initialisation took 157ns
2018/07/17 23:47:18 server ready on port :8080
2018/07/17 23:47:31 [GET] /1/queries/popular/2015?size=2000 took 7.343221463s
2018/07/17 23:47:37 [GET] /1/queries/popular/2015?size=2000 took 1.444681ms
^C
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
{
    "queries": [
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
}
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

some tests could fail if the big file 'hn_logs.tsv' is not in the sample folder
