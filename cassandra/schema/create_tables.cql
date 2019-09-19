USE scraper;

CREATE TABLE productpageinfo (
	pid bigint,
	purlid bigint,
	pname text,
	purl text,
	isexpired int,
	totaltries bigint,
	totalmisses bigint,
	createdon timestamp,
	modifiedon timestamp,
	modifiedby text,
	PRIMARY KEY ((pid, purlid), pname)
	) WITH CLUSTERING ORDER BY (pname DESC);

CREATE TABLE productinfo (
	pid bigint,
	purlid bigint, 
	metainfoid bigint,
	imagegroup list<bigint>,
	createdon timestamp,
	modifiedon timestamp,
	PRIMARY KEY ((purlid), pid)
	) WITH CLUSTERING ORDER BY (pid DESC);

CREATE TABLE productimageinfo (
	purlid bigint,
	iurlid bigint,
	iurl text,
	createdon timestamp,
	modifiedon timestamp,
	PRIMARY KEY ((iurlid), purlid)
	) WITH CLUSTERING ORDER BY (purlid DESC);

CREATE TABLE productmetainfo (
	pid bigint,
	purlid bigint,
	metainfoid bigint,  
	title text, 
	description text,
	retailer text,
	price text,
	seller text,
	createdon timestamp,
	modifiedon timestamp,
	PRIMARY KEY ((purlid), pid) 
	) WITH CLUSTERING ORDER BY (pid DESC);


