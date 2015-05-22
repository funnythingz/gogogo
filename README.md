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

auto migrate

```
% gom run migrate/migration.go
```

reset

```
% gom run migrate/migration.go reset
```

### test

```
% gom test -v
```

## use package

- Go Manager - bundle for go ... https://github.com/mattn/gom
- Web framework ... https://github.com/zenazn/goji
- Params binding ... https://github.com/goji/param
- Asset Management ... https://github.com/shaoshing/train
- ORMapper ... https://github.com/jinzhu/gorm
- Markdown to HTML ... https://github.com/russross/blackfriday
- Sanitizer ... https://github.com/microcosm-cc/bluemonday
- Truncate ... https://github.com/funnythingz/sunnyday
- Live reload ... https://github.com/codegangsta/gin
- Template ... https://github.com/yosssi/ace
- Validater ... https://github.com/asaskevich/govalidator
- PrintDebug ... https://github.com/k0kubun/pp
- Redis ... https://github.com/garyburd/redigo/redis
- TOML ... https://github.com/BurntSushi/toml

and more...

&copy; funnythingz
