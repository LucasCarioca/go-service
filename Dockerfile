FROM node:14 as builder-ui

WORKDIR /ui

COPY ./status/package*.json ./

RUN npm ci

COPY ./status/src ./src
COPY ./status/public ./public

RUN npm run build

FROM golang:latest

WORKDIR /app

COPY ./ ./
COPY --from=builder-ui ./ui/build ./public

RUN go build -o main .

CMD ["/app/main"]