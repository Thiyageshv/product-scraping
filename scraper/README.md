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

 
