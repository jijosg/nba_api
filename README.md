# ![](favicon.png) nba_api
Create a nba api in golang

This is the first draft and got the api to work finally :smile:

## TODO :notebook_with_decorative_cover:
- Release binary
- Add more endpoints
- Refine the existing api calls
- Better cli interface and help commands

## Build locally :arrow_forward:

```
git clone https://github.com/jijosg/nba_api.git && cd nba_api
go build .
sudo mv nba_api /usr/local/bin/
```
## Check latest scores :rocket:

```bash
> nba_api score 
No. Home Team     Away Team     Game Status Scores   
1   Celtics       Wizards       Final       144 - 102
2   Bucks         Mavericks     Final       112 - 118
3   Lakers        Nuggets       Final       118 - 129
4   Pacers        Pistons       Final       117 - 121
5   Cavaliers     76ers         Final       108 - 112
6   Magic         Knicks        Final       88  - 118
7   Raptors       Heat          Final       109 - 114
8   Rockets       Timberwolves  Final       132 - 139
9   Thunder       Suns          Final       117 -  96
10  Spurs         Trail Blazers Final       113 -  92
11  Kings         Warriors      Final       90  - 109
12  Clippers      Pelicans      Final       119 - 100
```
## List the NBA teams :basketball: :notebook:

```bash
> go run main.go list 
list called
TEAM_ID     ABBREVIATION  NICKNAME      YEAR_FOUNDED  CITY           FULLNAME                STATE
1610612737  ATL           Hawks         1946          Atlanta        Atlanta Hawks           Atlanta
1610612738  BOS           Celtics       1946          Boston         Boston Celtics          Massachusetts
1610612751  BKN           Nets          1976          Brooklyn       Brooklyn Nets           New York
1610612766  CHA           Hornets       1988          Charlotte      Charlotte Hornets       North Carolina
1610612741  CHI           Bulls         1966          Chicago        Chicago Bulls           Illinois
1610612739  CLE           Cavaliers     1970          Cleveland      Cleveland Cavaliers     Ohio
1610612742  DAL           Mavericks     1980          Dallas         Dallas Mavericks        Texas
1610612743  DEN           Nuggets       1976          Denver         Denver Nuggets          Colorado
1610612765  DET           Pistons       1948          Detroit        Detroit Pistons         Michigan
1610612744  GSW           Warriors      1946          Golden State   Golden State Warriors   California
1610612745  HOU           Rockets       1967          Houston        Houston Rockets         Texas
1610612754  IND           Pacers        1976          Indiana        Indiana Pacers          Indiana
1610612746  LAC           Clippers      1970          Los Angeles    Los Angeles Clippers    California
1610612747  LAL           Lakers        1948          Los Angeles    Los Angeles Lakers      California
1610612763  MEM           Grizzlies     1995          Memphis        Memphis Grizzlies       Tennessee
1610612748  MIA           Heat          1988          Miami          Miami Heat              Florida
1610612749  MIL           Bucks         1968          Milwaukee      Milwaukee Bucks         Wisconsin
1610612750  MIN           Timberwolves  1989          Minnesota      Minnesota Timberwolves  Minnesota
1610612740  NOP           Pelicans      2002          New Orleans    New Orleans Pelicans    Louisiana
1610612752  NYK           Knicks        1946          New York       New York Knicks         New York
1610612753  ORL           Magic         1989          Orlando        Orlando Magic           Florida
1610612755  PHI           76ers         1949          Philadelphia   Philadelphia 76ers      Pennsylvania
1610612756  PHX           Suns          1968          Phoenix        Phoenix Suns            Arizona
1610612757  POR           Trail Blazers 1970          Portland       Portland Trail Blazers  Oregon
1610612758  SAC           Kings         1948          Sacramento     Sacramento Kings        California
1610612759  SAS           Spurs         1976          San Antonio    San Antonio Spurs       Texas
1610612760  OKC           Thunder       1967          Oklahoma City  Oklahoma City Thunder   Oklahoma
1610612761  TOR           Raptors       1995          Toronto        Toronto Raptors         Ontario
1610612762  UTA           Jazz          1974          Utah           Utah Jazz               Utah
1610612764  WAS           Wizards       1961          Washington     Washington Wizards      District of Columbia
```

## Querying stats.nba.com
Finding league games for celtics
```bash
curl -XGET -m 30 https://stats.nba.com/stats/leaguegamefinder/\?PlayerOrTeam\='T'\&\&TeamID\=1610612738 \
-H "Accept: application/json, text/plain, */*" \
-H "Accept-Encoding: gzip, deflate, br" \
-H "x-nba-stats-origin: stats" \
-H "x-nba-stats-token: true" \
-H "Host: stats.nba.com" \
-H "User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:72.0) Gecko/20100101 Firefox/72.0" \
-H "Connection: keep-alive" \
-H "Referer: https://stats.nba.com/" \
-H "Pragma: no-cache" \
-H "Cache-Control: no-cache" \
-H "Accept-Language: en-US,en;q=0.5" \
--output nba_data.gz

gunzip nba_data.gz
```

## References
https://github.com/swar/nba_api