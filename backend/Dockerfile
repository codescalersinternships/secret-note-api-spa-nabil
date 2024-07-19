#Build Stage
FROM golang:1.22.4-alpine AS BuildStage

WORKDIR /backend

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the app code
COPY . ./

# Install golint
RUN go get -u golang.org/x/lint/golint

# Install gin
RUN go get -u github.com/gin-gonic/gin

# Install endless
RUN go get -u github.com/fvbock/endless

# Install assert
RUN go get github.com/stretchr/testify/assert

# Install bcrypt
RUN go get golang.org/x/crypto/bcrypt

# Install gorm
RUN go get -u gorm.io/gorm
RUN go get gorm.io/driver/postgres

# Install jwt
RUN go get github.com/golang-jwt/jwt/v5

# Install uuid
RUN go get github.com/google/uuid

# Build the Go binary
RUN go build -o secretnoteserver ./cmd/main.go

EXPOSE 8090

#Deploy Stage
FROM alpine:latest As secretnoteserver
WORKDIR /
COPY .env .env 
COPY --from=BuildStage /backend/secretnoteserver /secretnoteserver

EXPOSE 8090

ENTRYPOINT [ "/secretnoteserver" ]