FROM alpine:3.15
COPY wallet-api  ./config/settings.yml VERSION ./
EXPOSE 8821
ENTRYPOINT ["./wallet-example"]
CMD ["start", "-c", "./settings.yml"]