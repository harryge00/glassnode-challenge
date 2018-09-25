all: build-basic-image build-all run-all

VERSION?=v0.0.1

build-basic-image:
	docker build -t basic-go .

build-price:
	docker build -t price:$(VERSION) docker/price

build-ranking:
	docker build -t ranking:$(VERSION) docker/ranking

build-httpserver:
	docker build -t httpserver:$(VERSION) docker/httpserver

build-all: build-basic-image build-price build-ranking build-httpserver
	# build all the 3 images

run-all:
	docker run -d -p 8085:8080 --name price price:$(VERSION)
	docker run -d -p 8084:8080 --name ranking ranking:$(VERSION)
	docker run -d --net=host -e RANK_ADDR=localhost -e RANK_PORT=8084 -e PRICE_ADDR=localhost -e PRICE_PORT=8085 --name httpserver httpserver:$(VERSION)

clean:
	docker stop price ranking httpserver
	docker rm price ranking httpserver
