all:
	go build

docker:
	docker build -t spetr/mundi .

docker-run:
	docker run -it --rm spetr/mundi

update:
	go get -u github.com/kardianos/service
	go get -u github.com/gin-gonic/gin
	go get -u github.com/gin-contrib/static

	go get -u github.com/h2non/filetype
	go get -u code.sajari.com/docconv
