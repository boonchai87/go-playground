# https://docs.docker.com/language/golang/build-images/
FROM golang:1.22.1

# Set destination for COPY
WORKDIR /app

#https://stackoverflow.com/questions/42385977/accessing-a-docker-container-from-another-container
ENV GO_PORT=8080
ENV GO_URL=localhost:8080
ENV POSTGRES_HOST=host.docker.internal
ENV POSTGRES_PORT=5432
ENV POSTGRES_USERNAME=postgres
ENV POSTGRES_DBNAME=postgres
ENV POSTGRES_PASSWORD=root
ENV POSTGRES_SSLMODE=disable
ENV GIN_MODE=release

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
# COPY *.go ./
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080

# Run
CMD ["/docker-gs-ping"]