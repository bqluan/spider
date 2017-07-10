spider
======
spider

### Build dev env ###
```
docker build -t spider-dev .
```

### Dev shell ###
```
docker run --rm -it -v `pwd`:/go/src/github.com/bqluan/spider spider-dev bash
```

### Build for Windows ###
```
env GOOS=windows GOARCH=386 go build
```

### Build for Linux ###
```
go build
```
