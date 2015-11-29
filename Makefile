PKGS := $(shell go list ./... | grep -v 'go.pedge.io/protoeasy/vendor')

export GO15VENDOREXPERIMENT=1

all: build docker-build docker-launch install installplugins proto example

deps:
	GO15VENDOREXPERIMENT=0 go get -d -v $(PKGS)

updatedeps:
	GO15VENDOREXPERIMENT=0 go get -d -v -u -f $(PKGS)

testdeps:
	GO15VENDOREXPERIMENT=0 go get -d -v -t $(PKGS)

updatetestdeps:
	GO15VENDOREXPERIMENT=0 go get -d -v -t -u -f $(PKGS)

vendor:
	go get -v github.com/tools/godep
	rm -rf Godeps
	rm -rf vendor
	GOOS=linux GOARCH=AMD64 godep save \
			 $(PKGS) \
			 github.com/gengo/grpc-gateway/protoc-gen-grpc-gateway/... \
			 github.com/gengo/grpc-gateway/third_party/googleapis/... \
			 github.com/golang/protobuf/proto/... \
			 github.com/golang/protobuf/protoc-gen-go/... \
			 github.com/gogo/protobuf/protoc-gen-gofast/... \
			 github.com/gogo/protobuf/protoc-gen-gogo/... \
			 github.com/gogo/protobuf/protoc-gen-gogofast/... \
			 github.com/gogo/protobuf/protoc-gen-gogofaster/... \
			 github.com/gogo/protobuf/protoc-gen-gogoslick/... \
			 go.pedge.io/google-protobuf/... \
			 go.pedge.io/googleapis/... \
			 google.golang.org/grpc
	rm -rf Godeps

build:
	go build $(PKGS)

install:
	go install $(PKGS)

installplugins:
	go install \
		go.pedge.io/protoeasy/vendor/github.com/gengo/grpc-gateway/protoc-gen-grpc-gateway \
		go.pedge.io/protoeasy/vendor/github.com/golang/protobuf/protoc-gen-go \
		go.pedge.io/protoeasy/vendor/github.com/gogo/protobuf/protoc-gen-gofast \
		go.pedge.io/protoeasy/vendor/github.com/gogo/protobuf/protoc-gen-gogo \
		go.pedge.io/protoeasy/vendor/github.com/gogo/protobuf/protoc-gen-gogofast \
		go.pedge.io/protoeasy/vendor/github.com/gogo/protobuf/protoc-gen-gogofaster \
		go.pedge.io/protoeasy/vendor/github.com/gogo/protobuf/protoc-gen-gogoslick

proto: install
	go get -v go.pedge.io/pkg/cmd/strip-package-comments
	protoeasy --go --grpc --go-import-path go.pedge.io/protoeasy --exclude vendor --exclude example .
	find . -name *\.pb\*\.go | grep -v vendor | xargs strip-package-comments

example-complete: install
	rm -rf _example-out/complete
	protoeasy \
		--out=_example-out/complete \
		--cpp --cpp-rel-out=cpp \
		--csharp --csharp-rel-out=csharp \
		--objc --objc-rel-out=objc \
		--python --python-rel-out=python \
		--ruby --ruby-rel-out=ruby \
		--go --go-rel-out=go --go-import-path=go.pedge.io/protoeasy/_example-out/complete/go \
		--grpc \
		--grpc-gateway \
		example/complete
	cd _example-out/complete/go && go build ./...

example-proto2: install
	rm -rf _example-out/proto2
	protoeasy \
		--out=_example-out/proto2 \
		--cpp --cpp-rel-out=cpp \
		--python --python-rel-out=python \
		--go --go-rel-out=go --go-import-path=go.pedge.io/protoeasy/_example-out/proto2/go \
		--go-protoc-plugin=gogo \
		--grpc \
		example/proto2
	cd _example-out/proto2/go && go build ./...

lint:
	go get -v github.com/golang/lint/golint
	for file in $$(find . -name '*.go' | grep -v '^\./_example-out' | grep -v '^\./vendor' | grep -v '\.pb\.go'); do \
		golint $$file; \
		if [ -n "$$(golint $$file)" ]; then \
			exit 1; \
		fi; \
	done

vet:
	go vet $(PKGS)

errcheck:
	go get -v github.com/kisielk/errcheck
	errcheck $(PKGS)

pretest: lint vet errcheck

test: pretest
	go test $(PKGS)

clean:
	go clean $(PKGS)
	rm -rf _example-out

docker-build:
	docker build -t quay.io/pedge/protoeasy .

docker-push: docker-build
	docker push quay.io/pedge/protoeasy

docker-pull:
	docker pull quay.io/pedge/protoeasy

docker-launch:
	docker rm -f protoeasy || true
	docker run -d -p 6789:6789 --name=protoeasy quay.io/pedge/protoeasy

docker-test: docker-build
	docker run quay.io/pedge/protoeasy make test

docker-build-proto2:
	docker build -t quay.io/pedge/protoeasy-proto2 -f Dockerfile.proto2 .

docker-push-proto2: docker-build-proto2
	docker push quay.io/pedge/protoeasy-proto2

docker-pull-proto2:
	docker pull quay.io/pedge/protoeasy

docker-launch-proto2:
	docker rm -f protoeasy-proto2 || true
	docker run -d -p 6789:6789 --name=protoeasy-proto2 quay.io/pedge/protoeasy-proto2

.PHONY: \
	all \
	deps \
	updatedeps \
	testdeps \
	updatetestdeps \
	vendor \
	build \
	install \
	installplugins \
	proto \
	example-complete \
	example-proto2 \
	lint \
	vet \
	errcheck \
	pretest \
	test \
	clean \
	docker-build \
	docker-push \
	docker-pull \
	docker-launch \
	docker-test \
	docker-build-proto2 \
	docker-push-proto2 \
	docker-pull-proto2 \
	docker-launch-proto2 \
