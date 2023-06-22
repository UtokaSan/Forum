FROM golang:latest

RUN mkdir ../home/app
WORKDIR /../home/app
COPY . .

CMD ["go", "run", "."]
