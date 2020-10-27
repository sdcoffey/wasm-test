# Sonder WASM Demo

## To run
You need:
- Go installed on your machine

```bash
$ make boostrap # one-time setup
$ make run
```

## Overview

In this project:
1. `wasm_exec.js`: The Golang js runtime. Copied from `$GOROOT/misc/wasm`
1. `index.html`: HTML page which loads `wasm_exec.js` and `main.js`.
1. `main.go`: Our go file which exposes go code to js
1. `main.js`: our main js file which loads our javascript and makes calls into webassembly
