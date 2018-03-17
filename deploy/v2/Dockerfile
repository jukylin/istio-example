FROM alpine
RUN apk add --update bash && rm -rf /var/cache/apk/*
ADD ./deploy-v2 /deploy-v2
EXPOSE 8080
CMD ["/deploy-v2"]