FROM curlimages/curl:latest

WORKDIR /app
COPY ratetester.sh .

ENTRYPOINT ["/app/ratetester.sh"]
