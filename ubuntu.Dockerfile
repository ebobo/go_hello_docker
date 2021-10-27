FROM ubuntu:20.04

RUN ln -fs  /usr/share/zoneinfo/Europe/Oslo /etc/localtime
RUN apt -y update && apt-get install -y --no-install-recommends
RUN apt install clang cmake wget git libboost-all-dev libtinyxml2-dev -y

RUN wget https://dl.google.com/go/go1.17.2.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.17.2.linux-amd64.tar.gz

# Environment replacement
ENV PATH="/usr/local/go/bin:${PATH}"

WORKDIR /app
COPY . ./


RUN go mod download
RUN go build -o /hello

# CMD ["go","version"]

# Start app
CMD ["/hello","greeting","-n qi" ]