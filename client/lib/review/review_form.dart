import 'package:contra/custom_widgets/button_round_with_shadow.dart';
import 'package:contra/custom_widgets/button_solid_with_icon.dart';
import 'package:contra/login/contra_text.dart';
import 'package:contra/utils/colors.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:fixnum/fixnum.dart';
import 'package:mukgo/api/api.dart';
import 'package:mukgo/auth/auth_api.dart';
import 'package:mukgo/proto/model.pb.dart';
import 'package:mukgo/restaurant/restaurant_detail_test.dart';

import 'input_text_box_bigger.dart';
import 'dart:developer' as developer;

class ReviewForm extends StatefulWidget {
  ReviewForm({this.restaurant_id});
  final String restaurant_id;

  @override
  _ReviewForm createState() => _ReviewForm();
}

class _ReviewForm extends State<ReviewForm> {
  final menuController = TextEditingController();
  final commentController = TextEditingController();

  int numPeople = 1;
  int rating = 3;
  bool waiting = false;
  List<String> _menu = List<String>();

  @override
  void initState() {
    super.initState();
  }

  @override
  void dispose() {
    // Clean up the controller when the widget is removed from the widget tree.
    // This also removes the _printLatestValue listener.
    menuController.dispose();
    commentController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return SingleChildScrollView(
      child: Column(
        children: [
          Container(
            padding: EdgeInsets.symmetric(vertical: 0, horizontal: 20.0),
            child: Column(
              children: <Widget>[
                Padding(
                  padding: const EdgeInsets.only(top: 20, bottom: 10),
                  child: Row(
                    children: <Widget>[
                      ContraText(
                        size: 30.0,
                        text: 'Create Review',
                        alignment: Alignment.centerLeft,
                      ),
                    ],
                  ),
                ),
                Column(
                  children: <Widget>[
                    Container(
                      padding: EdgeInsets.only(
                          left: 20, right: 20, top: 10, bottom: 10),
                      margin: EdgeInsets.only(left: 2, right: 2, bottom: 2),
                      decoration: ShapeDecoration(
                          color: athens,
                          shape: RoundedRectangleBorder(
                            borderRadius: BorderRadius.all(Radius.circular(16)),
                          )),
                      child: Column(
                        children: <Widget>[
                          Row(
                            mainAxisAlignment: MainAxisAlignment.spaceBetween,
                            children: [
                              Text(
                                'Number of People :  ',
                                style: TextStyle(fontSize: 18),
                              ),
                              DropdownButton<int>(
                                value: numPeople,
                                icon: Icon(Icons.arrow_downward),
                                iconSize: 24,
                                elevation: 16,
                                style: TextStyle(color: Colors.black),
                                underline: Container(
                                  height: 2,
                                  color: Colors.deepPurpleAccent,
                                ),
                                onChanged: (int newValue) {
                                  setState(() {
                                    numPeople = newValue;
                                  });
                                },
                                items: <int>[1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
                                    .map<DropdownMenuItem<int>>((int value) {
                                  return DropdownMenuItem<int>(
                                    value: value,
                                    child: Text(value.toString()),
                                  );
                                }).toList(),
                              ),
                            ],
                          ),
                          Row(
                            mainAxisAlignment: MainAxisAlignment.spaceBetween,
                            children: [
                              Text(
                                'Rating :  ',
                                style: TextStyle(fontSize: 18),
                              ),
                              DropdownButton<int>(
                                value: rating,
                                icon: Icon(Icons.arrow_downward),
                                iconSize: 24,
                                elevation: 16,
                                style: TextStyle(color: Colors.black),
                                underline: Container(
                                  height: 2,
                                  color: Colors.deepOrangeAccent,
                                ),
                                onChanged: (int newValue) {
                                  setState(() {
                                    rating = newValue;
                                  });
                                },
                                items: <int>[1, 2, 3, 4, 5]
                                    .map<DropdownMenuItem<int>>((int value) {
                                  return DropdownMenuItem<int>(
                                    value: value,
                                    child: Text(value.toString()),
                                  );
                                }).toList(),
                              ),
                            ],
                          ),
                          Row(
                            mainAxisAlignment: MainAxisAlignment.spaceBetween,
                            children: [
                              Text(
                                'Waiting :  ',
                                style: TextStyle(fontSize: 18),
                              ),
                              Checkbox(
                                value: waiting,
                                onChanged: (bool value) {
                                  setState(() {
                                    waiting = value;
                                  });
                                },
                              ),
                            ],
                          ),
                        ],
                      ),
                    ),
                    SizedBox(
                      height: 24,
                    ),
                    Column(children: <Widget>[
                      Row(children: <Widget>[
                        Text(
                          'Menu :  ',
                          style: TextStyle(fontSize: 18),
                        ),
                        Container(
                            width: 180,
                            child: TextField(
                                controller: menuController,
                                decoration: InputDecoration(
                                  hintText: "Enter your menu",
                                  suffixIcon: IconButton(
                                    onPressed: () => menuController.clear(),
                                    icon: Icon(Icons.clear),
                                  ),
                                ))),
                        Container(
                          child: RaisedButton(
                            onPressed: () {
                              setState(() {
                                if (menuController.text.isNotEmpty) {
                                  if (!_menu.contains(menuController.text)) {
                                    _menu.add(menuController.text);
                                  }
                                }

                                menuController.clear();
                              });
                            },
                            child: Text('Submit'),
                          ),
                        ),
                      ]),
                      SizedBox(
                        height: 10,
                      ),
                      Wrap(
                        spacing: 2.0,
                        children: _menu
                            .map((menu) => Row(
                                    mainAxisAlignment: MainAxisAlignment.start,
                                    children: <Widget>[
                                      SizedBox(
                                        width: 30,
                                      ),
                                      Flexible(
                                          child: Text(menu,
                                              style: TextStyle(
                                                backgroundColor:
                                                    Colors.grey[300],
                                              ),
                                              overflow: TextOverflow.ellipsis,
                                              maxLines: 2,
                                              softWrap: true)),
                                      IconButton(
                                        icon: Icon(Icons.clear),
                                        iconSize: 20.0,
                                        color: Colors.blueGrey,
                                        onPressed: () {
                                          setState(() {
                                            _menu.remove(menu);
                                          });
                                        },
                                      ),
                                    ]))
                            .toList(),
                      ),
                    ]),
                    InputTextBoxBigger(
                      text: "comment",
                      iconPath: "assets/icons/mail.svg",
                      controller: commentController,
                    ),
                    SizedBox(
                      height: 24,
                    ),
                    ButtonPlainWithIcon(
                      color: wood_smoke,
                      textColor: white,
                      iconPath: 'assets/icons/arrow_next.svg',
                      isPrefix: false,
                      isSuffix: true,
                      text: 'Post Review',
                      callback: () async {
                        var auth = getAuth(context);
                        var review = Review()..comment = commentController.text;
                        review.score = rating;
                        for (int i = 0; i < _menu.length; i++) {
                          review.menus.add(_menu[i]);
                        }
                        review.wait = waiting;
                        review.numPeople = numPeople;
                        review.timestamp =
                            Int64(DateTime.now().millisecondsSinceEpoch);
                        var result = await postReviewData(auth.token,
                            data: review, id: widget.restaurant_id);
                        return Navigator.push(
                            context,
                            MaterialPageRoute(
                                builder: (context) => RestaurantDetailTestPage(
                                      restaurant_id: widget.restaurant_id,
                                    )));
                        //Navigator.of(context).pop();
                        //redirect to restaurant detail test page
                      },
                    ),
                  ],
                ),
              ],
            ),
          ),
          Padding(
            padding: const EdgeInsets.only(top: 20),
            child: ButtonRoundWithShadow(
              size: 48,
              iconPath: "assets/icons/close.svg",
              borderColor: black,
              shadowColor: black,
              color: white,
              callback: () {
                Navigator.of(context).pop();
              },
            ),
          )
        ],
      ),
    );
  }
}
