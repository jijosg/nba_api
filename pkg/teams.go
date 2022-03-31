package pkg

type Teams struct {
	Id               int
	Abbreviation     string
	Nickname         string
	YearFounded      int
	City             string
	FullName         string
	State            string
	ChampionshipYear []int
}

var NBATeams []Teams = []Teams{
	Teams{ 1610612737,  "ATL",  "Hawks",  1946,  "Atlanta",  "Atlanta Hawks",  "Atlanta",  []int{1958}},
	Teams{ 1610612738,  "BOS",  "Celtics",  1946,  "Boston",  "Boston Celtics",  "Massachusetts",  []int{1957, 1959, 1960, 1961, 1962, 1963, 1964, 1965, 1966, 1968, 1969, 1974, 1976, 1981, 1984, 1986, 2008}},
	Teams{ 1610612751,  "BKN",  "Nets",  1976,  "Brooklyn",  "Brooklyn Nets",  "New York",  []int{}},
	Teams{ 1610612766,  "CHA",  "Hornets",  1988,  "Charlotte",  "Charlotte Hornets",  "North Carolina",  []int{}},
	Teams{ 1610612741,  "CHI",  "Bulls",  1966,  "Chicago",  "Chicago Bulls",  "Illinois",  []int{1991, 1992, 1993, 1996, 1997, 1998}},
	Teams{ 1610612739,  "CLE",  "Cavaliers",  1970,  "Cleveland",  "Cleveland Cavaliers",  "Ohio",  []int{2016}},
	Teams{ 1610612742,  "DAL",  "Mavericks",  1980,  "Dallas",  "Dallas Mavericks",  "Texas",  []int{2011}},
	Teams{ 1610612743,  "DEN",  "Nuggets",  1976,  "Denver",  "Denver Nuggets",  "Colorado",  []int{}},
	Teams{ 1610612765,  "DET",  "Pistons",  1948,  "Detroit",  "Detroit Pistons",  "Michigan",  []int{1989, 1990, 2004}},
	Teams{ 1610612744,  "GSW",  "Warriors",  1946,  "Golden State",  "Golden State Warriors",  "California",  []int{1947, 1956, 1975, 2015, 2017, 2018}},
	Teams{ 1610612745,  "HOU",  "Rockets",  1967,  "Houston",  "Houston Rockets",  "Texas",  []int{1994, 1995}},
	Teams{ 1610612754,  "IND",  "Pacers",  1976,  "Indiana",  "Indiana Pacers",  "Indiana",  []int{}},
	Teams{ 1610612746,  "LAC",  "Clippers",  1970,  "Los Angeles",  "Los Angeles Clippers",  "California",  []int{}},
	Teams{ 1610612747,  "LAL",  "Lakers",  1948,  "Los Angeles",  "Los Angeles Lakers",  "California",  []int{1949, 1950, 1952, 1953, 1954, 1972, 1980, 1982, 1985, 1987, 1988, 2000, 2001, 2002, 2009, 2010, 2020}},
	Teams{ 1610612763,  "MEM",  "Grizzlies",  1995,  "Memphis",  "Memphis Grizzlies",  "Tennessee",  []int{}},
	Teams{ 1610612748,  "MIA",  "Heat",  1988,  "Miami",  "Miami Heat",  "Florida",  []int{2006, 2012, 2013}},
	Teams{ 1610612749,  "MIL",  "Bucks",  1968,  "Milwaukee",  "Milwaukee Bucks",  "Wisconsin",  []int{1971, 2021}},
	Teams{ 1610612750,  "MIN",  "Timberwolves",  1989,  "Minnesota",  "Minnesota Timberwolves",  "Minnesota",  []int{}},
	Teams{ 1610612740,  "NOP",  "Pelicans",  2002,  "New Orleans",  "New Orleans Pelicans",  "Louisiana",  []int{}},
	Teams{ 1610612752,  "NYK",  "Knicks",  1946,  "New York",  "New York Knicks",  "New York", []int{1970, 1973}},
	Teams{ 1610612753,  "ORL",  "Magic",  1989,  "Orlando",  "Orlando Magic",  "Florida", []int{}},
	Teams{ 1610612755,  "PHI",  "76ers",  1949,  "Philadelphia",  "Philadelphia 76ers",  "Pennsylvania", []int{1955, 1967, 1983}},
	Teams{ 1610612756,  "PHX",  "Suns",  1968,  "Phoenix",  "Phoenix Suns",  "Arizona", []int{}},
	Teams{ 1610612757,  "POR",  "Trail Blazers",  1970,  "Portland",  "Portland Trail Blazers",  "Oregon", []int{1977}},
	Teams{ 1610612758,  "SAC",  "Kings",  1948,  "Sacramento",  "Sacramento Kings",  "California", []int{1951}},
	Teams{ 1610612759,  "SAS",  "Spurs",  1976,  "San Antonio",  "San Antonio Spurs",  "Texas", []int{1999, 2003, 2005, 2007, 2014}},
	Teams{ 1610612760,  "OKC",  "Thunder",  1967,  "Oklahoma City",  "Oklahoma City Thunder",  "Oklahoma", []int{1979}},
	Teams{ 1610612761,  "TOR",  "Raptors",  1995,  "Toronto",  "Toronto Raptors",  "Ontario", []int{2019}},
	Teams{ 1610612762,  "UTA",  "Jazz",  1974,  "Utah",  "Utah Jazz",  "Utah", []int{}},
	Teams{ 1610612764,  "WAS",  "Wizards",  1961,  "Washington",  "Washington Wizards",  "District of Columbia", []int{1978}},
}
