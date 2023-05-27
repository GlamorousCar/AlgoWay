# Use an official Golang runtime as a parent image
FROM golang:1.20-alpine as gobuild
# Set the working directory to the app directory
WORKDIR /app

COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

COPY . ./

# Build the Go app
RUN go build -o app /app/cmd/web

# Set the entry point of the container to the web binary

ARG DBUSER
ARG DBNAME
ARG HOST
ARG DBPORT
ARG DBPASS
ARG SECRET_KEY

ENV PATH="/usr/local/go/bin:${PATH}"

RUN touch .env
RUN echo dbuser=$DBUSER > .env
RUN echo dbname=$DBNAME >> .env
RUN echo host=$HOST >> .env
RUN echo dbpass=$DBPASS >> .env
RUN echo dbport=$DBPORT >> .env
RUN echo secret_key=$SECRET_KEY >> .env
RUN echo ca=.postgresql/root.crt >> .env

RUN mkdir .postgresql && wget "https://storage.yandexcloud.net/cloud-certs/CA.pem" --output-document .postgresql/root.crt && chmod 0600 .postgresql/root.crt

RUN chmod a+x /app
ENTRYPOINT [ "./app" ]
