import 'package:contra/custom_widgets/custom_app_bar.dart';
import 'package:contra/login/contra_text.dart';
import 'package:contra/utils/colors.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

import 'package:mukgo/review/review_card.dart';
import 'package:mukgo/proto/model.pb.dart';
import 'package:mukgo/api/api.dart';
import 'package:mukgo/auth/auth_api.dart';

List<Color> colors = [dandelion, foam, mona_lisa, fair_pink, white];

class ReviewPageArguments {
  final String id;
  final String name;

  ReviewPageArguments({this.id, this.name});
}

class ReviewList extends StatefulWidget {
  @override
  _ReviewCardListState createState() => _ReviewCardListState();
}

class _ReviewCardListState extends State<ReviewList> {
  Future<Reviews> futureReviews;

  @override
  void initState() {
    super.initState();

    futureReviews = Future.microtask(() {
      ReviewPageArguments args = ModalRoute.of(context).settings.arguments;
      return fetchReviewsData(readAuth(context).token, restaurantId: args.id);
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: CustomAppBar(
        height: 150,
        child: Column(
          mainAxisAlignment: MainAxisAlignment.end,
          children: [
            Padding(
              padding: const EdgeInsets.only(top: 24.0, left: 24),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.start,
                children: <Widget>[
                  ContraText(
                    size: 44,
                    alignment: Alignment.bottomCenter,
                    text: "Reviews",
                  ),
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
                    title: review.comment,
                    bgColor: color);
              }).toList();

              return ListView.builder(
                  itemCount: data.length,
                  padding: EdgeInsets.only(left: 24, right: 24, bottom: 24),
                  itemBuilder: (BuildContext context, int index) {
                    return ReviewCard(
                      data: data[index],
                      isSubType: index == 4 ? true : false,
                      onTap: () {
                        Navigator.pushNamed(context, "/blog_detail_page");
                      },
                    );
                  });
            }

            return Center(
              child: CircularProgressIndicator(),
            );
          }),
    );
  }
}
