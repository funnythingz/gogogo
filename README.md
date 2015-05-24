# GoGoGo

This is web application in the Golang project template of DDD base.

## Usage

```
% git clone git@github.com:funnythingz/gogogo.git
% cd gogogo/
```

### database setting

Must database setting in MySQL before try `GoGoGo`

```
% cp database.toml.sample database.toml
```

create database and user/password.

### run apprication

```
% go get github.com/mattn/gom
% gom install
% gom migration/migrate.go
% gom exec gin
```

## Command

### LiveReload server

```
% gom exec gin
```

access to `http://localhost:3000`

### migration

```
% gom run migrate/migration.go
```

### test

```
% gom test -v
```

&copy; funnythingz
