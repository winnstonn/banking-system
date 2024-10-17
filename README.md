# banking-system
Repository for all codes related to simple Banking System app in Go

### Requirements

- Install [Go](https://go.dev/doc/install) --> current used version 1.22.6
- Install [Redis](https://redis.io/docs/latest/operate/oss_and_stack/install/install-redis/install-redis-on-mac-os/)
- Install [PostgreSQL](https://www.postgresql.org/download/macosx/)

#### Install dependencies
Execute command:
```sh
make all
```

### Run existing project
- Copy `.env.template` file & rename it to `.env`
- Execute command:
```sh
make run
```