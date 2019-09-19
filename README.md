# Product-scraping

Wlecome, what you are seeing is a prototype microservices model to demonstrate a product image scrapping application. The application scrapes product URLs from a database, fetches images and meta information in the URL and displays them in an Android App that created with Flutter. This application is designed to scale easily and support a lot more features (For eg: search indexin). It might therefore might look over-engineered for the purpose - isn't it better than under engineering? Also, this whole codebase, design and documentation offers a peek into how I would normally approach and solve a Software Engineering problemn. Please not thate this application is merely a proof of concept and in order to develop this whole application in record time, I have left out implementing some of the core features such as authentication/authorization, transaction rollback, CI pipelines and so on. 

## Components

There are 5  components to the whole application, 

1. scrapped_display aka the frontend - The Flutter code for displaying information in an Android Application
2. cassandra - A distributed db model to host product information and all scraped information
3. scraper - A Python service that scrapes a URL and returns image and meta information for a given API call
4. cacher - A Golang service that reads product info from database, calls scraper API on them and stores the retrieved information into Cassandra tables
5. fetcher - Fetches stored information from Cassandra tables and makes them available to frontend through an API call

## Get Started

###### Backend
Easy! Make sure you have installed docker, docker-compose. Go to the project root folder and 
```
docker-compose up

```

That is all it takes seriously! The images are hosted in the [docker hub](https://cloud.docker.com/repository/docker/thiyageshv/product-scraping). The command fetches and builds the images on a separate bridge network that would be accessible by localhost.

###### Frontend

Open the APK file residing in scrapped_display folder on Android Studio and run it through an emulator. Make sure the backend is up before that. 

## Available APIs

###### External 

1. /fetcher/api/v1/getInfo - Gets imageurl information and meta information from db
2. /fetcher/api/v1/getMetrics - Gets product page information and associated metrics, number of scraping failures, total tries, link expiry and so on.
3. /fetcher/api/v1/addInfo - Adds product url and product name into the productpage info table for it to be scraped in the next cycle

Use [this](https://www.getpostman.com/collections/b29f797b51d162b2a012) link to download the POSTMAN collection to try them. 
###### Internal
1. /scraper/api/v1/scrape - takes product URL as a parameter and returns Image URL inside the page and all meta information such as seller, price, product description and so on. 
    
