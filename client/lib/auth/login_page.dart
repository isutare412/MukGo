import 'package:contra/custom_widgets/button_solid_with_icon.dart';
import 'package:contra/utils/colors.dart';
import 'package:contra/login/contra_text.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';

import 'package:mukgo/auth/auth_api.dart';

class LoginForm extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Stack(alignment: Alignment.bottomCenter, children: <Widget>[
        Container(
          color: lightening_yellow,
          alignment: Alignment.center,
          child: Row(
            crossAxisAlignment: CrossAxisAlignment.start,
            mainAxisAlignment: MainAxisAlignment.start,
            children: <Widget>[
              SvgPicture.asset(
                "assets/images/peep_standing_left.svg",
                width: 370,
                height: 590,
              ),
            ],
          ),
        ),
        Container(
          alignment: Alignment.center,
          child: Row(
            crossAxisAlignment: CrossAxisAlignment.start,
            mainAxisAlignment: MainAxisAlignment.end,
            children: <Widget>[
              SvgPicture.asset(
                "assets/images/peep_standing_right.svg",
                width: 370,
                height: 590,
              ),
            ],
          ),
        ),
        Container(
          height: 300,
          alignment: Alignment.bottomCenter,
          margin: EdgeInsets.all(24),
          padding: EdgeInsets.all(24),
          decoration: BoxDecoration(
              borderRadius: BorderRadius.circular(16), color: white),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.center,
            mainAxisAlignment: MainAxisAlignment.center,
            children: <Widget>[
              ContraText(
                text: "Login",
                alignment: Alignment.center,
              ),
              SizedBox(
                height: 10,
              ),
              Text(
                "You don’t think you should login first and behave like human not robot.",
                textAlign: TextAlign.center,
                style: TextStyle(
                    fontSize: 17, color: trout, fontWeight: FontWeight.w500),
              ),
              SizedBox(
                height: 16,
              ),
              ButtonPlainWithIcon(
                color: google_red,
                textColor: white,
                iconPath: "assets/icons/google.svg",
                isPrefix: true,
                isSuffix: false,
                text: "Google 로그인",
                callback: () async {
                  await googleSignIn(context);
                  Navigator.pop(context);
                },
              ),
            ],
          ),
        )
      ]),
    );
  }
}
