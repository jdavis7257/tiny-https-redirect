FROM busybox

ADD src/redirect /redirect
RUN chmod +x redirect
ENTRYPOINT redirect