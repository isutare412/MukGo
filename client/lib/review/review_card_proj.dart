import 'package:contra/login/contra_text.dart';
import 'package:contra/utils/colors.dart';
import 'package:flutter/material.dart';

class ReviewCardData {
  final String user;
  final String comment;
  final int score;
  final int like;
  final String time;
  final Color bgColor;
  final List<String> menus;

  const ReviewCardData(
      {this.user,
      this.comment,
      this.score,
      this.like,
      this.time,
      this.bgColor,
      this.menus});
}

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
            Expanded(
              flex: 1,
              child: Container(
                width: 100,
                height: 100,
//                child: SvgPicture.asset(
//                  "assets/icons/placeholder_icon.svg",
//                  width: 40,
//                  height: 40,
//                ),
                child: Icon(
                  Icons.image,
                  color: white,
                  size: 40,
                ),
                decoration: ShapeDecoration(
                    color: reviewData.bgColor,
                    shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.all(Radius.circular(16)),
                        side: BorderSide(color: wood_smoke, width: 2))),
              ),
            ),
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
                  Row(
                    mainAxisAlignment: MainAxisAlignment.start,
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
                              Icons.favorite_border,
                              color: wood_smoke,
                            ),
                            ContraText(
                              text: reviewData.like.toString(),
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
