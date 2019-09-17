FROM ubuntu:16.04
RUN \
 apt-get update && \
 apt-get install -y nginx && \
 apt-get install -y curl && \
 apt-get install -y git && \
 apt-get install -y net-tools && \
 apt-get install -y python3 && \
 apt-get install -y python3-pip && \
 pip3 install requests && \
 pip3 install beautifulsoup4 && \
 pip3 install Flask && \
 pip3 install tldextract
EXPOSE 5000
ADD . /opt/src/scraper/
WORKDIR /opt/src/scraper/
RUN chmod 777 start_app.sh
ENTRYPOINT ./start_app.sh
