FROM alpine
COPY bin/helpwave /usr/local/bin/helpwave
RUN chmod +x /usr/local/bin/helpwave
EXPOSE 3000
CMD /usr/local/bin/helpwave
