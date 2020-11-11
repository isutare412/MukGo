import 'package:contra/login/contra_text.dart';
import 'package:contra/utils/colors.dart';
import 'package:flutter/material.dart';
import 'package:mukgo/review/review_card_data.dart';
import 'package:mukgo/user/user_model.dart';
import 'package:provider/provider.dart';
import 'package:flutter_svg/svg.dart';

List<String> profileAsset = [
  'assets/images/onboarding_image_five.svg',
  'assets/images/onboarding_image_one.svg',
  'assets/images/onboarding_image_two.svg',
  'assets/images/onboarding_image_three.svg',
  'assets/images/onboarding_image_four.svg'
];

class ReviewCard extends StatelessWidget {
  final ReviewCardData reviewData;
  final VoidCallback onTap;

  const ReviewCard({this.reviewData, this.onTap});

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: onTap,
      child: Container(
        margin: EdgeInsets.only(top: 24),
        padding: EdgeInsets.all(24),
        child: Row(
          children: <Widget>[
            Consumer<UserModel>(builder: (context, user, child) {
              int profileIndex;
              if (reviewData.userLevel <= 0 || reviewData.userLevel > 4) {
                profileIndex = 0;
              } else {
                profileIndex = reviewData.userLevel;
              }
              return SvgPicture.asset(
                profileAsset[profileIndex],
                height: 50,
                width: 50,
              );
            }),
            SizedBox(
              width: 24,
            ),
            Expanded(
              flex: 3,
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: <Widget>[
                  Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: <Widget>[
                      Text(
                        reviewData.time,
                        style: TextStyle(
                            fontSize: 11,
                            fontWeight: FontWeight.bold,
                            color: trout),
                      ),
                      Text(
                        reviewData.user,
                        style: TextStyle(
                            fontSize: 11,
                            fontWeight: FontWeight.bold,
                            color: trout),
                      ),
                    ],
                  ),
                  SizedBox(
                    height: 6,
                  ),
                  Text(
                    reviewData.comment,
                    overflow: TextOverflow.ellipsis,
                    maxLines: 2,
                    style: TextStyle(
                        fontWeight: FontWeight.bold,
                        color: wood_smoke,
                        fontSize: 21),
                  ),
                  Wrap(
                    spacing: 8.0,
                    runSpacing: 4.0,
                    children: reviewData.menus
                        .map((menu) => Padding(
                            padding: EdgeInsets.only(right: 6),
                            child: Text(
                              '#$menu',
                              style: TextStyle(
                                  fontSize: 11,
                                  fontWeight: FontWeight.bold,
                                  color: trout),
                            )))
                        .toList(),
                  ),
                  SizedBox(
                    height: 6,
                  ),
                  Row(
                    children: <Widget>[
                      Expanded(
                        flex: 1,
                        child: Row(
                          children: <Widget>[
                            Icon(
                              Icons.stars,
                              color: wood_smoke,
                            ),
                            ContraText(
                              text: reviewData.score.toString(),
                              size: 13,
                              alignment: Alignment.center,
                            )
                          ],
                        ),
                      ),
                      Expanded(
                        flex: 1,
                        child: Row(
                          children: <Widget>[
                            Icon(
                              Icons.person,
                              color: wood_smoke,
                            ),
                            ContraText(
                              text: reviewData.numPeople.toString(),
                              size: 13,
                              alignment: Alignment.center,
                            )
                          ],
                        ),
                      )
                    ],
                  )
                ],
              ),
            )
          ],
        ),
      ),
    );
  }
}
