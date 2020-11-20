import 'package:flutter/material.dart';
import 'package:contra/utils/colors.dart';
import 'package:mukgo/proto/model.pb.dart';
import 'package:mukgo/restaurant/restaurant_badge.dart';

class UserBadge {
  final RestaurantType type;
  final int count;

  UserBadge({this.type, this.count});
}

class BadgeGrid extends StatelessWidget {
  List<UserBadge> badges;

  BadgeGrid({this.badges});

  @override
  Widget build(BuildContext context) {
    if (badges.isEmpty) {
      badges = testBadges();
    }

    return GridView.builder(
        padding: EdgeInsets.all(24),
        gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
          crossAxisCount: 4,
          crossAxisSpacing: 16,
          mainAxisSpacing: 16,
          childAspectRatio: (1 / 1),
        ),
        controller: ScrollController(keepScrollOffset: false),
        shrinkWrap: true,
        scrollDirection: Axis.vertical,
        itemCount: badges.length,
        itemBuilder: (BuildContext context, int index) {
          var type = badges[index].type;
          var count = badges[index].count;

          return Container(
              child: RestaurantBadge(
            restaurantType: type,
          ));
        });
  }
}

List<UserBadge> testBadges() {
  return List.generate(12, (index) {
    return UserBadge(
        type: RestaurantType.values[index % RestaurantType.values.length],
        count: index);
  });
}
