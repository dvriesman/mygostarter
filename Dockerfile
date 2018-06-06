FROM scratch
ADD ca-certificates.crt /etc/ssl/certs/
ADD build/main /
CMD ["/main"]