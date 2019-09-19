### Cassandra Database Schema

This folder has all the schema files to set up a Cassandra database with the project schema. A single node cassandra service is set in the container and schema is applied after cassandra is up. 

## Getting started

Install cqlsh. Download the cassandra image from [docker hub](https://cloud.docker.com/repository/docker/thiyageshv/product-scraping) or build the dockerfile and 
 ``` 
   docker run -p 9042:9042 -d <image name>
 ```
You can then access the db and the tables through cqlsh. Simply type cqlsh from your command line. 

The default data in the cassandra db for product information is taken from input.json file. You can add or remove entities and reconstruct the docker for the state to be reflected.   
