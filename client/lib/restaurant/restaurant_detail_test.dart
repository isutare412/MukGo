import 'dart:convert';

import 'package:contra/custom_widgets/custom_app_bar.dart';
import 'package:contra/custom_widgets/custom_header.dart';
import 'package:contra/custom_widgets/custom_search_text.dart';
import 'package:contra/login/contra_text.dart';
import 'package:contra/utils/colors.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:mukgo/project/my_home_page.dart';

import 'package:mukgo/review/review_form.dart';
import 'package:mukgo/api/api.dart';
import 'package:mukgo/proto/model.pb.dart';
import 'dart:async';
import 'package:mukgo/auth/auth_api.dart';
import 'package:mukgo/auth/auth_model.dart';
import 'package:mukgo/review/review_card_proj.dart';
import 'restaurant.dart';

/*
class ReviewPageArguments {
  final String id;
  final String name;

  ReviewPageArguments({this.id, this.name});
}
*/
List<Color> colors = [dandelion, foam, mona_lisa, fair_pink, white];

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

  @override
  void initState() {
    super.initState();
    //GET restaurant info
    futureRestaurant = Future.microtask(() {
      /*
      var restaurantId = ModalRoute.of(context).settings.arguments.restaurantId;
      return fetchRestaurantData(tok, restaurantId: restaurantId);
      */
      return fetchRestaurantData(readAuth(context).token,
          restaurantId: widget.restaurant_id);
    });
    //GET review data
    futureReviews = Future.microtask(() {
      /*
      var restaurantId = ModalRoute.of(context).settings.arguments.restaurantId;
      return fetchReviewsData(tok, restaurantId: restaurantId);
      */
      return fetchReviewsData(readAuth(context).token,
          restaurantId: widget.restaurant_id);
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: white,
      appBar: CustomAppBar(
        height: 150,
        child: Column(
          mainAxisAlignment: MainAxisAlignment.end,
          children: [
            Padding(
              padding: const EdgeInsets.all(24.0),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                crossAxisAlignment: CrossAxisAlignment.end,
                children: <Widget>[
                  IconButton(
                    icon: Icon(Icons.arrow_back),
                    color: wood_smoke,
                    onPressed: () {
                      Navigator.push(
                          context,
                          MaterialPageRoute(
                              builder: (context) => MyHomePage(
                                    title: 'Mukgo Project',
                                  )));
                    },
                  ),
                  Expanded(
                    flex: 1,
                    child: ContraText(
                      size: 27,
                      alignment: Alignment.bottomCenter,
                      text: "contra",
                    ),
                  ),
                  IconButton(
                    color: wood_smoke,
                    icon: Icon(Icons.edit),
                    onPressed: () {
                      Navigator.push(
                          context,
                          MaterialPageRoute(
                              builder: (context) => ReviewForm(
                                    restaurant_id: widget.restaurant_id,
                                  )));
                    },
                  )
                ],
              ),
            ),
          ],
        ),
      ),
      body: FutureBuilder<Reviews>(
          future: futureReviews,
          builder: (context, snapshot) {
            if (snapshot.hasData) {
              var reviews = snapshot.data.reviews;
              var data = reviews.asMap().entries.map((entry) {
                var i = entry.key;
                var review = entry.value;
                var color = colors[i % colors.length];
                return ReviewCardData(
                    user: review.userName,
                    comment: review.comment,
                    score: review.score,
                    like: 4,
                    time: 'june 11',
                    bgColor: color);
              }).toList();

              return ListView.builder(
                  itemCount: data.length,
                  itemBuilder: (BuildContext context, int index) {
                    return ReviewCard(reviewData: data[index], onTap: () {});
                  });
            }

            return Center(
              child: CircularProgressIndicator(),
            );
          }),
    );
  }
}

/*
List<ReviewCardData> dummyData;
            dummyData.add(ReviewCardData(
              user: 'Gina',
              comment: 'MENU1',
              score: 4,
              like: 4,
              time: 'june 11',
              bgColor: colors[0]
            ));
            dummyData.add(ReviewCardData(
              user: 'Gina2',
              comment: 'MENU2',
              score: 4,
              like: 4,
              time: 'june 11',
              bgColor: colors[0]
            ));
            dummyData.add(ReviewCardData(
              user: 'Gina3',
              comment: 'MENU3',
              score: 4,
              like: 4,
              time: 'june 11',
              bgColor: colors[2]
            ));
            return ListView.builder(
                itemCount: dummyData.length,
                itemBuilder: (BuildContext context, int index) {
                  return ReviewListItem(
                    reviewcard: dummyData[index],
                    onTap: () {}
                  );
                }
            );
*/
