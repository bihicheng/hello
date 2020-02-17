FROM alpine
EXPOSE 8080
COPY hello /bin
ENTRYPOINT [ "/bin/hello" ]
