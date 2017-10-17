# Development

## Requirements

* Go 1.9+
* PostgreSQL 9.5+
* Redis 3.0+
* Node.js 6.0+

## Environment

* Use `yarn` rather than `npm` (for lockfile support)

## Serve

```bash
# Run Server
go run main.go

# Webpack Watch
yarn run watch
```

## Code Standard

### General

* Use `UTF-8` for file encoding
* Use `LF` for line breaking
* Reserve an empty line at the end of file

### Go

#### Package Layers

* 5 - main
* 4 - middleware/controller
* 3 - util
* 2 - db
* 1 - service

To avoid import loop, a package could only import packages lower then itself (including subpackages) or those inside `$GOPATH`.

Example:

```go
package util

import (
    _ "../db" // OK
    _ "../service" // OK
    _ "../middleware" // DON'T
)
```

```go
package db

import (
    _ "../util" // DON'T    
    _ "github.com/jinzhu/gorm" // OK
)
```

### ECMAScript (JavaScript)

* Indent with 4 spaces
* Do not indent empty lines
* Opening braces go on the same line
* Single quotes for strings
    * Except to avoid escaping
* Keep semicolons for each statements
    * Except for `export { ... }` or `return { ... }`
* Use `let` to declare local variables rather than `var`
* Avoid declaring unused variables
* Each variable has its own declaration statement
* Keep space after keywords and commas
* Keep space before and after infix operators
* No space after function names
* ES2015 import/export syntax

Example:

```js
import pangu from 'pangu';

function hello(text) {
    let str = text || 'Hello世界!';
    let length = str.length;

    if (typeof str === 'string')
        str = pangu.spacing(str);
    else
        str = 'Hello 世界!';

    return { str, length }
}

class Greeter {
    constructor(name) {
        this.name = name;
    }

    memberMethod() {
        if (this.name.length < 100) {
            let mixin = {
                quote: 'It\'s good.'
            };

            let obj = {
                ...mixin,
                quote: "It's better.",
                reverse: str => str.split('').reverse().join('')
            };
        } else {
            window.alert(this.name);
        }

        return this.name;
    }
}

export { hello as default, Greeter }
```

### CSS

* Use `scss` rather than `sass`
* No prefixes. [autoprefixer](https://github.com/postcss/autoprefixer) does this well.
