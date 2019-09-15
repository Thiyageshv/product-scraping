#!/bin/bash
keyspace="scraper"
tablename="productinfo"
cqlsh -f create_keyspaces.cql
cqlsh -f create_tables.cql
cqlsh -f create_indexes.cql
author=`echo "$(<input.json)" | jq -r ".author"`
timestamp=$(date +"%Y-%m-%d %H:%M:%S")
while read name id ; do
    pid=`echo "$name" | cksum | cut -f 1 -d ' '`
    pname=`echo "$name"` 
    purlid=`echo "$id" | cksum | cut -f 1 -d ' '`
    purl=`echo "$id"`
    query="INSERT INTO $keyspace.$tablename (purl, createdon, modifiedon, isexpired, pid, pname, purlid, modifiedby)  VALUES ('$purl', '$timestamp', '$timestamp', 0, $pid, '$pname', $purlid, '$author')"
    echo $query 
    cqlsh -e "$query"
done < <(echo "$(<input.json)" | jq -r '.payload[]|"\(.productname) \(.producturl)"')

