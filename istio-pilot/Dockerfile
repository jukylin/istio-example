FROM alpine 
RUN apk add --update bash && rm -rf /var/cache/apk/*
ADD ./pilot-discovery /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/pilot-discovery"]
