PKGS := $(shell go list ./... | grep -v 'go.pedge.io/protoeasy/vendor')

CMDS := \
	go.pedge.io/protoeasy/cmd/protoeasy \
	go.pedge.io/protoeasy/cmd/protoeasyd

EXTRA_PKGS := \
	github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
	github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
	github.com/grpc-ecosystem/grpc-gateway/runtime \
	github.com/golang/protobuf/protoc-gen-go \
	github.com/gogo/protobuf/protoc-gen-gofast \
	github.com/gogo/protobuf/protoc-gen-gogo \
	github.com/gogo/protobuf/protoc-gen-gogofast \
	github.com/gogo/protobuf/protoc-gen-gogofaster \
	github.com/gogo/protobuf/protoc-gen-gogoslick \
	go.pedge.io/pb/go/google/api \
	go.pedge.io/pb/go/google/devtools/cloudtrace/v1 \
	go.pedge.io/pb/go/google/example/library/v1 \
	go.pedge.io/pb/go/google/iam/v1 \
	go.pedge.io/pb/go/google/logging/type \
	go.pedge.io/pb/go/google/logging/v2 \
	go.pedge.io/pb/go/google/longrunning \
	go.pedge.io/pb/go/google/pubsub/v1 \
	go.pedge.io/pb/go/google/pubsub/v1beta2 \
	go.pedge.io/pb/go/google/rpc \
	go.pedge.io/pb/go/google/type \
	go.pedge.io/pb/go/pb/common \
	go.pedge.io/pb/go/pb/geo \
	go.pedge.io/pb/go/pb/money \
	go.pedge.io/pb/go/pb/net \
	go.pedge.io/pb/go/pb/phone \
	go.pedge.io/pb/go/pb/time \
	go.pedge.io/pb/gogo/google/api \
	go.pedge.io/pb/gogo/google/devtools/cloudtrace/v1 \
	go.pedge.io/pb/gogo/google/example/library/v1 \
	go.pedge.io/pb/gogo/google/iam/v1 \
	go.pedge.io/pb/gogo/google/logging/type \
	go.pedge.io/pb/gogo/google/logging/v2 \
	go.pedge.io/pb/gogo/google/longrunning \
	go.pedge.io/pb/gogo/google/protobuf \
	go.pedge.io/pb/gogo/google/pubsub/v1 \
	go.pedge.io/pb/gogo/google/pubsub/v1beta2 \
	go.pedge.io/pb/gogo/google/rpc \
	go.pedge.io/pb/gogo/google/type \
	go.pedge.io/pb/gogo/pb/common \
	go.pedge.io/pb/gogo/pb/geo \
	go.pedge.io/pb/gogo/pb/money \
	go.pedge.io/pb/gogo/pb/net \
	go.pedge.io/pb/gogo/pb/phone

EXTRA_CMDS := \
	go.pedge.io/protoeasy/vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
	go.pedge.io/protoeasy/vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
	go.pedge.io/protoeasy/vendor/github.com/golang/protobuf/protoc-gen-go \
	go.pedge.io/protoeasy/vendor/github.com/gogo/protobuf/protoc-gen-gofast \
	go.pedge.io/protoeasy/vendor/github.com/gogo/protobuf/protoc-gen-gogo \
	go.pedge.io/protoeasy/vendor/github.com/gogo/protobuf/protoc-gen-gogofast \
	go.pedge.io/protoeasy/vendor/github.com/gogo/protobuf/protoc-gen-gogofaster \
	go.pedge.io/protoeasy/vendor/github.com/gogo/protobuf/protoc-gen-gogoslick

export GO15VENDOREXPERIMENT=1

all: build

testall: integration docker-integration

deps:
	GO15VENDOREXPERIMENT=0 go get -d -v $(PKGS) $(EXTRA_PKGS)

updatedeps:
	GO15VENDOREXPERIMENT=0 go get -d -v -u -f $(PKGS) $(EXTRA_PKGS)

testdeps:
	GO15VENDOREXPERIMENT=0 go get -d -v -t $(PKGS) $(EXTRA_PKGS)

updatetestdeps:
	GO15VENDOREXPERIMENT=0 go get -d -v -t -u -f $(PKGS) $(EXTRA_PKGS)

vendorupdate:
	GO15VENDOREXPERIMENT=0 GOOS=linux GOARCH=amd64 go get -d -v -t -u -f $(PKGS) $(EXTRA_PKGS)

vendornoupdate:
	go get -u -v github.com/kardianos/govendor
	rm -rf vendor
	govendor init
	GOOS=linux GOARCH=amd64 govendor add $(EXTRA_PKGS)
	GOOS=linux GOARCH=amd64 govendor add +external
	GOOS=linux GOARCH=amd64 govendor update +vendor
	mkdir -p vendor/go.pedge.io/pb
	cp -R ../pb/proto vendor/go.pedge.io/pb/proto

vendor: vendorupdate vendornoupdate

build:
	go build $(PKGS)

install:
	for i in $(CMDS) $(EXTRA_CMDS); do echo $$i; go install $$i; done

proto:
	go get -v go.pedge.io/pkg/cmd/strip-package-comments
	protoeasy
	find . -name *\.pb\*\.go | grep -v vendor | xargs strip-package-comments

example-complete:
	rm -rf example/out
	protoeasy \
		--no-file \
		--out=example/out/complete \
		--cpp --cpp-rel-out=cpp \
		--csharp --csharp-rel-out=csharp \
		--objc --objc-rel-out=objc \
		--python --python-rel-out=python \
		--ruby --ruby-rel-out=ruby \
		--go --go-rel-out=go --go-import-path=go.pedge.io/protoeasy/example/out/complete/go \
		--grpc \
		--grpc-gateway --extra-plugin-flag=swagger\
		example/complete
	go build ./example/out/complete/go/...
	rm -rf example/out

example-complete-file:
	rm -rf example/out
	cd example && protoeasy --out=out/complete && cd -
	go build ./example/out/complete/go/...
	rm -rf example/out

example-proto2:
	rm -rf example/out
	protoeasy \
		--out=example/out/proto2 \
		--cpp --cpp-rel-out=cpp \
		--python --python-rel-out=python \
		--go --go-rel-out=go --go-import-path=go.pedge.io/protoeasy/example/out/proto2/go \
		--gogo --gogo-rel-out=gogo --gogo-import-path=go.pedge.io/protoeasy/example/out/proto2/gogo \
		--grpc \
		example/proto2
	go build ./example/out/proto2/go/...
	#go build ./example/out/proto2/gogo/...
	rm -rf example/out

examples: install example-complete example-complete-file example-proto2

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

#pretest: lint vet errcheck
pretest: lint vet

test: pretest
	go test $(PKGS)

clean:
	go clean -i $(PKGS)
	rm -rf example/out

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

integration: build docker-build docker-launch proto examples

docker-integration: build docker-build
	docker run quay.io/pedge/protoeasy make proto examples

.PHONY: \
	all \
	testall \
	deps \
	updatedeps \
	testdeps \
	updatetestdeps \
	vendorupdate \
	vendornoupdate \
	vendor \
	build \
	install \
	proto \
	example-complete \
	example-complete-file \
	example-proto2 \
	examples \
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
	integration \
	docker-integration
