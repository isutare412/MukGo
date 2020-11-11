import 'package:contra/custom_widgets/button_plain_with_shadow.dart';
import 'package:contra/custom_widgets/button_round_with_shadow.dart';
import 'package:contra/custom_widgets/custom_app_bar.dart';
import 'package:contra/login/contra_text.dart';
import 'package:contra/utils/colors.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:mukgo/review/review_card_data.dart';
import 'package:mukgo/user/user_model.dart';
import 'package:provider/provider.dart';

class ReviewDetailPage extends StatefulWidget {
  ReviewDetailPage({this.review_data});
  final ReviewCardData review_data;

  @override
  _ReviewDetailPageState createState() => _ReviewDetailPageState();
}

class _ReviewDetailPageState extends State<ReviewDetailPage> {
  bool like;

  @override
  void initState() {
    super.initState();
    like = false;
  }

  @override
  Widget build(BuildContext context) {
    double statusBarHeight = MediaQuery.of(context).padding.top;
    return SingleChildScrollView(
      // padding: EdgeInsets.all(24),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: <Widget>[
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: <Widget>[
              Consumer<UserModel>(builder: (context, user, child) {
                return SvgPicture.asset(
                  user.profileAsset(),
                  height: 48,
                  width: 48,
                );
              }),
              Expanded(
                  child: Padding(
                      padding: EdgeInsets.only(left: 10.0),
                      child: ContraText(
                        size: 16,
                        text: widget.review_data.user,
                        alignment: Alignment.centerLeft,
                      ))),
              ContraText(
                size: 16,
                text: widget.review_data.time,
                alignment: Alignment.centerRight,
              )
            ],
          ),
          SizedBox(
            height: 20,
          ),
          Text(
            'Menus',
            textAlign: TextAlign.left,
            style: TextStyle(
                fontWeight: FontWeight.normal, fontSize: 24, color: trout),
          ),
          Wrap(
            spacing: 8.0, // gap between adjacent chips
            runSpacing: 4.0,
            children: widget.review_data.menus
                .map((menu) => Padding(
                    padding: EdgeInsets.only(right: 6),
                    child: Text(
                      '#$menu',
                      style: TextStyle(
                          fontSize: 16,
                          fontWeight: FontWeight.bold,
                          color: trout),
                    )))
                .toList(),
          ),
          SizedBox(
            height: 10,
          ),
          Text(
            'Comment',
            textAlign: TextAlign.left,
            style: TextStyle(
                fontWeight: FontWeight.normal, fontSize: 24, color: trout),
          ),
          Text(
            widget.review_data.comment,
            style: TextStyle(
                fontWeight: FontWeight.normal, fontSize: 17, color: trout),
          ),
          SizedBox(
            height: 20,
          ),
          Row(mainAxisAlignment: MainAxisAlignment.start, children: <Widget>[
            Text(
              'Waiting',
              textAlign: TextAlign.left,
              style: TextStyle(
                  fontWeight: FontWeight.normal, fontSize: 24, color: trout),
            ),
            Checkbox(
              value: widget.review_data.waiting,
            ),
          ]),
          SizedBox(
            height: 20,
          ),
          Container(
            height: 40,
            decoration: ShapeDecoration(
              color: Colors.grey[300],
              shape: RoundedRectangleBorder(
                  borderRadius: BorderRadius.all(Radius.circular(16))),
            ),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: <Widget>[
                Row(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: <Widget>[
                    Icon(
                      Icons.star,
                      color: wood_smoke,
                    ),
                    ContraText(
                      text: widget.review_data.score.toString(),
                      size: 13,
                      alignment: Alignment.center,
                    )
                  ],
                ),
                Row(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: <Widget>[
                    Icon(
                      Icons.person,
                      color: wood_smoke,
                    ),
                    ContraText(
                      text: widget.review_data.numPeople.toString(),
                      size: 13,
                      alignment: Alignment.center,
                    )
                  ],
                ),
                Row(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: <Widget>[
                    IconButton(
                      icon: like
                          ? Icon(Icons.favorite)
                          : Icon(Icons.favorite_border),
                      tooltip: 'Like this review',
                      onPressed: () {
                        setState(() {
                          like = !like;
                          //chane the number of likes in the server
                        });
                      },
                    ),
                    ContraText(
                      text: "no data yet",
                      size: 13,
                      alignment: Alignment.center,
                    )
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
