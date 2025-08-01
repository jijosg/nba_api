# nba_api [![.github/workflows/release.yml](https://github.com/jijosg/nba_api/actions/workflows/release.yml/badge.svg)](https://github.com/jijosg/nba_api/actions/workflows/release.yml) [![Go](https://github.com/jijosg/nba_api/actions/workflows/go.yml/badge.svg)](https://github.com/jijosg/nba_api/actions/workflows/go.yml)

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

## Get player stats
```bash
> nba_api players -p 2544
PLAYER_ID     PLAYER_NAME   TimeFrame     PTS           AST           REB           PIE           
2544          LeBron James  2021-22       30.3          6.2           8.2           0.18          

> nba_api players -p 201939
PLAYER_ID     PLAYER_NAME   TimeFrame     PTS           AST           REB           PIE           
201939        Stephen Curry 2021-22       25.5          6.3           5.2           0.155  
```

## CLI Commands and Flags

### Team Commands

#### Show Team Info (Root Command)

Show information about a specific NBA team (default: Lakers).

```bash
nba_api [flags]
```

**Flags:**
- `--team string` &nbsp;&nbsp;Specify a team abbreviation to get info (default: `LAL`)
- `--version` &nbsp;&nbsp;Show version information

**Description:**  
If run without arguments, shows help and available commands.  
If run with `--team`, displays information about the specified team.

---

#### List All NBA Teams

List all NBA teams that exist today.

```bash
nba_api list
```

**Description:**  
Displays a table of all NBA teams, including team ID, abbreviation, nickname, year founded, city, full name, and state.

---

### Player Commands

#### Players

Get player info and stats by player ID or search by name.

```bash
nba_api players [flags]
```

**Flags:**
- `-p`, `--playerId string` &nbsp;&nbsp;Player ID to get info (default: `2544`)
- `-n`, `--name string` &nbsp;&nbsp;Player name to search for (case-insensitive substring match)
- `-r`, `--recent` &nbsp;&nbsp;Show recent games for the player (shows 5 most recent games)

**Examples:**
```bash
nba_api players -p 2544
nba_api players -n curry
nba_api players -p 2544 --recent
```

---

### Score Commands

#### View Today's Scores

Show the scores for today's NBA games.

```bash
nba_api score
```

**Description:**  
Displays a table of today's NBA games, including home team, away team, game status, and scores.  
If there are no games today, it will display a message indicating so.

---
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
