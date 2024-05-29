# import golang compiler
FROM golang:1.22.3-alpine

# use dir /app on image to run app
WORKDIR /app

# copy local project to work directory
COPY . .

# install dumb-init process management
RUN apk add dumb-init

# build project to binary file
RUN go build -o App .

# run main program
ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["./App"]