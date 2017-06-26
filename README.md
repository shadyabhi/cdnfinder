# Summary

* This program gets the list of public nameservers from https://public-dns.info/.
* Picks one nameserver from every country. 
* Do a CNAME query for the specified domain against these picked nameservers.

## Things to note

* Default timeout is 3 seconds

# Sample Output

This output is when timeout for DNS requests is set to 0.1ms. Set it to higher amount to get more results.
Default value of timeout is 3 seconds.

```
➜ $?=0 shadyabhi/cdnfinder [ 3:45PM] (master↑1|…) % go build && ./cdnfinder -domain=static.licdn.com -timeout=0.1s
|-----------|----------------|----------------------------------------|-------------|
|  COUNTRY  |      CITY      |              CDN/HOSTNAME              |    TIME     |
|-----------|----------------|----------------------------------------|-------------|
| Singapore | SG - Singapore | EdgeCast   - cs627.wac.epsiloncdn.net. | 52.727914ms |
| Nepal     | NP -           | EdgeCast   - cs627.wac.epsiloncdn.net. | 77.331402ms |
| Hong Kong | HK -           | EdgeCast   - cs627.wac.epsiloncdn.net. | 99.823144ms |
|-----------|----------------|----------------------------------------|-------------|
INFO[0000] Took 259.059486ms to query DNS servers...

>>>  1s elasped...
➜ $?=0 shadyabhi/cdnfinder [ 3:45PM] (master↑1|…) %
```
