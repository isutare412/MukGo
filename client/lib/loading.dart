import 'dart:async';
import 'package:flutter/material.dart';
import 'package:animated_text_kit/animated_text_kit.dart';
import 'package:contra/utils/colors.dart';

class LoadingScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    Timer(Duration(milliseconds: 1200), () {
      Navigator.pushReplacementNamed(context, '/project');
    });
    return Scaffold(
      body: Container(
        color: lightening_yellow,
        child: Center(
          child: TextLiquidFill(
            loadDuration: Duration(milliseconds: 1000),
            waveDuration: Duration(milliseconds: 700),
            text: 'MukGo',
            waveColor: persian_blue,
            boxBackgroundColor: lightening_yellow,
            textStyle: TextStyle(
              fontSize: 80.0,
              fontWeight: FontWeight.bold,
            ),
            boxHeight: 300.0,
          ),
        ),
      ),
    );
  }
}
