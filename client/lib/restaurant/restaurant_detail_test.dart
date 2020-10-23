import 'dart:convert';

import 'package:contra/custom_widgets/custom_app_bar.dart';
import 'package:contra/custom_widgets/custom_header.dart';
import 'package:contra/custom_widgets/custom_search_text.dart';
import 'package:contra/utils/colors.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

import 'package:mukgo/review/review_form.dart';
import 'package:mukgo/api/api.dart';
import 'package:mukgo/proto/model.pb.dart';
import 'dart:async';
import 'package:mukgo/auth/auth_api.dart';
import 'package:mukgo/auth/auth_model.dart';

import 'restaurant_review_list_item.dart';
import 'restaurant.dart';

class RestaurantDetailTestPage extends StatefulWidget {
  RestaurantDetailTestPage({this.restaurant_id});
  final String restaurant_id;

  @override
  _RestaurantDetailTestPageState createState() => _RestaurantDetailTestPageState();
}

class _RestaurantDetailTestPageState extends State<RestaurantDetailTestPage> {
  TextEditingController _textEditingController = TextEditingController();
  List<RestaurantReview> reviews = List<RestaurantReview>();

  //restaurant data
  String name;
  double latitude;
  double longitude;

  @override
  void initState() {
    super.initState();
    /*
    var auth=readAuth(context);
    var tok= auth.token;
    //GET restaurant info (phone number, name, rating etc)
    var restaurantData= fetchRestaurantsData(tok, id: widget.restaurant_id);
    //GET review data
    var reviewsData = fetchReviewsData(tok, id: widget.restaurant_id);
    */
    //var restaurantData= getDummyRestaurant(widget.restaurant_id);
    //var reviewsData= getDummyReviews(widget.restaurant_id);
    //extract reviews from reviewsData and set them to reviews
    /*
    setState(() {
      name: restaurantData.name;

    });
    */  
    
    reviews.add(RestaurantReview(
        name: "Angela Mehra",
        designation: "Designer",
        profile: "assets/images/peep_man_three.svg",
        textColor: white,
        bgColor: carribean_green));
    reviews.add(RestaurantReview(
        name: "Konami Ravi",
        designation: "Muscian",
        textColor: white,
        profile: "assets/images/peep_lady_one.svg",
        bgColor: flamingo));
    reviews.add(RestaurantReview(
        name: "Hard Cops",
        textColor: white,
        designation: "Bill Rizer",
        profile: "assets/images/peep_man_right.svg",
        bgColor: Colors.yellow));
    reviews.add(RestaurantReview(
      textColor: black,
      name: "Kalia Youknow",
      designation: "Muscian",
      profile: "assets/images/peep_lady_right.svg",
    ));
    reviews.add(RestaurantReview(
      textColor: white,
      name: " Genbei Yagy ",
      designation: "Planet Designer",
      bgColor: caribbean_color,
      profile: "assets/images/peep_lady_right.svg",
    ));
    
  }
  
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: white,
      appBar: CustomAppBar(
        height: 200,
        child: CustomHeader(
          fg_color: wood_smoke,
          bg_color: white,
          color: wood_smoke,
          lineOneText: "name",
          //should be set to info extacted from Restaurant
          lineTwotext: "other info",
          //should be set to info extacted from Restaurant
        ),
      ),
      body: SingleChildScrollView(
        child: Column(
          children: <Widget>[
            Padding(
              padding: EdgeInsets.only(right: 20.0),
              child: GestureDetector(
                onTap: () {
                  Navigator.push(
                    context,
                    MaterialPageRoute(
                      builder: (context) => ReviewForm(
                        restaurant_id: widget.restaurant_id,
                    ))
                  );
                },
                child: Text(
                  'Write Review',
                  textAlign: TextAlign.center,
                  style: TextStyle(
                      fontSize: 17, color: trout, fontWeight: FontWeight.w500),
                  ),
              )
            ),
            
            ListView.builder(
                shrinkWrap: true,
                padding: EdgeInsets.all(24),
                physics: NeverScrollableScrollPhysics(),
                itemCount: reviews.length,
                itemBuilder: (BuildContext context, int index) {
                  return RestaurantReviewListItem(
                    review: reviews[index],
                  );
                })
                
          ],
        ),
      ),
    );
  }
}
