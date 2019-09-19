#Overview 

The project uses beautifulsoup library to parse through the web page. It works only on product pages. There are individual classes to parse Amazon(flaky), Ebay, Walmart and Alibaba. It is easy to add more classes. There is also a generic class but it is not guranteed to work on all pages and websites. 

# Getting started

Download the scraper image from the [docker hub](https://cloud.docker.com/repository/docker/thiyageshv/product-scraping) or build docker file and run
```
docker run -p 5000:5000 -d <image name>

```
Once the container starts running, you can give it instructions through the scrape API call 

```
POST http://127.0.0.1:5000/scraper/api/v1/scrape

Bodyparams:
{
	"producturl": <product url>

}

```

 
