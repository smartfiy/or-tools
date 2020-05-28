FROM ortools/make:debian_swig AS env
ENV GOROOT=/usr/local/go
ENV GOPATH=/home/go
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin
ENV GO111MODULE=on
RUN apt-get update -qq \
    && wget https://dl.google.com/go/go1.14.1.linux-amd64.tar.gz \
    && mkdir -p /usr/local/go \
    && mkdir -p /home/go \
    && tar -xvf go1.14.1.linux-amd64.tar.gz -C /usr/local \
    && rm -rf go1.14.1.linux-amd64.tar.gz \
    && go get github.com/golang/protobuf/protoc-gen-go@v1.3 \
    && apt-get install -yq  \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
RUN make -version

FROM env AS devel
WORKDIR /home/project
COPY . .

FROM devel AS build
RUN make third_party
RUN make clean_go install_go
RUN ldconfig

FROM build AS test
RUN make test_go

FROM build AS package
RUN make golib_archive

FROM scratch AS export
COPY --from=package /home/project/or-tools_Debian* .
