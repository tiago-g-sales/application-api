FROM alpine:3

RUN mkdir /application && addgroup -S app && adduser -S app -G app \
    && chown -R app:app  /application

USER app

COPY --chown=app application /application/app
COPY --chown=app ./migrations /application/migrations
COPY --chown=app ./docs /application/docs

WORKDIR /application

EXPOSE 8080

ENTRYPOINT ["./app"]
