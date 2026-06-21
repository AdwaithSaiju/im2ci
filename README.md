# im2ci

Convert images to ASCII art using Floyd-Steinberg dithering.

## Usage

```
im2ci [-i] [-w width] <image>
```

| Flag | Default | Description |
|------|---------|-------------|
| `-i` | off | Invert character mapping for dark/transparent terminals |
| `-w` | 100 | Output width in characters |

## Algorithm

### 1. Resize

The image is resized using nearest-neighbour interpolation. The height is halved relative to the width to compensate for the typical 2:1 aspect ratio of terminal characters (characters are roughly twice as tall as they are wide).

### 2. Luminance

Each pixel is converted to a grayscale value using the BT.601 luma coefficients:

```
Y = 0.299*R + 0.587*G + 0.114*B
```

The `RGBA()` method returns 16-bit colour values (0-65535), so luminance is computed in that range.

### 3. Character mapping

Ten characters are used, ordered by approximate ink coverage:

```
@ % # * + - = : . <space>
```

`@` has the most ink (appears darkest on a light background, brightest on a dark background). `<space>` has the least. The `-i` flag reverses this order.

### 4. Floyd-Steinberg dithering

The 10-level quantisation produces visible banding. Floyd-Steinberg error diffusion distributes the quantisation error to neighbouring pixels (left-to-right, top-to-bottom scan):

```
          x    7/16
   3/16  5/16  1/16
```

The error at each pixel is spread to the right, bottom-left, bottom, and bottom-right neighbours. This preserves the perceived brightness of the original image despite the limited character set.

## Install

```
go build -o im2ci
```

Requires Go 1.26 or later. Supports PNG and JPEG inputs.
