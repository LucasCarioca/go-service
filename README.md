```
go get -u github.com/gin-gonic/gin
brew install watchexec
make start
```

```
docker build . -t go-service
docker run -it -p 8080:8080 go-service
```