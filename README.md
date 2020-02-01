# grocerylistsbot

A Telegram Bot to track your groceries :) [t.me/grocerylistsbot](https://t.me/grocerylistsbot)

## Installation

```bash
$ make 

# postgres
$ docker run -itd --rm -p 5432:5432 -e ALLOW_EMPTY_PASSWORD=yes bitnami/postgresql

# the bot
$ TELEGRAM_TOKEN=my_token_here ./bin/grocerylistsbot
```

## License

Apache-2.0