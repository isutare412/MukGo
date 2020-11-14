import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/svg.dart';
import 'package:contra/login/contra_text.dart';
import 'package:contra/utils/colors.dart';

import 'package:mukgo/ranking/ranking_user.dart';

class RankingListItem extends StatelessWidget {
  final RankingUser user;

  const RankingListItem({this.user});

  Color getBgColorFromRank(int rank) {
    if (rank == 1) {
      return lightening_yellow;
    } else if (rank == 2) {
      return santas_gray;
    } else if (rank == 3) {
      return flamingo;
    } else {
      return lavandar_bush;
    }
  }

  @override
  Widget build(BuildContext context) {
    var itemColor = user.rank < 4 ? white : trout;

    return Container(
      margin: EdgeInsets.symmetric(vertical: 12),
      decoration: ShapeDecoration(
          color: getBgColorFromRank(user.rank),
          shape: RoundedRectangleBorder(
              borderRadius: BorderRadius.all(Radius.circular(16)),
              side: BorderSide(color: wood_smoke, width: 2))),
      child: Stack(
        children: [
          Align(
            alignment: Alignment.centerRight,
            child: SvgPicture.asset(
              user.asset,
              width: 210,
              height: 230,
            ),
          ),
          Padding(
            padding: const EdgeInsets.only(top: 12, left: 24, right: 24),
            child: Column(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              crossAxisAlignment: CrossAxisAlignment.start,
              mainAxisSize: MainAxisSize.max,
              children: <Widget>[
                ContraText(
                  text: '# ${user.rank.toString()}',
                  size: 36,
                  color: wood_smoke,
                  weight: FontWeight.w800,
                  alignment: Alignment.centerLeft,
                  textAlign: TextAlign.left,
                ),
                ContraText(
                  text: user.name,
                  size: 48,
                  color: itemColor,
                  weight: FontWeight.w800,
                  alignment: Alignment.centerLeft,
                  textAlign: TextAlign.left,
                ),
                SizedBox(
                  height: 60,
                ),
                Row(
                  children: <Widget>[
                    Icon(
                      Icons.person,
                      color: itemColor,
                    ),
                    SizedBox(
                      width: 12,
                    ),
                    ContraText(
                      text: user.level.toString(),
                      size: 17,
                      color: wood_smoke,
                      weight: FontWeight.w800,
                      alignment: Alignment.centerLeft,
                      textAlign: TextAlign.left,
                    ),
                    SizedBox(
                      width: 48,
                    ),
                    Icon(
                      Icons.rate_review,
                      color: itemColor,
                    ),
                    SizedBox(
                      width: 12,
                    ),
                    ContraText(
                      text: user.reviewCount.toString(),
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
        ],
      ),
    );
  }
}
