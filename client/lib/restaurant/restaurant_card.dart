import 'package:contra/utils/colors.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:mukgo/proto/model.pb.dart';
import 'package:mukgo/restaurant/restaurant_badge.dart';

class RestaurantCard extends StatelessWidget {
  final String restaurantName;
  final RestaurantType restaurantType;

  const RestaurantCard({this.restaurantName, this.restaurantType});

  @override
  Widget build(BuildContext context) {
    return Padding(
        padding: const EdgeInsets.only(left: 12, right: 12, top: 12),
        child: Container(
          height: 120.0,
          margin: EdgeInsets.symmetric(horizontal: 8),
          decoration: ShapeDecoration(
              color: Colors.yellow,
              shape: RoundedRectangleBorder(
                  borderRadius: BorderRadius.all(Radius.circular(16)),
                  side: BorderSide(color: wood_smoke, width: 2))),
          child: Row(
            children: <Widget>[
              Column(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                crossAxisAlignment: CrossAxisAlignment.start,
                children: <Widget>[
                  Container(
                    width: MediaQuery.of(context).size.width * 0.65,
                    child: Padding(
                      padding: const EdgeInsets.only(left: 24.0, top: 16),
                      child: Text(
                        restaurantName,
                        maxLines: 2,
                        overflow: TextOverflow.ellipsis,
                        style: TextStyle(
                            color: wood_smoke,
                            fontSize: 24,
                            fontWeight: FontWeight.w800),
                      ),
                    ),
                  ),
                ],
              ),
              Expanded(child: RestaurantBadge(restaurantType: restaurantType))
            ],
          ),
        ));
  }
}
