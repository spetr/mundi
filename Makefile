all:
	go build -tags ocr -a -o mundi

run:
	./mundi

docker:
	docker build -t spetr/mundi .

docker-run:
	docker run --interactive --tty --rm -p 25794:25794 --name mundi spetr/mundi

docker-push:
	docker push spetr/mundi

update:
	go get -u github.com/pkg/errors
	go get -u github.com/kardianos/service
	go get -u github.com/gin-gonic/gin
	go get -u github.com/gin-contrib/static

	go get -u github.com/vimeo/go-magic/magic
	go get -u github.com/otiai10/gosseract
	go get -tags ocr -u github.com/spetr/docconv
