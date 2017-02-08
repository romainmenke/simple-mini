# Simple Mini

A lazy mans minify

`go get github.com/romainmenke/simple-mini`

---

I needed a tool to remove whitespace from css and js files for golang web projects. I did't want a transpiler and I didn't want to add node as a dependency.
Since everything is gzipped anyway I also didn't want to minify naming.

It finds files with `.js` or `.css` extensions and generates `.min.js` and `.min.css` files. `.min` files are obviously ignored.

I use it with `//go:generate simple-mini` and [modd](https://github.com/cortesi/modd).

---

It removes :
- tabs
- linebreaks

It replaces :
- double spaces with single spaces

---

Size reduction depends heavily on coding style.
For me it averages to 20% (I like whitespace)
