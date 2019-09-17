USE scraper;

CREATE INDEX PNAME ON productpageinfo(pname);

CREATE INDEX PID ON productinfo(pid);

CREATE INDEX METAINFOID ON productinfo(metainfoid);

CREATE INDEX PURLID ON productimageinfo(purlid)

CREATE INDEX RETAILER ON productmetainfo(retailer);

CREATE INDEX SELLER ON productmetainfo(seller);	


