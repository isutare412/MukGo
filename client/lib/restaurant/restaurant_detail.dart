import 'package:contra/custom_widgets/custom_app_bar.dart';
import 'package:contra/custom_widgets/custom_header.dart';
import 'package:contra/custom_widgets/custom_search_text.dart';
import 'package:contra/utils/colors.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

import 'restaurant_review_list_item.dart';
import 'restaurant.dart';

class RestaurantDetailPage extends StatefulWidget {
  @override
  _RestaurantDetailPageState createState() => _RestaurantDetailPageState();
}

class _RestaurantDetailPageState extends State<RestaurantDetailPage> {
  TextEditingController _textEditingController = TextEditingController();
  List<RestaurantReview> reviews = List<RestaurantReview>();

  @override
  void initState() {
    super.initState();
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
          lineOneText: "Popular",
          lineTwotext: "Artists",
        ),
      ),
      body: SingleChildScrollView(
        child: Column(
          children: <Widget>[
            Padding(
              padding: const EdgeInsets.only(
                left: 24.0,
                right: 24,
              ),
              child: CustomSearchText(
                iconPath: "assets/icons/ic_search.svg",
                text: "Search with love ...",
                enable: true,
                callback: () {},
                controller: _textEditingController,
              ),
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
