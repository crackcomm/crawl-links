FROM alpine:latest
MAINTAINER Łukasz Kurowski <crackcomm@gmail.com>

# Install certificates
RUN apk --update add ca-certificates

# Copy application
COPY ./dist/crawl-links /crawl-links

#
# Environment variables for crawler
#
ENV TOPIC crawl_links
ENV CHANNEL consumer
ENV NSQ_ADDR nsq:4150
ENV NSQLOOKUP_ADDR nsqlookup:4161
ENV CONCURRENCY 100

ENTRYPOINT ["/crawl-links"]
