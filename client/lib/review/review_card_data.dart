import 'dart:ui';

class ReviewCardData {
  final String user;
  final String comment;
  final int score;
  final int numPeople;
  final String time;
  final Color bgColor;
  final List<String> menus;
  final bool waiting;

  const ReviewCardData(
      {this.user,
      this.comment,
      this.score,
      this.numPeople,
      this.time,
      this.bgColor,
      this.menus,
      this.waiting});
}
