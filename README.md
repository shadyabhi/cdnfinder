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
➜ $?=0 shadyabhi/cdnfinder [ 3:52PM] (master|✚2…) % go build && ./cdnfinder -domain=static.licdn.com -timeout=0.15s
|------------------------------------------------------|----------------------|-----------------------------------------|--------------|----------------|
|                       COUNTRY                        |         CITY         |              CDN/HOSTNAME               |      IP      |      TIME      |
|------------------------------------------------------|----------------------|-----------------------------------------|--------------|----------------|
| United States of America                             | US - Mountain View   | Akamai     - e9706.dscg.akamaiedge.net. | 36.423185ms  | 8.8.8.8        |
| Nepal                                                | NP -                 | Akamai     - e9706.dscg.akamaiedge.net. | 65.7544ms    | 115.187.17.50  |
| Singapore                                            | SG - Singapore       | EdgeCast   - cs627.wac.epsiloncdn.net.  | 62.629647ms  | 202.136.162.11 |
| Bhutan                                               | BT -                 | Akamai     - e9706.dscg.akamaiedge.net. | 96.372663ms  | 103.78.117.5   |
| Hong Kong                                            | HK -                 | EdgeCast   - cs627.wac.epsiloncdn.net.  | 99.706276ms  | 202.14.67.14   |
| Myanmar                                              | MM -                 | Akamai     - e9706.dscg.akamaiedge.net. | 106.64258ms  | 203.81.161.9   |
| Philippines                                          | PH -                 | Akamai     - e9706.dscg.akamaiedge.net. | 106.837845ms | 203.115.130.72 |
| Kenya                                                | KE -                 | Akamai     - e9706.dscg.akamaiedge.net. | 110.129266ms | 197.248.2.44   |
| Sri Lanka                                            | LK - Raddolugama     | EdgeCast   - cs627.wac.epsiloncdn.net.  | 109.15236ms  | 124.43.25.209  |
| Viet Nam                                             | VN - Hanoi           | EdgeCast   - cs627.wac.epsiloncdn.net.  | 116.046334ms | 203.119.36.106 |
| Indonesia                                            | ID -                 | EdgeCast   - cs627.wac.epsiloncdn.net.  | 116.6236ms   | 121.52.87.128  |
| Japan                                                | JP - Tokyo           | EdgeCast   - cs627.wac.epsiloncdn.net.  | 128.879125ms | 203.112.2.4    |
| United Kingdom of Great Britain and Northern Ireland | GB - St Albans       | EdgeCast   - cs627.wac.epsiloncdn.net.  | 147.154655ms | 212.118.241.1  |
| Luxembourg                                           | LU - Luxembourg      | EdgeCast   - cs627.wac.epsiloncdn.net.  | 149.565665ms | 195.46.239.58  |
| Spain                                                | ES - Barcelona       | EdgeCast   - cs627.wac.epsiloncdn.net.  | 148.801557ms | 212.36.64.16   |
| Switzerland                                          | CH -                 | EdgeCast   - cs627.wac.epsiloncdn.net.  | 147.153861ms | 195.186.1.111  |
| Ireland                                              | IE -                 | EdgeCast   - cs627.wac.epsiloncdn.net.  | 151.051385ms | 62.40.32.33    |
| Finland                                              | FI -                 | Akamai     - e9706.dscg.akamaiedge.net. | 151.15823ms  | 178.33.144.62  |
|------------------------------------------------------|----------------------|-----------------------------------------|--------------|----------------|
INFO[0000] Took 310.20445ms to query DNS servers...

>>>  1s elasped...
➜ $?=0 shadyabhi/cdnfinder [ 3:52PM] (master|✚2…) %
```
