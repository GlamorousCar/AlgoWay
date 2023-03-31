# Use an official Golang runtime as a parent image
FROM golang:1.16-alpine as gobuild
# Set the working directory to the app directory
WORKDIR /AlgoWay/backend/cmd/web
#RUN mkdir .postgresql && wget "https://storage.yandexcloud.net/cloud-certs/CA.pem" --output-document .postgresql/root.crt && chmod 0600 .postgresql/root.crt
# Copy the source code into the container
#COPY backend/.env .
COPY . /AlgoWay
#COPY go.* ./
COPY backend/go.mod .
COPY backend/go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
#RUN #go mod download
# Build the Go app
RUN go build -o app .


# Set the entry point of the container to the web binary


FROM alpine:latest

ARG DBUSER
ARG DBNAME
ARG HOST
ARG DBPORT
ARG DBPASS

WORKDIR /app
COPY --from=gobuild /AlgoWay/backend/cmd/web/app .


RUN touch .env
RUN echo dbuser=$DBUSER > .env
RUN echo dbname=$DBNAME >> .env
RUN echo host=$HOST >> .env
RUN echo dbpass=$DBPASS >> .env
RUN echo dbport=$DBPORT >> .env
RUN echo ca=.postgresql/root.crt >> .env


RUN cat .env


#COPY backend/.env .
RUN mkdir .postgresql && wget "https://storage.yandexcloud.net/cloud-certs/CA.pem" --output-document .postgresql/root.crt && chmod 0600 .postgresql/root.crt

#RUN #chmod a+x /app

RUN chmod a+x /app
ENTRYPOINT [ "./app" ]
#EXPOSE 4000

#FROM golang:1.16 as gobuild
#WORKDIR /app
#COPY . .
#RUN go build -o app


