# LotusCMS

`Simple, Fast, Flexible`

#### Main Features

* One binary file to deploy
* SEO tools (generate static pagws and headers)
* AMP / Facebook Instant Articles support
* The only two dependencies:
  1.  [Golang](https://golang.org/)
  2.  [SQLite3](https://www.sqlite.org/)

#### Possible Features

* _PostgreSQL / MySQL as main database_
* Redis as cache
* Wordpress immigration tool
* Configurable storage locations ( **other folders** or via http etc.)
* Hooks for extensions
  > #### Admin Panel
  >
  > Compiled Angular App

```
TO DO:
1. [done] Import SQLite code and create cgo wrappers
2. Create router
3. Initialise database
4. ...
```

##Notes:

* To speed up compilling time with cgo, we can firstly generate a compiled object (.o .a file) by

  > gcc -g -c -Wall sqlite3.c -o .\sqlite3.o

  and then replace

  > #cgo amd64 386 CFLAGS: -DX86=1
  > #include <sqlite3.c>

  with

  > #cgo LDFLAGS: ${SRCDIR}/sqlite3.o
  > #include <sqlite3.h>

* Add [sqlite3 source code](https://www.sqlite.org/download.html) (sqlite3.c, sqlite3.h, sqlite3ext.h and extensions: json1, fts5, icu) by yourself

* To enable extensions:

  #cgo CFLAGS: -DSQLITE_ENABLE_FTS5
  #cgo LDFLAGS: -lm

  #cgo CFLAGS: -DSQLITE_ENABLE_JSON1

  #cgo LDFLAGS: -licuuc -licui18n
  #cgo CFLAGS: -DSQLITE_ENABLE_ICU
