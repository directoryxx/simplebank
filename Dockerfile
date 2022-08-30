FROM golang:1.19.0-alpine

# Install Required Package
RUN apk add curl openssh-client gcc git make

# Change Workdir to /tmp
WORKDIR /tmp

# Getting Migrate CLI
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xz

# Install it to /bin
RUN cp /tmp/migrate /bin/migrate

# Install Air for hot reload
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# Create Folder for development
RUN mkdir /app

# Set Workdir to development
WORKDIR /app

# Add user for golang application (Please sync your uid & gid)
RUN addgroup -g 1000 www
RUN adduser -s /bin/sh -u 1000 -G www -D www

# Copy File
COPY --chown=www . /app

# Change owner root folder development to www
RUN chown www:www /app

# Change user
USER www

# Install Package
RUN go install

# Expose a port
EXPOSE 8080

# Temp, it will change to air
CMD air
