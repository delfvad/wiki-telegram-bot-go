FROM alpine
ENV LANGUAGE="en"
COPY /src/src .
RUN apk add --no-cache ca-certificates &&\
    chmod +x src
EXPOSE 80/tcp
CMD [ "./src" ]