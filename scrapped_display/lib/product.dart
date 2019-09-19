class ImgInfo {
    final int iurlgroupid;
    final int iurlid;
    final String iurl;
    final String createdon;
    final String modifiedon; 
    ImgInfo({this.iurlgroupid, this.iurlid, this.iurl, this.createdon, this.modifiedon});
    factory ImgInfo.fromJson(Map<String, dynamic> json) => _imageFromJson(json);
}

ImgInfo _imageFromJson(Map<String, dynamic> json) =>ImgInfo(
    iurlgroupid: json['iurlgroupid'] as int,
    iurlid: json['iurlid'] as int,
    iurl: json['iurl'] as String,
    createdon: json['createdon'] as String,
    modifiedon: json['modifiedon'] as String,
  );

class MetaInfo {
    final int pid;
    final int purlid;
    final int metainfoid;
    final String title;
    final String description;
    final String retailer;
    final String price;
    final String seller;
    final String createdon;
    final String modifiedon;
    MetaInfo({this.pid, this.purlid, this.metainfoid, this.title, this.description, this.retailer, this.price, this.seller, this.createdon, this.modifiedon});
    factory MetaInfo.fromJson(Map<String, dynamic> json) => _metaFromJson(json); 
}

MetaInfo _metaFromJson(Map<String, dynamic> json) => MetaInfo(
  pid: json['pid'] as int,
  purlid: json['purlid'] as int,
  metainfoid: json['metainfoid'] as int,
  title: json['title'] as String,
  description: json['description'] as String,
  retailer: json['retailer'] as String,
  price: json['price'] as String,
  seller: json['seller'] as String,
  createdon: json['createdon'] as String,
  modifiedon: json['modifiedon'] as String,
);

class Product {
  final List<ImgInfo> imageurls;
  final MetaInfo metainfo;
  Product({this.imageurls, this.metainfo});
  factory Product.fromJSON(Map<String, dynamic> json) => _productFromJson(json);
}

Product _productFromJson(Map<String, dynamic> json) {
    var imagesJson = json['imageurls'] as List;
    var metaJson = json['metainfo'];
    List<ImgInfo> images = imagesJson != null ? imagesJson.map((i) => ImgInfo.fromJson(i)).toList() : null;
    MetaInfo meta = metaJson != null ? MetaInfo.fromJson(metaJson) : null;
    return Product(
      imageurls: images,
      metainfo: meta,
    );
}

class Response {
  final int status;
  final String message;
  final Product response;
  Response({this.status, this.message, this.response});
  factory Response.fromJSON(Map<String, dynamic> json) => _responseFromJson(json);
}

Response _responseFromJson(Map<String, dynamic> json) {
  var status = json["status"] as int;
  var message = json["message"] as String;
  var pdt = json["response"];
  Product p = pdt != null ? Product.fromJSON(pdt) : null;
  return Response(
    status: status,
    message: message,
    response: p,
  );
}