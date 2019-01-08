 HELP_FUN = \
         %help; \
         while(<>) { push @{$$help{$$2 // 'options'}}, [$$1, $$3] if /^(\w+)\s*:.*\#\#(?:@(\w+))?\s(.*)$$/ }; \
         print "usage: make [target]\n\n"; \
     for (keys %help) { \
         print "$$_:\n"; $$sep = " " x (20 - length $$_->[0]); \
         print "  $$_->[0]$$sep$$_->[1]\n" for @{$$help{$$_}}; \
         print "\n"; }     

help:           ##@miscellaneous Show this help.
	@perl -e '$(HELP_FUN)' $(MAKEFILE_LIST)

setup:##@setup Setup first time for your project
	install build migrate
run:##@run run your app
	$(shell ./app)
runapp:##@run run database (postgres & redis) and run your app
	rundb run
build:##@setup build your porject
	go build -o cli command/main.go
	go build -o app main.go
test:##@test test your project
	go test ./... -coverprofile=cover.out
coverage:##@test check your test coverage
	go tool cover -html=cover.out
install:##@setup install depedency of project
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
	dep ensure
init:##@setup init your project, please init before do anything
	$(shell echo "MODE: $(mode)" > ./config.yaml)
migrate:##@run migrate your data
	./cli db --action=migrate
rundb:##@run run database postgresql & redis
	$(shell cd docker && docker-compose up --build)
install_docker:##@setup how to install docker ?
	$(shell echo "plesae read documentation at https://docs.docker.com/install/")