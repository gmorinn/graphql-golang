# To improve
- Masterize migration
- support websocket
- unit test flow
- push to github clean
- dataloader
- file
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
├── graphql
    ├─── resolver.go
    ├─── generated.go
    ├─── schem.resolvers.go
    ├─── model
    │   ├─── models_gen.go
    ├─── mypkg
├── service
├── utils

Dockerfile
Makefile
go.mod
go.sum
.gitignore
gqlgen.yml
sqlc.yaml
schema.graphql
Readme.md
```