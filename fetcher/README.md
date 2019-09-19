#Getting started

Download image from [docker hub](https://cloud.docker.com/repository/docker/thiyageshv/product-scraping) or build the docker file from parent project directory 
 ```
 docker build -f /fetcher/Dockerfile .

```

Then run 

```
docker run -d -p 6000:6000 <image name>

```

You can interact with container through a bunch of API calls 

```
GET /fetcher/api/v1/getInfo returns ImageURLS information and metainformation (check lib_cassandra for model structure)

POST /fetcher/api/v1/addInfo adds a record to the productpageinfo table 

Bodyparams (xxx-form-url-encoded):
	productname: <name> 
	producturl: <url> 

GET /fetcher/api/v1/getMetrics returns metric info (hits and misses) from productpageinfo table. 

```
