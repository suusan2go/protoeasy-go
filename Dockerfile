FROM golang:1.7.3
MAINTAINER peter.edge@gmail.com

RUN \
  apt-get update -yq && \
  apt-get install -yq --no-install-recommends \
    autoconf \
    automake \
    build-essential \
    git \
    libtool \
    unzip && \
  apt-get clean && \
  rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN \
  wget https://codeload.github.com/google/protobuf/tar.gz/v3.0.2 && \
  tar xvzf v3.0.2 && \
  rm v3.0.2 && \
  cd protobuf-3.0.2 && \
  ./autogen.sh && \
  ./configure --prefix=/usr && \
  make && \
  make check && \
  make install && \
  cd - && \
  rm -rf protobuf-3.0.2

RUN \
  git clone https://github.com/grpc/grpc.git && \
  cd grpc && \
  git checkout v1.0.0 && \
  git submodule update --init && \
  make && \
  make install

RUN mkdir -p /go/src/go.pedge.io/protoeasy
ADD . /go/src/go.pedge.io/protoeasy/
WORKDIR /go/src/go.pedge.io/protoeasy
RUN make install
CMD ["protoeasyd"]
