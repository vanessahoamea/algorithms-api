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

k6:
	cd k6 \
		&& npm run build \
		&& K6_WEB_DASHBOARD=true \
		K6_WEB_DASHBOARD_EXPORT=reports/dashboard.html \
		k6 run \
		-e BASE_URL=${APP_BASE_URL} \
		-e OPTIONS_FILE=${K6_OPTIONS_FILE} \
		dist/script.js

compose-up:
	docker compose up --build --exit-code-from test

compose-down:
	docker compose down

.PHONY: k6