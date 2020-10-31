import 'package:contra/custom_widgets/button_round_with_shadow.dart';
import 'package:contra/custom_widgets/button_solid_with_icon.dart';
import 'package:contra/login/contra_text.dart';
import 'package:contra/utils/colors.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:mukgo/api/api.dart';
import 'package:mukgo/auth/auth_api.dart';
import 'package:mukgo/proto/model.pb.dart';
import 'package:mukgo/restaurant/restaurant_detail_test.dart';
import 'package:mukgo/user/user_model.dart';
import 'package:provider/provider.dart';

import 'input_text_box_bigger.dart';
import 'login_input_email_text.dart';
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
    return Scaffold(
      body: Stack(
        children: [
          Container(
            padding: EdgeInsets.all(24),
            child: Column(
              children: <Widget>[
                Expanded(
                  flex: 1,
                  child: Row(
                    children: <Widget>[
                      ContraText(
                        text: "Create Review",
                        alignment: Alignment.centerLeft,
                      ),
                    ],
                  ),
                ),
                Expanded(
                  flex: 5,
                  child: Column(
                    children: <Widget>[
                      Container(
                        padding: EdgeInsets.only(
                            left: 20, right: 20, top: 10, bottom: 10),
                        margin: EdgeInsets.only(left: 2, right: 2, bottom: 2),
                        decoration: ShapeDecoration(
                            color: athens,
                            shape: RoundedRectangleBorder(
                              borderRadius:
                                  BorderRadius.all(Radius.circular(16)),
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
                      LoginEmailText(
                        //nim people, score
                        text: "menu",
                        iconPath: "assets/icons/ic_search.svg",
                        controller: menuController,
                      ),
                      SizedBox(
                        height: 24,
                      ),
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
                        iconPath: "assets/icons/arrow_next.svg",
                        isPrefix: false,
                        isSuffix: true,
                        text: "Post Review",
                        callback: () async {
                          var auth = getAuth(context);
                          var review = Review()
                            ..comment = commentController.text;
                          review..score = rating;
                          review..menus.add(menuController.text);
                          review..wait = waiting;
                          review..numPeople = numPeople;
                          /*erase after checking whether userName, id are unnecessary
                          var userModel=Provider.of<UserModel>(context, listen: false);
                          var userName=userModel.name;
                          review..userName=userName;
                          review..id='1';
                          */

                          var result = await postReviewData(auth.token,
                              data: review, id: widget.restaurant_id);
                          return Navigator.push(
                              context,
                              MaterialPageRoute(
                                  builder: (context) =>
                                      RestaurantDetailTestPage(
                                        restaurant_id: widget.restaurant_id,
                                      )));
                          //Navigator.of(context).pop();
                          //redirect to restaurant detail test page
                        },
                      ),
                    ],
                  ),
                ),
              ],
            ),
          ),
          Positioned(
            right: 20,
            top: 40,
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
