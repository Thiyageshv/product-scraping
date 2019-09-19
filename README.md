# Product-scraping

Mini microservices model to demonstrate a product image scrapping application. The application scrapes product URLs from a database, fetches images and meta information and displays them in an Android App that is created with Flutter. This application is designed with scalability in mind and therefore might look over-engineered for the purpose - isn't it better than under engineering? 

## Components

There are 5  components to the whole application, 

1. scrapped_display - The Flutter code for displaying information in an Android Application
2. cassandra - A distributed db model to host product information and all scraped information
3. scraper - A Python service that scrapes a URL and returns image and meta information for a given API call
4. cacher - A Golang service that reads product info from database, calls scraper API on them and stores the retrieved information into Cassandra tables
5. fetcher - Fetches stored information from Cassandra tables and makes them available to frontend through an API call

## Get Started

### Backend
```
docker-compose up

```

That is all it takes seriously! The images are hosted in the docker hub. The command fetches and builds the images on a separate bridge network that would be accessible by localhost. 

### Frontend

Open the APK file on Android Studio and run it through an emulator. Make sure the backend is up before that. 

## Available APIs

### External 

1. /fetcher/api/v1/getInfo - Gets imageurl information and meta information from db
2. /fetcher/api/v1/getMetrics - Gets product page information and associated metrics, number of scraping failures, total tries, link expiry and so on.
3. /fetcher/api/v1/addInfo - Adds product url and product name into the productpage info table for it to be scraped in the next cycle

### Internal
1. /scraper/api/v1/scrape - takes product URL as a parameter and returns Image URL inside the page and all meta information such as seller, price, product description and so on. 
    
