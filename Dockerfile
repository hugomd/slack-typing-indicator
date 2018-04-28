FROM scratch
ADD main /
ADD ca-certificates.crt /etc/ssl/certs/
CMD ["/main"]
