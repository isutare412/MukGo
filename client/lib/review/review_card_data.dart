import 'dart:ui';

class ReviewCardData {
  final String reviewId;
  final String user;
  final String comment;
  final int score;
  final int numPeople;
  final String time;
  final List<String> menus;
  final bool waiting;
  final int userLevel;
  final int likeCount;
  final bool likeByMe;

  const ReviewCardData(
      {this.reviewId,
      this.user,
      this.comment,
      this.score,
      this.numPeople,
      this.time,
      this.menus,
      this.waiting,
      this.userLevel,
      this.likeCount,
      this.likeByMe});
}
