# ffpreset

## Build
- Install [Go](https://golang.org/dl/)
- Run `go build`

## Usage
**Prerequisites:**
- Add `ffmpeg` to your `$PATH` (`webp`, `avif` (`libaom`) support)
- Add `cjxl` to your `$PATH` (`cjxl` support)

**Usage:**
```bash
ffpreset [format] <input1> <input2> ...
```

**Formats:**
- `webp`, `w` ➡️ WebP
- `avif`, `a` ➡️ AVIF
- `cjxl`, `j` ➡️ Jpeg XL