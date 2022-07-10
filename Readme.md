# To improve
- avoid global variable
- cmd folder (take example for goa design)
- Masterize migration
- Remove or move useless file from their respective folders
- server.go/cron.go not in the graphql folder
- support websocket
- unit test flow
- push to github clean
- dataloader
- authentification

# Structure
```
├── cmd
    ├── server.go
    ├── cron.go
    ├── main.go
├── config
    ├── config.go
├── sql
    |── migration
    |── schema
    |── query
    |── internal
    |── sqlc.yaml
├── graphql
    ├─── resolver.go
    ├─── generated.go
    ├─── model.go
    ├─── mypkg
├── service
├── utils

Dockerfile
Makefile
go.mod
go.sum
.gitignore
gqlgen.yml
schema.graphql
Readme.md
```