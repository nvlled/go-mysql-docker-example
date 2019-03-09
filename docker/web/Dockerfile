FROM golang:1.8
LABEL author="ronald"
LABEL description="a sample docker image"

WORKDIR /go/src/app
COPY . .

# Setup ssh keys for accessing private repositories
ARG ssh_prv_key
ARG ssh_pub_key

## Add know hosts
RUN mkdir -p /root/.ssh && \
    chmod 0700 /root/.ssh && \
    ssh-keyscan github.com > /root/.ssh/known_hosts

## Add the keys and set permissions
RUN echo "$ssh_prv_key" > /root/.ssh/id_rsa && \
    echo "$ssh_pub_key" > /root/.ssh/id_rsa.pub && \
    chmod 600 /root/.ssh/id_rsa && \
    echo "XY =" "$XY"

## make sure your domain is accepted
RUN touch /root/.ssh/known_hosts
RUN ssh-keyscan doctype.se >> /root/.ssh/known_hosts
RUN ssh-keyscan brewery.doctype.se >> /root/.ssh/known_hosts

## Use ssh when cloning for these domains (for private repositories)
RUN git config --global --add url."git@example.com:".insteadOf "https://example.com/"

# Copy sql files for initialization
#RUN mkdir -p ./schema /docker-entrypoint-initdb.d
#RUN mv -vfu ./schema/* /docker-entrypoint-initdb.d/

# Install go dependencies
RUN go get -d -v ./...
RUN go install -v ./...

CMD ["./wait-for-it.sh", "-t", "90", "db:3306", "--", "app"]
