FROM ortools/make:debian_swig AS env
ENV GOROOT=/usr/local/go
ENV GOPATH=/home/go
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin
ENV GO111MODULE=on
RUN wget https://golang.org/dl/go1.16.3.linux-amd64.tar.gz \
    && mkdir -p /usr/local/go \
    && mkdir -p /home/go \
    && tar -xvf go1.16.3.linux-amd64.tar.gz -C /usr/local \
    && rm -rf go1.16.3.linux-amd64.tar.gz \
    && go get github.com/golang/protobuf/protoc-gen-go@v1.5.2
RUN make -version

FROM env AS devel
WORKDIR /home/project
COPY . .

FROM devel AS build
RUN make third_party
RUN make clean_go go
RUN ldconfig

FROM build AS test
RUN make test_go

FROM build AS package
RUN make golib_archive

FROM scratch AS export
COPY --from=package /home/project/or-tools_Debian* .
