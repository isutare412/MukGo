# MukGo

Game-like restaurant review service.

# House Rules

- Need to know what other teammates are doing. Request a seminar if you don't
  understand.
- Follow code convention clearly.
- Follow git commit [convention](https://djkeh.github.io/articles/How-to-write-a-git-commit-message-kor/).
- `git merge` is allowed only for branch merging. Use `git rebase` for
  common use.
  + `git merge` → `git rebase`
  + `git pull` → `git fetch && git rebase`

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
