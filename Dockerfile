FROM alpine
EXPOSE 8080
ADD ./hello /bin
ENTRYPOINT [ "/bin/hello" ]
