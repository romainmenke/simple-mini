# Simple Mini

A lazy man's minify

`go get github.com/romainmenke/simple-mini`

---

I needed a tool to remove whitespace from css and js files for golang web projects. I did't want a transpiler and I didn't want to add node as a dependency. Since everything is gzipped anyway I also didn't want to minify naming.

- it loops over all files in a directory
- removes whitespace
- saves to `.min` version

I use it with `//go:generate simple-mini` and [modd](https://github.com/cortesi/modd).

---

### Options

- `-h`            : help
- `-source`       : source directory
- `-out`          : output directory
- `trailing args` : exclusion -> simple `must not contain` logic

---

It removes :
- tabs
- linebreaks

It replaces :
- double spaces with single spaces

---

### Simple

- [simple-mini](https://github.com/romainmenke/simple-mini)
- [simple-bundle](https://github.com/romainmenke/simple-bundle)
- [simple-gzip](https://github.com/romainmenke/simple-gzip)
