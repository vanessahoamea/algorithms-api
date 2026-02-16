-include .env

# Variable declarations
regexp ?= Test.*

# Scripts
build:
	cd src && go build -o ../bin/app

run:
	bin/app

test:
	cd src/solvers && go test -run $(regexp)

swagger:
	cd src && swag init

performance:
	cd k6 \
		&& npm run build \
		&& K6_WEB_DASHBOARD=true \
		K6_WEB_DASHBOARD_EXPORT=dashboard.html \
		k6 run \
		-e BASE_URL=${BASE_URL} \
		-e OPTIONS_FILE=${OPTIONS_FILE} \
		dist/script.js