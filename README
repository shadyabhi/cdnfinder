# Summary

* This program gets the list of public nameservers from https://public-dns.info/.
* Picks one nameserver from every country. 
* Do a CNAME query for the specified domain against these picked nameservers.

## Things to note

* Default timeout is 3 seconds

# Sample Output

```
➜ $?=1 shadyabhi/cdnfinder [ 3:33PM] (master|…) % ./cdn-finder -domain=static.licdn.com -timeout=1s
WARN[0000] No results for domain: static.licdn.com.
|------------------------------------------------------|------------------------------|-------------------------------------------------|
|                       COUNTRY                        |             CITY             |                  CDN/HOSTNAME                   |
|------------------------------------------------------|------------------------------|-------------------------------------------------|
| Singapore                                            | SG - Singapore               | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Hong Kong                                            | HK -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Viet Nam                                             | VN - Hanoi                   | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Japan                                                | JP - Tokyo                   | Fastly     - dualstack.linkedin.map.fastly.net. |
| Taiwan, Province of China                            | TW -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| Myanmar                                              | MM -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| United States of America                             | US - Mountain View           | Akamai     - e9706.dscg.akamaiedge.net.         |
| France                                               | FR - Montevrain              | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| United Kingdom of Great Britain and Northern Ireland | GB - St Albans               | Akamai     - e9706.dscg.akamaiedge.net.         |
| Luxembourg                                           | LU - Luxembourg              | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Slovenia                                             | SI -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| Switzerland                                          | CH -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| Germany                                              | DE -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Italy                                                | IT - Arezzo                  | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Belgium                                              | BE -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Slovakia                                             | SK -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Finland                                              | FI -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| Spain                                                | ES - Barcelona               | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Netherlands                                          | NL - Eindhoven               | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Czech Republic                                       | CZ -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| Syrian Arab Republic                                 | SY -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Ireland                                              | IE -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Lithuania                                            | LT -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Norway                                               | NO - Stavanger               | Akamai     - e9706.dscg.akamaiedge.net.         |
| Moldova (Republic of)                                | MD - Chisinau                | Akamai     - e9706.dscg.akamaiedge.net.         |
| Poland                                               | PL - Warsaw                  | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Romania                                              | RO - Agigea                  | Fastly     - dualstack.linkedin.map.fastly.net. |
| Lebanon                                              | LB -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Montenegro                                           | ME -                         | Fastly     - dualstack.linkedin.map.fastly.net. |
| Belarus                                              | BY - Minsk                   | Akamai     - e9706.dscg.akamaiedge.net.         |
| Jersey                                               | JE - Saint Helier            | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Isle of Man                                          | IM - Laxey                   | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Bulgaria                                             | BG -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| Bosnia and Herzegovina                               | BA - Zivinice                | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Latvia                                               | LV - Riga                    | Akamai     - e9706.dscg.akamaiedge.net.         |
| Sri Lanka                                            | LK - Raddolugama             | Fastly     - dualstack.linkedin.map.fastly.net. |
| Turkey                                               | TR - Istanbul                | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Serbia                                               | RS -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Gambia                                               | GM - Kanifing                | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Åland Islands                                        | AX - Saltvik                 | Fastly     - dualstack.linkedin.map.fastly.net. |
| Portugal                                             | PT - Matosinhos Municipality | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Mauritania                                           | MR -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| Ghana                                                | GH -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Iraq                                                 | IQ -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| New Zealand                                          | NZ -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| Micronesia (Federated States of)                     | FM -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| Georgia                                              | GE -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| Andorra                                              | AD - Andorra la Vella        | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Canada                                               | CA - Cambridge               | Fastly     - dualstack.linkedin.map.fastly.net. |
| Guernsey                                             | GG - St Peter Port           | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Gabon                                                | GA - Libreville              | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Kazakhstan                                           | KZ -                         | Fastly     - dualstack.linkedin.map.fastly.net. |
| Turkmenistan                                         | TM -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Afghanistan                                          | AF - Kabul                   | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Liechtenstein                                        | LI - Eschen                  | Akamai     - e9706.dscg.akamaiedge.net.         |
| Ethiopia                                             | ET -                         | Fastly     - dualstack.linkedin.map.fastly.net. |
| Côte d'Ivoire                                        | CI -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| United Arab Emirates                                 | AE - Dubai                   | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Albania                                              | AL -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Tunisia                                              | TN -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| China                                                | CN - Nanjing                 | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Burkina Faso                                         | BF -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| Indonesia                                            | ID -                         |                                              -  |
| Macedonia (the former Yugoslav Republic of)          | MK -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Nigeria                                              | NG -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Macao                                                | MO -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| Dominican Republic                                   | DO - Santo Domingo           | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Bahamas                                              | BS - Nassau                  | Akamai     - e9706.dscg.akamaiedge.net.         |
| Estonia                                              | EE - Tallinn                 | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Australia                                            | AU - Melbourne               | Akamai     - e9706.dscg.akamaiedge.net.         |
| Timor-Leste                                          | TL -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| Israel                                               | IL -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| Anguilla                                             | AI - The Valley              | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Guatemala                                            | GT -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Iceland                                              | IS - Akureyri                | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Nepal                                                | NP -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Cameroon                                             | CM - Douala                  | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Morocco                                              | MA - Tangier                 | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Bonaire, Sint Eustatius and Saba                     | BQ -                         |                                              -  |
| Libya                                                | LY -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Papua New Guinea                                     | PG - Port Moresby            | Akamai     - e9706.dscg.akamaiedge.net.         |
| Kuwait                                               | KW - Kuwait City             | Akamai     - e9706.dscg.akamaiedge.net.         |
| Barbados                                             | BB - Hastings                | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Grenada                                              | GD - St. George's            | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Jordan                                               | JO -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| Brunei Darussalam                                    | BN - Bandar Seri Begawan     | Akamai     - e9706.dscg.akamaiedge.net.         |
| Malaysia                                             | MY - Damansara               | Akamai     - e9706.dscg.akamaiedge.net.         |
| Ukraine                                              | UA -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Namibia                                              | NA - Windhoek                | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Mauritius                                            | MU -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Martinique                                           | MQ - Le Gros-Morne           | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| French Polynesia                                     | PF - Punaauia                | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Egypt                                                | EG -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Jamaica                                              | JM - Kingston                | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Suriname                                             | SR - Paramaribo              | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Cayman Islands                                       | KY - George Town             | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Uganda                                               | UG -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Philippines                                          | PH -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| Azerbaijan                                           | AZ - Baku                    | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Virgin Islands (British)                             | VG - Road Town               | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Sweden                                               | SE - Lund                    | Akamai     - e9706.dscg.akamaiedge.net.         |
| Mozambique                                           | MZ -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Puerto Rico                                          | PR - Sabana Hoyos            | Akamai     - e9706.dscg.akamaiedge.net.         |
| Réunion                                              | RE - Saint-Paul              | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Brazil                                               | BR -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| Saint Pierre and Miquelon                            | PM - Miquelon                | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Tajikistan                                           | TJ - Dushanbe                | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Korea (Republic of)                                  | KR - Seongnam-si             | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Djibouti                                             | DJ -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| South Africa                                         | ZA - Waverley                | Fastly     - dualstack.linkedin.map.fastly.net. |
| Kyrgyzstan                                           | KG -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| Saint Kitts and Nevis                                | KN - Basseterre              | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Haiti                                                | HT -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Saint Barthélemy                                     | BL -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| El Salvador                                          | SV - San Salvador            | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Swaziland                                            | SZ -                         | Fastly     - dualstack.linkedin.map.fastly.net. |
| Sint Maarten (Dutch part)                            | SX - Philipsburg             | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Russian Federation                                   | RU -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Saint Vincent and the Grenadines                     | VC - Kingstown               | Akamai     - e9706.dscg.akamaiedge.net.         |
| Uruguay                                              | UY - Montevideo              | Akamai     - e9706.dscg.akamaiedge.net.         |
| Liberia                                              | LR -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Costa Rica                                           | CR - San José                | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Curaçao                                              | CW - Willemstad              | Akamai     - e9706.dscg.akamaiedge.net.         |
| Antigua and Barbuda                                  | AG - St. John's              | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Denmark                                              | DK -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| Virgin Islands (U.S.)                                | VI -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Paraguay                                             | PY -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Algeria                                              | DZ -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| Panama                                               | PA -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Bahrain                                              | BH -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Venezuela (Bolivarian Republic of)                   | VE - Caracas                 | Akamai     - e9706.dscg.akamaiedge.net.         |
| Fiji                                                 | FJ - Suva                    | Akamai     - e9706.dscg.akamaiedge.net.         |
| Ecuador                                              | EC - Quito                   | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Sao Tome and Principe                                | ST -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| Mongolia                                             | MN - Sühbaatar               | Akamai     - e9706.dscg.akamaiedge.net.         |
| Colombia                                             | CO - Cundinamarca            | Akamai     - e9706.dscg.akamaiedge.net.         |
| Guam                                                 | GU - Hagåtña                 | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Qatar                                                | QA - Doha                    | Akamai     - e9706.dscg.akamaiedge.net.         |
| Angola                                               | AO -                         | Fastly     - dualstack.linkedin.map.fastly.net. |
| Trinidad and Tobago                                  | TT - Tunapuna                | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Armenia                                              | AM -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| Argentina                                            | AR -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Zimbabwe                                             | ZW -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| Honduras                                             | HN -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Zambia                                               | ZM -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| Nauru                                                | NR -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| Tonga                                                | TO -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Chile                                                | CL - Santiago                | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Bhutan                                               | BT -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| Pakistan                                             | PK -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| Bangladesh                                           | BD -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| Kenya                                                | KE -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| Cabo Verde                                           | CV - Praia                   | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Saint Lucia                                          | LC - Castries                | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Central African Republic                             | CF -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| New Caledonia                                        | NC - Noumea                  | Fastly     - dualstack.linkedin.map.fastly.net. |
| Seychelles                                           | SC -                         | Fastly     - dualstack.linkedin.map.fastly.net. |
| Palestine, State of                                  | PS - Ramallah                | Akamai     - e9706.dscg.akamaiedge.net.         |
| Saudi Arabia                                         | SA -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| American Samoa                                       | AS - Pago Pago               | Fastly     - dualstack.linkedin.map.fastly.net. |
| Bolivia (Plurinational State of)                     | BO - La Paz                  | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Greenland                                            | GL - Nuuk                    | Akamai     - e9706.dscg.akamaiedge.net.         |
| Cook Islands                                         | CK -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Malawi                                               | MW -                         | Fastly     - dualstack.linkedin.map.fastly.net. |
| Tanzania, United Republic of                         | TZ -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Botswana                                             | BW -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
| Burundi                                              | BI -                         | EdgeCast   - cs627.wac.epsiloncdn.net.          |
| Guyana                                               | GY -                         | Akamai     - e9706.dscg.akamaiedge.net.         |
|------------------------------------------------------|------------------------------|-------------------------------------------------|
INFO[0001] Took 1.165006686s to query DNS servers...

>>>  1s elasped...
➜ $?=0 shadyabhi/cdnfinder [ 3:33PM] (master|…) %
```
