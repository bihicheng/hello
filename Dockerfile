FROM golang:1.13.7
EXPOSE 8080
COPY ./hello /usr/local/bin/
CMD ["hello"]
