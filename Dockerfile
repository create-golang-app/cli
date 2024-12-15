FROM scratch
COPY gocli /
ENTRYPOINT ["/gocli"]