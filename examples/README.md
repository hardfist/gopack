# Gopack Configuration

This directory contains example configuration files for Gopack.

## gopack.config.json

Basic configuration file format:

```json
{
  "entry": "./src/index.js",
  "output": {
    "dir": "./dist",
    "filename": "bundle.js"
  },
  "mode": "development",
  "resolve": {
    "extensions": [".js", ".jsx", ".ts", ".tsx"]
  },
  "module": {
    "rules": []
  }
}
```
