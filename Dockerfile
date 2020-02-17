FROM golang:1.13.7
EXPOSE 8080
ADD hello /bin
ENTRYPOINT [ "/bin/hello" ]
