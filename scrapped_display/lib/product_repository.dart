import 'package:http/http.dart' as http;
import 'dart:convert';
import 'product.dart';
  
Future<Stream<Product>> getInfo() async {
 final String url = 'http://10.0.2.2:6000/fetcher/api/v1/getInfo';

 final client = new http.Client();
 final streamedRest = await client.send(
  http.Request('get', Uri.parse(url))
 );

 return streamedRest.stream
     .transform(utf8.decoder)
     .transform(json.decoder)
     .expand((data) => (data as List))
     .map((data) => Product.fromJSON(data));
}