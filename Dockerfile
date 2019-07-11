FROM ubuntu:disco AS builder
RUN apt-get update && \
    apt-get install -y golang-1.12-go libmagic-dev git && \
    rm -rf /var/lib/apt/lists/*
RUN /usr/lib/go-1.12/bin/go get github.com/spetr/mundi
WORKDIR /project
COPY *.go ./
RUN CGO_ENABLED=1 GOOS=linux /usr/lib/go-1.12/bin/go build -a -o mundi

FROM ubuntu:disco
RUN apt-get update && \
    apt-get install -y poppler-utils wv unrtf tidy && \
    rm -rf /var/lib/apt/lists/*
COPY --from=builder /project/mundi /mundi
ENTRYPOINT ["/mundi"]