FROM ubuntu:disco AS builder
RUN apt-get update && \
    apt-get install -y golang-1.12-go git libmagic-dev libtesseract-dev && \
    rm -rf /var/lib/apt/lists/*
RUN /usr/lib/go-1.12/bin/go get github.com/pkg/errors && \
    /usr/lib/go-1.12/bin/go get github.com/kardianos/service && \
    /usr/lib/go-1.12/bin/go get github.com/gin-gonic/gin && \
    /usr/lib/go-1.12/bin/go get github.com/gin-contrib/static && \
    /usr/lib/go-1.12/bin/go get github.com/vimeo/go-magic/magic && \
    /usr/lib/go-1.12/bin/go get -t github.com/otiai10/gosseract && \
    /usr/lib/go-1.12/bin/go get -tags ocr github.com/spetr/docconv
WORKDIR /project
COPY *.go ./
RUN CGO_ENABLED=1 GOOS=linux /usr/lib/go-1.12/bin/go build -a -o mundi

FROM ubuntu:disco
RUN apt-get update && \
    apt-get install -y libmagic1 poppler-utils wv unrtf tidy tesseract-ocr && \
    rm -rf /var/lib/apt/lists/*
COPY --from=builder /project/mundi /mundi
COPY mundi.yaml /
COPY www/* /www/
EXPOSE 25794/tcp
ENTRYPOINT ["/mundi"]