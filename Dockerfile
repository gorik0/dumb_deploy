FROM alpine

RUN sleep 2

ENTRYPOINT [ "echo", "-e"]
CMD [ "hmelo"]
