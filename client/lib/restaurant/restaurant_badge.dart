import 'package:flutter/material.dart';
import 'package:flutter_svg/svg.dart';
import 'package:mukgo/proto/model.pb.dart';

Map<RestaurantType, String> badgePathMap = {
  RestaurantType.INVALID: 'assets/icons/r_invalid.svg',
  RestaurantType.CHICKEN: 'assets/icons/r_chicken.svg',
  RestaurantType.CAFE: 'assets/icons/r_cafe.svg',
  RestaurantType.FASTFOOD: 'assets/icons/r_fastfood.svg',
  RestaurantType.MEAT: 'assets/icons/r_meat.svg',
  RestaurantType.DESSERT: 'assets/icons/r_dessert.svg',
  RestaurantType.JAPANESE: 'assets/icons/r_japanese.svg',
  RestaurantType.KOREAN: 'assets/icons/r_korean.svg',
  RestaurantType.CHINESE: 'assets/icons/r_chinese.svg',
};

class RestaurantBadge extends StatelessWidget {
  final RestaurantType restaurantType;

  RestaurantBadge({this.restaurantType});

  @override
  Widget build(BuildContext context) {
    var path =
        badgePathMap[restaurantType] ?? badgePathMap[RestaurantType.INVALID];

    return Padding(
      padding: const EdgeInsets.all(8.0),
      child: SvgPicture.asset(path),
    );
  }
}
