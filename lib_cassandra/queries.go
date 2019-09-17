package cassandra 


var getPageBasicInfo = `SELECT pid, purlid, pname, purl FROM scraper.productpageinfo ALLOW FILTERING`
var getPageMetricInfo = `SELECT pid, purlid, pname, isexpired, totaltries, totalmisses, modifiedby, createdon, modifiedon FROM scraper.productpageinfo ALLOW FILTERING`
var updateMetrics = `UPDATE scraper.productpageinfo SET isexpired=%d, totaltries=%d, totalmisses=%d WHERE pid=%d AND purlid=%d AND pname='%s'`

var getProductInfo = `SELECT pid, purlid, metainfoid, imagegroup, createdon, modifiedon FROM scraper.productinfo WHERE purlid=%d AND pid=%d`
var addProductInfo = `INSERT INTO scraper.productinfo (pid, purlid, metainfoid, imagegroup, createdon, modifiedon) VALUES (%d, %d, %d, %v, toUnixTimestamp(now()), toUnixTimestamp(now()))`
var addImageID = `UPDATE scraper.productinfo SET imagegroup = imagegroup + [%d] WHERE purlid=%d AND pid=%d`

var getProductMetaInfo = `SELECT pid, purlid, metainfoid, description, retailer, price, seller, createdon, modifiedon  FROM scraper.productmetainfo WHERE purlid=%d AND pid=%d`
var addProductMetaInfo = `INSERT INTO scraper.productmetainfo (pid, purlid, metainfoid, description, retailer, price, seller, createdon, modifiedon) VALUES (%d, %d, %d, '%s', '%s', '%s', '%s', toUnixTimestamp(now()), toUnixTimestamp(now()))`


var getProductImageInfo = `SELECT purlid, iurlid, iurl, createdon, modifiedon FROM scraper.productimageinfo WHERE iurlid=%d AND purlid=%d`
var addProductImage = `INSERT INTO scraper.productimageinfo (purlid, iurlid, iurl, createdon, modifiedon) VALUES (%d, %d, '%s', toUnixTimestamp(now()), toUnixTimestamp(now()))`
