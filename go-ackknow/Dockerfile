FROM golang:1.16

WORKDIR /app

COPY . .

RUN apt-get update && apt-get install -y bpfcc-tools git build-essential cmake libelf-dev && \
    apt-get install -y gcc-7 g++-7 && \
    update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc-7 100 && \
    update-alternatives --install /usr/bin/g++ g++ /usr/bin/g++-7 100 && \
    git clone --depth 1 https://github.com/iovisor/bcc.git /tmp/bcc && \
    mkdir /tmp/bcc/build && cd /tmp/bcc/build && \
    cmake .. && make -j$(nproc) && make install && \
    cd / && rm -rf /tmp/bcc && \
    go get -d -v ./... && \
    go build -o /app/main .

CMD ["/app/main"]
