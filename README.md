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

A 10-level character set (`@ % # * + - = : . <space>`) can only represent 10 discrete luminance levels. Without dithering, smooth gradients break into visible contour bands. Floyd-Steinberg error diffusion fixes this by distributing quantisation error to neighbouring pixels, preserving the original image's perceived brightness.

**How it works**

Each pixel is processed left-to-right, top-to-bottom. For a pixel at position `(x, y)`:

1. Map its luminance (0-65535) to the nearest character index (0-9).
2. Compute the quantisation error: the difference between the original luminance and the luminance that the chosen character actually represents.
3. Spread this error to four neighbours that haven't been processed yet:

```
           x    7/16     (pixel to the right)
    3/16  5/16  1/16     (bottom-left, bottom, bottom-right)
```

The coefficients (7, 3, 5, 1) sum to 16, so total image brightness is conserved. Each neighbour's luminance value is adjusted by `error * coefficient / 16` before it is processed.

**Example**

A pixel with luminance 40000 maps to character `*` (which represents luminance 43690). The error is `40000 - 43690 = -3690`. This error is distributed:

- Pixel `(x+1, y)` gets `-3690 * 7/16 ≈ -1614` added to its luminance
- Pixel `(x-1, y+1)` gets `-3690 * 3/16 ≈ -692`
- Pixel `(x, y+1)` gets `-3690 * 5/16 ≈ -1153`
- Pixel `(x+1, y+1)` gets `-3690 * 1/16 ≈ -231`

This error diffusion means the local average of quantised pixels closely matches the original continuous-tone image, producing smooth gradients instead of banding.

## Install

```
go build -o im2ci
```

Requires Go 1.26 or later. Supports PNG and JPEG inputs.
