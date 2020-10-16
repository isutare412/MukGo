import 'package:contra/utils/colors.dart';
import 'package:contra/custom_widgets/button_plain.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:provider/provider.dart';
import 'package:mukgo/user/user_model.dart';

class UserDetailTestPage extends StatefulWidget {
  const UserDetailTestPage({Key key}) : super(key: key);

  @override
  _UserDetailTestPageState createState() => _UserDetailTestPageState();
}

class _UserDetailTestPageState extends State<UserDetailTestPage> {
  @override
  void initState() {
    super.initState();
    Future.microtask(() {
      // fetch user info after randering
      return context.read<UserModel>().fetch();
    });
  }

  @override
  Widget build(BuildContext context) {
    return Material(
      child: Container(
        color: white,
        child: Column(
          children: <Widget>[
            Expanded(
              flex: 4,
              child: Container(
                child: Column(
                  mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                  children: <Widget>[
                    SizedBox(
                      height: 30,
                    ),
                    Center(
                      child: SvgPicture.asset(
                        'assets/images/onboarding_image_one.svg',
                        height: 320,
                        width: 320,
                      ),
                    ),
                  ],
                ),
              ),
            ),
            Expanded(
              flex: 2,
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.center,
                children: <Widget>[
                  Padding(
                    padding: const EdgeInsets.only(
                        left: 24.0, right: 24.0, top: 12.0, bottom: 12.0),
                    child: Consumer<UserModel>(builder: (context, user, child) {
                      return Text(
                        user.name ?? '로딩중..',
                        textAlign: TextAlign.center,
                        style: TextStyle(
                            fontSize: 36,
                            color: wood_smoke,
                            fontWeight: FontWeight.w800),
                      );
                    }),
                  ),
                  Padding(
                    padding: const EdgeInsets.only(
                        left: 24.0, right: 24.0, top: 12.0, bottom: 12.0),
                    child: Consumer<UserModel>(builder: (context, user, child) {
                      return Text(
                        user.level != null
                            ? '레벨: ${user.level}, 총경험치: ${user.totalExp}입니다.'
                            : '',
                        textAlign: TextAlign.center,
                        style: TextStyle(
                            fontSize: 21,
                            color: trout,
                            fontWeight: FontWeight.w500),
                      );
                    }),
                  ),
                  Padding(
                    padding: const EdgeInsets.only(
                        left: 24.0, right: 24.0, top: 12.0, bottom: 12.0),
                    child: ButtonPlain(
                      color: google_red,
                      textColor: white,
                      text: '다시 불러오기',
                      onTap: () async {
                        var user = context.read<UserModel>();
                        user.clear();
                        await user.fetch();
                        print('reloaded');
                      },
                    ),
                  ),
                ],
              ),
            ),
            Expanded(
              flex: 1,
              child: Container(),
            )
          ],
        ),
      ),
    );
  }
}
