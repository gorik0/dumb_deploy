FROM golang:1.22-alpine AS build

# create a working directory inside the image
WORKDIR /app

# copy Go modules and dependencies to image
COPY go.mod ./

# download Go modules and dependencies
RUN go mod download

# copy directory files i.e all files ending with .go
COPY *.go ./

# compile application
RUN go build -o /godocker

##
## STEP 2 - DEPLOY
##
FROM scratch

WORKDIR /

COPY --from=build /godocker /godocker



ENTRYPOINT ["echo"]