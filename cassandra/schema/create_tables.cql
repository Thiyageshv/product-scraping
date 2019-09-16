USE scraper;

CREATE TABLE productinfo (
	pid bigint,
	purlid bigint,
	pname text,
	purl text,
	isexpired int,
	createdon timestamp,
	modifiedon timestamp,
	modifiedby text,
	PRIMARY KEY (purl, createdon, modifiedon)
	) WITH CLUSTERING ORDER BY (createdon DESC, modifiedon DESC);

CREATE TABLE productimageinfo (
	pid bigint,
	purlid bigint, 
	iurlid bigint,
	metainfoid bigint,
	iurl text,
	createdon timestamp,
	modifiedon timestamp,
	PRIMARY KEY ((iurl, metainfoid), createdon, modifiedon)
	) WITH CLUSTERING ORDER BY (createdon DESC, modifiedon DESC);

CREATE TABLE productmetainfo (
	pid bigint,
	purlid bigint,
	metainfoid bigint,  
	description text,
	retailer text,
	price float,
	seller text,
	createdon timestamp,
	modifiedon timestamp,
	PRIMARY KEY ((metainfoid), price) 
	) WITH CLUSTERING ORDER BY (price ASC);


