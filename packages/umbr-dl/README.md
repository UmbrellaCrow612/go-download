# Umbr download

A npm package that wraps a go binary to download files or folders really fast using go

# Install

```bash
npm i umbr-dl --save-dev
```

# Usage

```js
const { download } = require("umbr-dl");

download("URL", { name: "download.zip", path: "." });
// Download the URL, names it download.zip and saves it in the current working directory
```
