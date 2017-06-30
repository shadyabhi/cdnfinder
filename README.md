[![Build Status](https://travis-ci.org/shadyabhi/cdnfinder.png)](https://travis-ci.org/shadyabhi/cdnfinder)
[![codecov](https://codecov.io/gh/shadyabhi/cdnfinder/branch/master/graph/badge.svg)](https://codecov.io/gh/shadyabhi/cdnfinder)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/shadyabhi/cdnfinder)

# Summary

* This program gets the list of public nameservers from https://public-dns.info/.
* Picks one nameserver from every country. 
* Do a CNAME query for the specified domain against these picked nameservers.

## Things to note

* Default timeout is 3 seconds

# Sample Output

* -number-ns=5: Will set the number of nameservers to query as 5
* -timeout=0.1s: Sets DNS timeout as 0.1s
* -filter-countries=IN,SG,JP: Sets the countries to query CNAME in
* -domain: Set the domain to query CNAME for

```
➜ $?=0 shadyabhi/cdnfinder [ 4:48PM] (master|…) % ./cdnfinder -domain=static.licdn.com -timeout=0.1s -number-ns=5 -filter-countries=IN,SG,JP
WARN[0000] Skipping a line while parsing csv at line: 46346, column: 0
|-----------|----------------------|-----------------------------------------|----------------|-------------|
|  COUNTRY  |         CITY         |              CDN/HOSTNAME               |       IP       |    TIME     |
|-----------|----------------------|-----------------------------------------|----------------|-------------|
| Singapore | SG - Singapore       | EdgeCast   - cs627.wac.epsiloncdn.net.  | 202.136.162.11 | 44.04332ms  |
| India     | IN - Mumbai          | EdgeCast   - cs627.wac.epsiloncdn.net.  | 203.115.71.66  | 20.700975ms |
| India     | IN - Bhubaneswar     | EdgeCast   - cs627.wac.epsiloncdn.net.  | 202.62.224.2   | 61.284302ms |
| India     | IN - Mumbai          | EdgeCast   - cs627.wac.epsiloncdn.net.  | 202.138.120.4  | 27.423725ms |
| Singapore | SG -                 | Akamai     - e9706.dscg.akamaiedge.net. | 121.52.206.130 | 44.463913ms |
|-----------|----------------------|-----------------------------------------|----------------|-------------|
INFO[0000] Took 669.099039ms to query DNS servers...

>>>  0s elasped...
➜ $?=0 shadyabhi/cdnfinder [ 4:49PM] (master|…) %
```

# TODO

* Limit goroutines for DNS queries
* Option to force refresh nameserver.csv
