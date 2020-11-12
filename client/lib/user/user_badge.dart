import 'package:flutter/material.dart';
import 'package:contra/utils/colors.dart';
import 'package:mukgo/proto/model.pb.dart';

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
              color: dandelion,
              child: Padding(
                  padding: EdgeInsets.all(6),
                  child: Center(
                      child: Column(
                          mainAxisSize: MainAxisSize.min,
                          children: <Widget>[
                        Text(
                          '$type',
                          style: TextStyle(
                            fontSize: 14,
                            fontWeight: FontWeight.bold,
                            color: wood_smoke,
                          ),
                        ),
                        Text(
                          '$count',
                          style: TextStyle(
                            fontSize: 20,
                            fontWeight: FontWeight.bold,
                            color: wood_smoke,
                          ),
                        )
                      ]))));
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
