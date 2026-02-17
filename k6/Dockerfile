# syntax=docker/dockerfile:1

# Stage 1: Create compiled JavaScript file
FROM node:25-alpine AS node

WORKDIR /build

COPY package.json package-lock.json ./

RUN npm ci

COPY . ./

RUN npm run build

# Stage 2: Run k6 tests against the API
FROM grafana/k6:latest

WORKDIR /test

COPY --from=node /build/dist ./dist
COPY --from=node /build/configs ./configs

ENTRYPOINT ["k6", "run"]

CMD ["dist/script.js"]