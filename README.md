A log viewer tool for different projects using golang.


## Prerequisites
Run the following commands in your command line
```
// Package Manager
go get -u github.com/kardianos/govendor
// live reloading tool
go get github.com/oxequa/realizego
// DB migration tool
get -u github.com/pressly/goose/cmd/goose
```

### Govendor
```
Fetch or Install:
	run `govendor fetch <packagename>`
Pull packages into vendor folder
	run `govendor sync`
```

## Development
* Clone this repo
* Run `govendor sync`
* Add an environment file for config
	* cp `.env.example` `.env`  ( in the root dir of the project)
	* Fill in appropriate DB connection
* Install Go Live Reloading `go get github.com/oxequa/realize`
* Run `realize start`

### Frontend
* `npm install -g @vue/cli` || `yarn global add @vue/cli`
* `cd vewlog/views`
* Refer to `vewlog/views/Readme.md` for more info

** Documentation in progress **

## Production & Deployment
** Documentation in progress **

## Sources
* [Gin-Gonic Golang Web Framework](https://github.com/gin-gonic/gin)
* [Go HTTP/NET](https://golang.org/pkg/net/http/)
* [Realize - Go live reloading](https://github.com/oxequa/realize)
* [Go Validator](https://github.com/thedevsaddam/govalidator)
