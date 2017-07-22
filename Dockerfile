FROM alpine

ADD redirect /redirect
RUN chmod +x redirect
ENTRYPOINT /redirect