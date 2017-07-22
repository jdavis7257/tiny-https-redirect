FROM busybox

ADD redirect /redirect
RUN chmod +x redirect
ENTRYPOINT ./redirect