import 'package:flutter/material.dart';
import 'dart:async';
import 'product_repository.dart';
import 'product.dart';

void main() => runApp(ProductListApp());

class ImageTile extends StatelessWidget {
  final ImgInfo _img;
  final MetaInfo _meta;
  ImageTile(this._img, this._meta);

  @override
  Widget build(BuildContext context) => Column(
    children: <Widget>[
      ListTile(
        title: Text(_meta.title),
        subtitle: Text(_meta.price + " (" + _meta.retailer + ")"),
        leading: Container(
          margin: EdgeInsets.only(left: 6.0),
          child: Image.network(_img.iurl, height: 50.0, fit: BoxFit.fill,)
        ),
      ),
      Divider()
    ],
  );
}

class ProductTile extends StatelessWidget {
  final Product _product;
  ProductTile(this._product);

  @override
  Widget build(BuildContext context) {
    if (_product.imageurls.length != 0) {
      return new Column(
              children: <Widget>[
                      SizedBox(
                          width: 500.0,
                          height: 70.0,
                          child: ListView.builder(
                          itemCount: _product.imageurls.length,
                          itemBuilder: (context, index) => ImageTile(_product.imageurls[index], _product.metainfo),
                      ),
              ),
              Divider()
            ],
      );
    }
    return new Container();
  } 
}

class ProductListApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) => MaterialApp(
    title: 'Product List App',
    debugShowCheckedModeBanner: false,
    theme: ThemeData(
      primaryColor: Colors.black,
      accentColor: Colors.black
    ),
    home: Home(),
  );
}


class Home extends StatefulWidget {
  @override
  _HomeState createState() => _HomeState();
}

class _HomeState extends State<Home> {
  List<Product> _products = <Product>[];

  @override
  void initState() {
    super.initState();
    const interval = const Duration(seconds:10);
    new Timer.periodic(interval, (Timer t) => listenForProducts());
    listenForProducts();
  }

  void listenForProducts() async {
    _products = <Product>[];
    final Stream<Product> stream = await getInfo();
    stream.listen((Product p) =>
      setState(() =>  _products.add(p))
    );
  }

  Widget build(BuildContext context) => Scaffold(
    appBar: AppBar(
      centerTitle: true,
      title: Text('Scraped Products'),
    ),
    body: ListView.builder(
      itemCount: _products.length,
      itemBuilder: (context, index) => ProductTile(_products[index]),
    ),
  );
}