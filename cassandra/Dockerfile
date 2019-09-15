FROM cassandra
RUN apt-get -y update && apt-get -y install jq
ADD . /opt/src/cassandra
COPY entrypoint-wrap.sh /entrypoint-wrap.sh
RUN chmod 777 /entrypoint-wrap.sh
ENV CASSANDRA_KEYSPACE=scraper
WORKDIR /opt/src/cassandra/schema
ENTRYPOINT ["/entrypoint-wrap.sh"]
CMD ["cassandra", "-f"]

