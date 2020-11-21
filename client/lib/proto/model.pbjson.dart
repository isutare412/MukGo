///
//  Generated code. Do not modify.
//  source: proto/model.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

const Code$json = const {
  '1': 'Code',
  '2': const [
    const {'1': 'SUCCESS', '2': 0},
    const {'1': 'INTERNAL_ERROR', '2': 1},
    const {'1': 'METHOD_NOT_ALLOWED', '2': 2},
    const {'1': 'AUTH_FAILED', '2': 3},
    const {'1': 'PROTOCOL_MISMATCH', '2': 4},
    const {'1': 'INVALID_DATA', '2': 5},
    const {'1': 'USER_EXISTS', '2': 6},
    const {'1': 'USER_NOT_EXISTS', '2': 7},
    const {'1': 'RESTAURANT_NOT_EXISTS', '2': 8},
    const {'1': 'LIKE_EXISTS', '2': 9},
    const {'1': 'NO_PERMISSION', '2': 10},
  ],
};

const RestaurantType$json = const {
  '1': 'RestaurantType',
  '2': const [
    const {'1': 'INVALID', '2': 0},
    const {'1': 'CHICKEN', '2': 1},
    const {'1': 'CAFE', '2': 2},
    const {'1': 'FASTFOOD', '2': 3},
    const {'1': 'MEAT', '2': 4},
    const {'1': 'DESSERT', '2': 5},
    const {'1': 'JAPANESE', '2': 6},
    const {'1': 'KOREAN', '2': 7},
    const {'1': 'CHINESE', '2': 8},
    const {'1': 'WESTERN', '2': 9},
    const {'1': 'DRINK', '2': 10},
    const {'1': 'MISC', '2': 11},
  ],
};

const RestaurantTypeCount$json = const {
  '1': 'RestaurantTypeCount',
  '2': const [
    const {'1': 'type', '3': 1, '4': 1, '5': 14, '6': '.proto.RestaurantType', '10': 'type'},
    const {'1': 'count', '3': 2, '4': 1, '5': 5, '10': 'count'},
  ],
};

const User$json = const {
  '1': 'User',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
    const {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'level', '3': 3, '4': 1, '5': 5, '10': 'level'},
    const {'1': 'total_exp', '3': 4, '4': 1, '5': 3, '10': 'totalExp'},
    const {'1': 'level_exp', '3': 5, '4': 1, '5': 3, '10': 'levelExp'},
    const {'1': 'cur_exp', '3': 6, '4': 1, '5': 3, '10': 'curExp'},
    const {'1': 'exp_ratio', '3': 7, '4': 1, '5': 1, '10': 'expRatio'},
    const {'1': 'sight_radius', '3': 8, '4': 1, '5': 1, '10': 'sightRadius'},
    const {'1': 'review_count', '3': 9, '4': 1, '5': 5, '10': 'reviewCount'},
    const {'1': 'like_count', '3': 10, '4': 1, '5': 5, '10': 'likeCount'},
    const {'1': 'restaurant_type_count', '3': 11, '4': 3, '5': 11, '6': '.proto.RestaurantTypeCount', '10': 'restaurantTypeCount'},
  ],
};

const Coordinate$json = const {
  '1': 'Coordinate',
  '2': const [
    const {'1': 'latitude', '3': 1, '4': 1, '5': 1, '10': 'latitude'},
    const {'1': 'longitude', '3': 2, '4': 1, '5': 1, '10': 'longitude'},
  ],
};

const Restaurant$json = const {
  '1': 'Restaurant',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
    const {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'coord', '3': 3, '4': 1, '5': 11, '6': '.proto.Coordinate', '10': 'coord'},
    const {'1': 'type', '3': 4, '4': 1, '5': 14, '6': '.proto.RestaurantType', '10': 'type'},
  ],
};

const Review$json = const {
  '1': 'Review',
  '2': const [
    const {'1': 'review_id', '3': 1, '4': 1, '5': 9, '10': 'reviewId'},
    const {'1': 'user_id', '3': 2, '4': 1, '5': 9, '10': 'userId'},
    const {'1': 'user_name', '3': 3, '4': 1, '5': 9, '10': 'userName'},
    const {'1': 'score', '3': 4, '4': 1, '5': 5, '10': 'score'},
    const {'1': 'comment', '3': 5, '4': 1, '5': 9, '10': 'comment'},
    const {'1': 'menus', '3': 6, '4': 3, '5': 9, '10': 'menus'},
    const {'1': 'wait', '3': 7, '4': 1, '5': 8, '10': 'wait'},
    const {'1': 'num_people', '3': 8, '4': 1, '5': 5, '10': 'numPeople'},
    const {'1': 'timestamp', '3': 9, '4': 1, '5': 3, '10': 'timestamp'},
    const {'1': 'user_level', '3': 10, '4': 1, '5': 5, '10': 'userLevel'},
    const {'1': 'like_count', '3': 11, '4': 1, '5': 5, '10': 'likeCount'},
    const {'1': 'liked_by_me', '3': 12, '4': 1, '5': 8, '10': 'likedByMe'},
  ],
};

const Users$json = const {
  '1': 'Users',
  '2': const [
    const {'1': 'users', '3': 1, '4': 3, '5': 11, '6': '.proto.User', '10': 'users'},
  ],
};

const Restaurants$json = const {
  '1': 'Restaurants',
  '2': const [
    const {'1': 'restaurants', '3': 1, '4': 3, '5': 11, '6': '.proto.Restaurant', '10': 'restaurants'},
  ],
};

const Reviews$json = const {
  '1': 'Reviews',
  '2': const [
    const {'1': 'reviews', '3': 1, '4': 3, '5': 11, '6': '.proto.Review', '10': 'reviews'},
  ],
};

const ErrorReason$json = const {
  '1': 'ErrorReason',
  '2': const [
    const {'1': 'code', '3': 1, '4': 1, '5': 14, '6': '.proto.Code', '10': 'code'},
  ],
};

