FROM ubuntu:16.04
RUN apt-get update && \
    apt-get install -y git curl make && \
    curl -O https://storage.googleapis.com/golang/go1.10.1.linux-amd64.tar.gz && \
    tar -xvf go1.10.1.linux-amd64.tar.gz
ADD . /opt/src/product-scraping/
ENV GOPATH=/opt/
ENV GOROOT=/go
ENV PATH=$PATH:$GOROOT/bin
ENV GOBIN=/opt/pkg
WORKDIR /opt/src/product-scraping/fetcher
RUN /bin/bash -c "make" && chmod 777 start_app.sh
ENTRYPOINT ./start_app.sh
