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
➜ $?=0 shadyabhi/cdnfinder [ 4:04PM] (master|✚2…) % go build && ./cdnfinder -domain=static.licdn.com -timeout=0.15s -number-ns=3
|---------------------------|-----------------------|-----------------------------------------|-----------------|--------------|
|          COUNTRY          |         CITY          |              CDN/HOSTNAME               |       IP        |     TIME     |
|---------------------------|-----------------------|-----------------------------------------|-----------------|--------------|
| India                     | IN - Mumbai           | EdgeCast   - cs627.wac.epsiloncdn.net.  | 203.115.81.38   | 60.184988ms  |
| Singapore                 | SG - Singapore        | EdgeCast   - cs627.wac.epsiloncdn.net.  | 202.136.162.11  | 76.090818ms  |
| India                     | IN -                  | EdgeCast   - cs627.wac.epsiloncdn.net.  | 202.148.202.4   | 42.188365ms  |
| Hong Kong                 | HK -                  | EdgeCast   - cs627.wac.epsiloncdn.net.  | 202.14.67.14    | 120.966881ms |
| Thailand                  | TH -                  | Akamai     - e9706.dscg.akamaiedge.net. | 202.28.162.1    | 124.96468ms  |
| Myanmar                   | MM -                  | Akamai     - e9706.dscg.akamaiedge.net. | 203.81.161.9    | 130.96638ms  |
| United States of America  | US - Mountain View    | Akamai     - e9706.dscg.akamaiedge.net. | 8.8.8.8         | 140.998307ms |
| Viet Nam                  | VN - Hanoi            | EdgeCast   - cs627.wac.epsiloncdn.net.  | 203.119.36.106  | 142.207267ms |
| Philippines               | PH -                  | Akamai     - e9706.dscg.akamaiedge.net. | 203.115.130.72  | 140.294227ms |
| India                     | IN - Mumbai           | EdgeCast   - cs627.wac.epsiloncdn.net.  | 203.115.71.66   | 41.051515ms  |
| Macao                     | MO -                  | Akamai     - e9706.dscg.akamaiedge.net. | 202.175.89.60   | 153.657013ms |
| Hong Kong                 | HK -                  | EdgeCast   - cs627.wac.epsiloncdn.net.  | 202.14.67.4     | 99.528695ms  |
| Viet Nam                  | VN - Hanoi            | Akamai     - e9706.dscg.akamaiedge.net. | 203.119.8.106   | 121.120333ms |
| Hong Kong                 | HK - Central District | EdgeCast   - cs627.wac.epsiloncdn.net.  | 202.130.97.66   | 111.606855ms |
| Sri Lanka                 | LK -                  | Akamai     - e9706.dscg.akamaiedge.net. | 220.247.213.129 | 97.972967ms  |
| Taiwan, Province of China | TW - Taipei           | Akamai     - e9706.dscg.akamaiedge.net. | 139.175.252.16  | 140.030192ms |
|---------------------------|-----------------------|-----------------------------------------|-----------------|--------------|
INFO[0000] Took 610.451149ms to query DNS servers...

>>>  2s elasped...
➜ $?=0 shadyabhi/cdnfinder [ 4:04PM] (master|✚2…) %
```
