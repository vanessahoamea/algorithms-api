# Variable declarations
regexp ?= Test.*

# Scripts
build:
	cd src && go build -o ../bin/app

run:
	bin/app

test:
	cd src/solvers && go test -run $(regexp)