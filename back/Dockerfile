FROM golang:1.19-alpine
RUN mkdir -p /app
WORKDIR app
COPY . .
RUN go mod download
RUN go install github.com/cosmtrek/air@latest
EXPOSE 3001
CMD ["air", "run"]
