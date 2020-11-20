import 'dart:async';

import 'package:modal_bottom_sheet/modal_bottom_sheet.dart';
import 'package:intl/intl.dart';
import 'package:flutter/material.dart';
import 'package:contra/utils/colors.dart';
import 'package:contra/custom_widgets/button_solid_with_icon.dart';
import 'package:mukgo/restaurant/shopping_card_pager_item.dart';
import 'package:mukgo/review/review_card_data.dart';

import 'package:mukgo/review/review_detail_test.dart';
import 'package:mukgo/review/review_form.dart';
import 'package:mukgo/api/api.dart';
import 'package:mukgo/proto/model.pb.dart';
import 'package:mukgo/auth/auth_api.dart';
import 'package:mukgo/review/review_card_proj.dart';
import 'package:mukgo/restaurant/review_filter.dart';
import 'restaurant.dart';

class RestaurantDetailTestPage extends StatefulWidget {
  RestaurantDetailTestPage({this.restaurant_id});
  final String restaurant_id;

  @override
  _RestaurantDetailTestPageState createState() =>
      _RestaurantDetailTestPageState();
}

class _RestaurantDetailTestPageState extends State<RestaurantDetailTestPage> {
  List<RestaurantReview> reviews = List<RestaurantReview>();

  Future<Reviews> futureReviews;
  Future<Restaurant> futureRestaurant;
  final DateFormat formatter = DateFormat('MMMd');
  Filter filter;

  void setFilter(Filter newFilter) {
    setState(() {
      filter = newFilter;
    });
  }

  void orderReviews(List<Review> reviews) {
    var multiplyer = filter.order.ascending ? 1 : -1;
    if (filter.order.key == 'score') {
      reviews.sort((a, b) => multiplyer * a.score.compareTo(b.score));
    } else {
      reviews.sort((a, b) => multiplyer * a.timestamp.compareTo(b.timestamp));
    }
  }

  List<Review> filterReviews(List<Review> reviews) {
    return reviews.where((review) {
      var checkPeople = review.numPeople <= filter.numPeople;
      bool checkLine;
      if (filter.wait == 'true') {
        checkLine = review.wait;
      } else if (filter.wait == 'false') {
        checkLine = !review.wait;
      } else {
        checkLine = true;
      }

      return checkPeople && checkLine;
    }).toList();
  }

  @override
  void initState() {
    super.initState();
    filter = Filter(
        numPeople: 10,
        wait: 'disable',
        order: Order(key: 'time', ascending: false));
    //GET restaurant info
    futureRestaurant = Future.microtask(() {
      return fetchRestaurantData(readAuth(context).token,
          restaurantId: widget.restaurant_id);
    });
    //GET review data
    futureReviews = Future.microtask(() {
      return fetchReviewsData(readAuth(context).token,
          restaurantId: widget.restaurant_id);
    });
  }

  @override
  Widget build(BuildContext context) {
    print(filter.order.key);
    return Stack(children: <Widget>[
      SingleChildScrollView(
          child: Column(children: <Widget>[
        FutureBuilder<Restaurant>(
            future: futureRestaurant,
            builder: (context, snapshot) {
              if (snapshot.hasData) {
                var restaurantData = snapshot.data;
                return ShoppingCardPagerItem(
                    restaurantName: restaurantData.name,
                    restaurantType: restaurantData.type);
              }
              return Center(
                child: CircularProgressIndicator(),
              );
            }),
        FilterOpenWidget(
          onTap: () {
            showMaterialModalBottomSheet(
              context: context,
              builder: (context, scrollController) =>
                  FilterModal(preFilter: filter, callback: setFilter),
            );
          },
        ),
        FutureBuilder<Reviews>(
            future: futureReviews,
            builder: (context, snapshot) {
              if (snapshot.hasData) {
                var reviews = snapshot.data.reviews;
                reviews = filterReviews(reviews);
                orderReviews(reviews);
                var data = reviews.asMap().entries.map((entry) {
                  var i = entry.key;
                  var review = entry.value;
                  var time = formatter.format(
                      DateTime.fromMillisecondsSinceEpoch(
                          review.timestamp.toInt()));
                  var menus = review.menus;
                  return ReviewCardData(
                      reviewId: review.reviewId,
                      user: review.userName,
                      comment: review.comment,
                      score: review.score,
                      numPeople: review.numPeople,
                      time: time,
                      menus: menus,
                      waiting: review.wait,
                      userLevel: review.userLevel,
                      likeCount: review.likeCount,
                      likeByMe: review.likedByMe);
                }).toList();

                return ListView.builder(
                  shrinkWrap: true,
                  physics: NeverScrollableScrollPhysics(),
                  itemCount: data.length,
                  itemBuilder: (BuildContext context, int index) {
                    return ReviewCard(
                        reviewData: data[index],
                        onTap: () {
                          Navigator.push(
                              context,
                              MaterialPageRoute(
                                  builder: (context) => ReviewDetailPage(
                                      review_id: data[index].reviewId,
                                      restaurant_id: widget.restaurant_id)));
                        });
                  },
                );
              }
              return Center(child: CircularProgressIndicator());
            }),
        SizedBox(height: 60),
      ])),
      Align(
        alignment: Alignment.bottomCenter,
        child: Padding(
            padding: const EdgeInsets.symmetric(horizontal: 24.0, vertical: 12),
            child: ButtonPlainWithIcon(
              color: wood_smoke,
              textColor: white,
              iconPath: 'assets/icons/arrow_next.svg',
              isPrefix: false,
              isSuffix: true,
              text: 'Write Review',
              callback: () {
                Navigator.push(
                    context,
                    MaterialPageRoute(
                        builder: (context) =>
                            ReviewForm(restaurant_id: widget.restaurant_id)));
              },
            )),
      )
    ]);
  }
}
