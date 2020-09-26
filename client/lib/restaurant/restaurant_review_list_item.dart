import 'package:contra/custom_widgets/contra_button.dart';
import 'package:contra/login/contra_text.dart';
import 'package:contra/utils/colors.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/svg.dart';

import 'restaurant.dart';

class RestaurantReviewListItem extends StatelessWidget {
  final RestaurantReview review;

  const RestaurantReviewListItem({this.review});

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: EdgeInsets.symmetric(vertical: 12),
      decoration: ShapeDecoration(
          color: review.bgColor,
          shape: RoundedRectangleBorder(
              borderRadius: BorderRadius.all(Radius.circular(16)),
              side: BorderSide(color: wood_smoke, width: 2))),
      child: Stack(
        children: [
          Padding(
            padding: const EdgeInsets.all(24.0),
            child: Column(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              crossAxisAlignment: CrossAxisAlignment.start,
              mainAxisSize: MainAxisSize.max,
              children: <Widget>[
                ContraText(
                  text: review.name,
                  size: 24,
                  color: wood_smoke,
                  weight: FontWeight.w800,
                  alignment: Alignment.centerLeft,
                  textAlign: TextAlign.left,
                ),
                SizedBox(
                  height: 12,
                ),
                ContraText(
                  text: review.designation,
                  size: 17,
                  color: white,
                  weight: FontWeight.bold,
                  alignment: Alignment.centerLeft,
                  textAlign: TextAlign.left,
                ),
                SizedBox(
                  height: 48,
                ),
                Row(
                  children: <Widget>[
                    ContraButton(
                      color: wood_smoke,
                      borderColor: wood_smoke,
                      textColor: white,
                      text: "message",
                      size: 100,
                      textSize: 12,
                      isPrefix: false,
                      isSuffix: false,
                      shadowColor: wood_smoke,
                    ),
                    SizedBox(
                      width: 12,
                    ),
                    Icon(
                      Icons.star_border,
                      color: white,
                    ),
                    ContraText(
                      text: "4.9",
                      size: 17,
                      color: wood_smoke,
                      weight: FontWeight.w800,
                      alignment: Alignment.centerLeft,
                      textAlign: TextAlign.left,
                    ),
                  ],
                )
              ],
            ),
          ),
          Align(
            alignment: Alignment.centerRight,
            child: SvgPicture.asset(
              review.profile,
              width: 210,
              height: 230,
            ),
          ),
        ],
      ),
    );
  }
}
