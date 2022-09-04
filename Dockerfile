FROM golang:1.19.0-buster

# Install Required Package
RUN apt update && apt install -y curl openssh-client gcc git make g++ python3

# Change Workdir to /tmp
WORKDIR /tmp

# Getting Migrate CLI
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xz

# Install it to /bin
RUN cp /tmp/migrate /bin/migrate

# Install sqlc
# RUN curl -L https://github.com/kyleconroy/sqlc/releases/download/v1.15.0/sqlc_1.15.0_linux_amd64.tar.gz | tar xz

# install it to /bin
# RUN cp /tmp/sqlc /bin/sqlc

# Install Air for hot reload
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# Create Folder for development
RUN mkdir /app

# Set Workdir to development
WORKDIR /app

# Add user for golang application (Please sync your uid & gid)
RUN addgroup --gid 1000 www
RUN useradd --shell /bin/sh -g www -m -u 1000 www

# Copy File
COPY --chown=www . /app

# Change owner root folder development to www
RUN chown www:www /app

# Change user
USER www

# Install sqlc
RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

# Install Package
RUN go install

# Expose a port
EXPOSE 8080

# Temp, it will change to air
CMD air
