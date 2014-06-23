
deps: 
	@( git submodule init && git submodule update )

run:
	@( GOPATH=`pwd` go run main.go )

build:
	@( GOPATH=`pwd` go build main.go )

build_linux:
	@( GOPATH=`pwd` GOOS=linux GOARCH=amd64 go build main.go )

deploy: 
	@( bundle exec cap -q production deploy)
