import 'package:contra/custom_widgets/button_round_with_shadow.dart';
import 'package:contra/login/contra_text.dart';
import 'package:contra/utils/colors.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:intl/intl.dart';
import 'package:mukgo/api/api.dart';
import 'package:mukgo/auth/auth_api.dart';
import 'package:mukgo/proto/model.pb.dart';
import 'package:mukgo/restaurant/restaurant_detail.dart';
import 'package:mukgo/user/user_model.dart';
import 'package:provider/provider.dart';

List<String> profileAsset = [
  'assets/images/onboarding_image_five.svg',
  'assets/images/onboarding_image_one.svg',
  'assets/images/onboarding_image_two.svg',
  'assets/images/onboarding_image_three.svg',
  'assets/images/onboarding_image_four.svg'
];

class ReviewDetailPage extends StatefulWidget {
  ReviewDetailPage({this.review_id, this.restaurant_id});
  final String review_id;
  final String restaurant_id;

  @override
  _ReviewDetailPageState createState() => _ReviewDetailPageState();
}

class _ReviewDetailPageState extends State<ReviewDetailPage> {
  Future<Review> futureReview;
  bool like;
  final DateFormat formatter = DateFormat('MMMd');

  @override
  void initState() {
    super.initState();
    like = true;
    futureReview = Future.microtask(() {
      return fetchReviewData(readAuth(context).token,
          reviewId: widget.review_id);
    });
    Future.microtask(() {
      // fetch user info after randering
      return context.read<UserModel>().fetch();
    });
  }

  void onClickHandler(likedByMe) {
    if (likedByMe) {
      futureReview = Future.microtask(() {
        return deleteLikeData(readAuth(context).token,
            reviewId: widget.review_id);
      });
    } else {
      futureReview = Future.microtask(() {
        return postLikeData(readAuth(context).token,
            reviewId: widget.review_id);
      });
    }
    setState(() {
      like = !like;
    });
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder<Review>(
        future: futureReview,
        builder: (context, snapshot) {
          if (snapshot.hasData) {
            var review = snapshot.data;
            return SingleChildScrollView(
                child: Padding(
              padding: EdgeInsets.all(24),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: <Widget>[
                  Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: <Widget>[
                      Consumer<UserModel>(builder: (context, user, child) {
                        int profileIndex;
                        if (review.userLevel <= 0 || review.userLevel > 4) {
                          profileIndex = 0;
                        } else {
                          profileIndex = review.userLevel;
                        }
                        return SvgPicture.asset(
                          profileAsset[profileIndex],
                          height: 50,
                          width: 50,
                        );
                      }),
                      Expanded(
                          child: Padding(
                              padding: EdgeInsets.only(left: 10.0),
                              child: ContraText(
                                size: 16,
                                text: review.userName,
                                alignment: Alignment.centerLeft,
                              ))),
                      ContraText(
                        size: 16,
                        text: formatter.format(
                            DateTime.fromMillisecondsSinceEpoch(
                                review.timestamp.toInt())),
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
                        fontWeight: FontWeight.normal,
                        fontSize: 24,
                        color: trout),
                  ),
                  Wrap(
                    spacing: 8.0, // gap between adjacent chips
                    runSpacing: 4.0,
                    children: review.menus
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
                        fontWeight: FontWeight.normal,
                        fontSize: 24,
                        color: trout),
                  ),
                  Text(
                    review.comment,
                    style: TextStyle(
                        fontWeight: FontWeight.normal,
                        fontSize: 17,
                        color: trout),
                  ),
                  SizedBox(
                    height: 20,
                  ),
                  Row(
                      mainAxisAlignment: MainAxisAlignment.start,
                      children: <Widget>[
                        Text(
                          'Waiting',
                          textAlign: TextAlign.left,
                          style: TextStyle(
                              fontWeight: FontWeight.normal,
                              fontSize: 24,
                              color: trout),
                        ),
                        Checkbox(
                          value: review.wait,
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
                            Padding(
                                padding: const EdgeInsets.only(
                                    left: 20.0, right: 10.0),
                                child: Icon(
                                  Icons.star,
                                  color: wood_smoke,
                                )),
                            ContraText(
                              text: review.score.toString(),
                              size: 16,
                              alignment: Alignment.center,
                            )
                          ],
                        ),
                        Row(
                          mainAxisAlignment: MainAxisAlignment.center,
                          children: <Widget>[
                            Padding(
                                padding: const EdgeInsets.only(right: 10.0),
                                child: Icon(
                                  Icons.person,
                                  color: wood_smoke,
                                )),
                            ContraText(
                              text: review.numPeople.toString(),
                              size: 16,
                              alignment: Alignment.center,
                            )
                          ],
                        ),
                        Row(
                          mainAxisAlignment: MainAxisAlignment.center,
                          children: <Widget>[
                            IconButton(
                              icon: review.likedByMe
                                  ? Padding(
                                      padding:
                                          const EdgeInsets.only(right: 10.0),
                                      child: Icon(Icons.favorite))
                                  : Padding(
                                      padding:
                                          const EdgeInsets.only(right: 10.0),
                                      child: Icon(Icons.favorite_border)),
                              tooltip: 'Like this review',
                              onPressed: () {
                                onClickHandler(review.likedByMe);
                                //chane the number of likes in the server
                              },
                            ),
                            Padding(
                                padding: const EdgeInsets.only(right: 25.0),
                                child: ContraText(
                                  // number of likes
                                  text: review.likeCount.toString(),
                                  size: 16,
                                  alignment: Alignment.center,
                                ))
                          ],
                        ),
                      ],
                    ),
                  ),
                  Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: <Widget>[
                        Padding(
                          padding: const EdgeInsets.only(top: 20),
                          child: Center(
                            child: ButtonRoundWithShadow(
                              size: 48,
                              iconPath: "assets/icons/arrow_back.svg",
                              borderColor: black,
                              shadowColor: black,
                              color: white,
                              callback: () {
                                Navigator.push(
                                    context,
                                    MaterialPageRoute(
                                        builder: (context) =>
                                            RestaurantDetailTestPage(
                                              restaurant_id:
                                                  widget.restaurant_id,
                                            ))).then((value) {
                                  setState(() {
                                    like = true;
                                  });
                                });
                              },
                            ),
                          ),
                        ),
                        Padding(
                          padding: const EdgeInsets.only(top: 20),
                          child: Consumer<UserModel>(
                              builder: (context, user, child) {
                            if (user.id == review.userId) {
                              return IconButton(
                                  icon: Icon(Icons.delete),
                                  iconSize: 48,
                                  onPressed: () async {
                                    var result = await deleteReviewData(
                                        getAuth(context).token,
                                        reviewId: widget.review_id);
                                    if (result) {
                                      Navigator.pushReplacement(
                                          context,
                                          MaterialPageRoute(
                                              builder: (context) =>
                                                  RestaurantDetailTestPage(
                                                    restaurant_id:
                                                        widget.restaurant_id,
                                                  ))).then((value) {
                                        setState(() {
                                          like = true;
                                        });
                                      });
                                    }
                                  });
                            } else {
                              return Icon(
                                Icons.delete,
                                color: white,
                              );
                            }
                          }),
                        )
                      ]),
                ],
              ),
            ));
          }
          return Center(
            child: CircularProgressIndicator(),
          );
        });
  }
}
