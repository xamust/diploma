FROM golang
MAINTAINER Stepan K. <xamust@gmail.com>
WORKDIR /server/
VOLUME ["/opt/server"]
COPY server ./
RUN make build
EXPOSE 8585
EXPOSE 8383
CMD [ "build/server" ]
