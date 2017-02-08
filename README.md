# Simple Mini

A lazy mans minify

---

I needed a tool to remove whitespace from css and js files for golang web projects. I did't want a transpiler and I didn't want to add node as a dependency.
Since everything is gzipped anyway I also didn't want to minify naming.

It finds files with `.js` or `.css` extensions and generates `.min.js` and `.min.css` files. `.min` files are obviously ignored.

I use it with `//go:generate simple-mini`.

---

It removes :
- tabs
- linebreaks

It replaces :
- double spaces with single spaces

---

Size reduction depends heavily on coding style.
For me it averages to 20% (I like whitespace)
