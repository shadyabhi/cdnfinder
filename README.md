# Summary

* This program gets the list of public nameservers from https://public-dns.info/.
* Picks one nameserver from every country. 
* Do a CNAME query for the specified domain against these picked nameservers.

## Things to note

* Default timeout is 3 seconds

# Sample Output

This output is when timeout for DNS requests is set to 0.15ms (to keep the output short). 
Set it to higher amount to get more results. Default value of timeout is 3 seconds.

```
➜ $?=0 shadyabhi/cdnfinder [ 3:55PM] (master|✚1…) % go build && ./cdnfinder -domain=static.licdn.com -timeout=0.15s
|------------------------------------------------------|----------------------|-------------------------------------------------|-----------------|--------------|
|                       COUNTRY                        |         CITY         |                  CDN/HOSTNAME                   |       IP        |     TIME     |
|------------------------------------------------------|----------------------|-------------------------------------------------|-----------------|--------------|
| Singapore                                            | SG - Singapore       | EdgeCast   - cs627.wac.epsiloncdn.net.          | 202.136.162.11  | 53.453296ms  |
| Nepal                                                | NP -                 | Akamai     - e9706.dscg.akamaiedge.net.         | 115.187.17.50   | 71.278873ms  |
| Hong Kong                                            | HK -                 | EdgeCast   - cs627.wac.epsiloncdn.net.          | 202.14.67.14    | 82.608864ms  |
| Bhutan                                               | BT -                 | Akamai     - e9706.dscg.akamaiedge.net.         | 103.78.117.5    | 95.372476ms  |
| Macao                                                | MO -                 | Akamai     - e9706.dscg.akamaiedge.net.         | 202.175.89.60   | 104.383627ms |
| Myanmar                                              | MM -                 | Akamai     - e9706.dscg.akamaiedge.net.         | 203.81.161.9    | 108.876443ms |
| Sri Lanka                                            | LK - Raddolugama     | Fastly     - dualstack.linkedin.map.fastly.net. | 124.43.25.209   | 114.318868ms |
| Philippines                                          | PH -                 | Akamai     - e9706.dscg.akamaiedge.net.         | 203.115.130.72  | 117.897499ms |
| Indonesia                                            | ID -                 | Akamai     - e9706.dscg.akamaiedge.net.         | 121.52.87.128   | 120.740483ms |
| Taiwan, Province of China                            | TW -                 | Akamai     - e9706.dscg.akamaiedge.net.         | 168.95.192.1    | 119.106063ms |
| Japan                                                | JP - Tokyo           | EdgeCast   - cs627.wac.epsiloncdn.net.          | 203.112.2.4     | 127.342667ms |
| Kenya                                                | KE -                 | Akamai     - e9706.dscg.akamaiedge.net.         | 197.248.2.44    | 130.061689ms |
| Viet Nam                                             | VN - Hanoi           | EdgeCast   - cs627.wac.epsiloncdn.net.          | 203.119.36.106  | 127.424526ms |
| Switzerland                                          | CH -                 | Fastly     - dualstack.linkedin.map.fastly.net. | 195.186.1.111   | 139.938856ms |
| United Kingdom of Great Britain and Northern Ireland | GB - St Albans       | EdgeCast   - cs627.wac.epsiloncdn.net.          | 212.118.241.1   | 139.429856ms |
| Ireland                                              | IE -                 | EdgeCast   - cs627.wac.epsiloncdn.net.          | 62.40.32.33     | 141.841277ms |
| Germany                                              | DE -                 | EdgeCast   - cs627.wac.epsiloncdn.net.          | 194.150.168.168 | 143.024783ms |
| Spain                                                | ES - Barcelona       | EdgeCast   - cs627.wac.epsiloncdn.net.          | 212.36.64.16    | 147.139443ms |
| Jersey                                               | JE - Saint Helier    | EdgeCast   - cs627.wac.epsiloncdn.net.          | 212.9.28.233    | 144.226394ms |
| Yemen                                                | YE -                 | Akamai     - e9706.dscg.akamaiedge.net.         | 195.94.2.226    | 147.223571ms |
| Mongolia                                             | MN - Sühbaatar       | Akamai     - e9706.dscg.akamaiedge.net.         | 183.81.169.102  | 149.676441ms |
| Finland                                              | FI -                 | Akamai     - e9706.dscg.akamaiedge.net.         | 178.33.144.62   | 147.815821ms |
| Hungary                                              | HU - Ozora           | EdgeCast   - cs627.wac.epsiloncdn.net.          | 212.108.197.216 | 151.143186ms |
|------------------------------------------------------|----------------------|-------------------------------------------------|-----------------|--------------|
INFO[0000] Took 314.098237ms to query DNS servers...

>>>  1s elasped...
➜ $?=0 shadyabhi/cdnfinder [ 3:55PM] (master|✚1…) %
```
