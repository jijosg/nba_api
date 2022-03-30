# nba_api
Create a nba api in golang

This is the first draft and got the api to work finally 😆
## Querying stats.nba.com
```bash
curl -XGET -m 30 https://stats.nba.com/stats/leaguegamefinder/\?PlayerOrTeam\='T'\&\&TeamID\=1610612738 -H "Accept: application/json, text/plain, */*" -H "Accept-Encoding: gzip, deflate, br" -H "x-nba-stats-origin: stats" -H "x-nba-stats-token: true" -H "Host: stats.nba.com" -H "User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:72.0) Gecko/20100101 Firefox/72.0" -H "Connection: keep-alive" -H "Referer: https://stats.nba.com/" -H "Pragma: no-cache" -H "Cache-Control: no-cache" -H "Accept-Language: en-US,en;q=0.5" --output nba_data.gz

gunzip nba_data.gz
```
