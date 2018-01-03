# GeniusLurker

## Install and run

```
go get -u github.com/yappo/geniuslurker
go build go/src/github.com/yappo/geniuslurker/geniuslurker_server/main.go
./main
```

## How to use

Search Genius for lyrics
```
curl 'http://localhost:3000/search?q=great%20day%20mf%20doom'
```

Now you can pick any result from the previous request and
use its url to get lyrics from Genius
```
curl 'http://localhost:3000/lyrics?url=https://genius.com/Madvillain-great-day-lyrics'
```
