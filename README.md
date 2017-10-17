# CantonRES

## Requirements

* Go 1.9+
* PostgreSQL 9.5+
* Redis 3.0+
* Node.js 6.0+

## Setup in production

### Configurations

Copy `config.example.yml` to `config.yml`, then edit it.

Since it's in production, make sure:

```
debug: false
```

### Front End Resources

1. Install Node.js 6.0+ ([nvm](http://nvm.sh) is recommended) and [yarn](https://yarnpkg.com/)
2. Install Front End dependencies with `yarn install` (or just `yarn`)
3. Build JS and CSS with `yarn run production`

`yarn run production` should be run each times the program is updated.

## Development

See [DEV.md](DEV.md)
