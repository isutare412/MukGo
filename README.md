# MukGo

![Build Status](http://redshore.asuscomm.com:8081/buildStatus/icon?job=MukGo%2Frelease)

![Mukgo service diagram](./data/Mukgo%20service%20diagram.png)

Game-like restaurant review service.

# Screenshots

- User page
- Achievement badges

<img src="./data/screenshots/user_page.png" width="220">
<img src="./data/screenshots/achievement_badge.png" width="220">

- Mukgoers ranking

<img src="./data/screenshots/ranking.png" width="220">

- Navigate restaurant within your sight

<img src="./data/screenshots/map_page.png" width="220">
<img src="./data/screenshots/map_tapped.png" width="220">

- Restaurant reviews
- Leave likes on reivew

<img src="./data/screenshots/restaurant_reviews.png" width="220">
<img src="./data/screenshots/filtering_sorting.png" width="220">
<img src="./data/screenshots/review_page.png" width="220">

- Leave your own review

<img src="./data/screenshots/review_post.png" width="220">

# House Rules

- Need to know what other teammates are doing. Request a seminar if you don't
  understand.
- Follow code convention clearly.
- Follow git commit [convention](https://djkeh.github.io/articles/How-to-write-a-git-commit-message-kor/).
- `git merge` is allowed only for branch merging. Use `git rebase` for
  common use.
  - `git merge` → `git rebase`
  - `git pull` → `git fetch && git rebase`

# How to run

## Server

- Install [`Docker Compose`](https://docs.docker.com/compose/install/)
- ```sh
  # Build services
  docker-compose build

  # Runs all services
  docker-compose up

  # Run services background
  docker-compose up -d

  # Stop background services
  docker-compose down
  ```

## Client

- Follow [Flutter VScode guide](https://flutter.dev/docs/get-started/test-drive?tab=vscode)

## Protobuf

- Define custome error code between server and client.

### How to add new error code

1. Fix `protobuf/code.proto`

2. `make protocol` (need docker)

3. New error codes will be compiled to files.
   - `server/api/code`
   - `client/lib/protocol`