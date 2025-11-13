### `fmtc` color tags

#### Editors support

If you are using SublimeText 4 (`4075+`), we strongly recommend that you install [extended Go syntax highlighting](https://github.com/essentialkaos/blackhole-theme-sublime/blob/master/fmtc.sublime-syntax) with support for `fmtc` tags.

![#colors](../.github/images/fmtc_highlight.png)

#### Environment variables

`fmtc` supports configuration via environment variables.

- `NO_COLOR` — disable all colors and modificators;
- `FMTC_NO_BOLD` — disable **bold** text;
- `FMTC_NO_ITALIC` — disable _italic_ text;
- `FMTC_NO_BLINK` — disable blinking text.

#### Modificators

| Name          | Tag   | Reset Tag | Code |
|---------------|-------|-----------|------|
| Reset         | `{!}` | —         | `0`  |
| Bold          | `{*}` | `{!*}`    | `1`  |
| Dim           | `{^}` | `{!^}`    | `2`  |
| Italic        | `{&}` | `{!&}`    | `3`  |
| Underline     | `{_}` | `{!_}`    | `4`  |
| Blink         | `{~}` | `{!~}`    | `5`  |
| Reverse       | `{@}` | `{!@}`    | `7`  |
| Hidden        | `{+}` | `{!+}`    | `8`  |
| Strikethrough | `{=}` | `{!=}`    | `9`  |

#### 8/16 Colors

##### Foreground (Text)

| Name          | Tag   | Code  | Color Preview |
|---------------|-------|-------|---------------|
| Black         | `{d}` |  `30` | ![#color](../.github/images/color_d.svg) |
| Red           | `{r}` |  `31` | ![#color](../.github/images/color_r.svg) |
| Green         | `{g}` |  `32` | ![#color](../.github/images/color_g.svg) |
| Yellow        | `{y}` |  `33` | ![#color](../.github/images/color_y.svg) |
| Blue          | `{b}` |  `34` | ![#color](../.github/images/color_b.svg) |
| Magenta       | `{m}` |  `35` | ![#color](../.github/images/color_m.svg) |
| Cyan          | `{c}` |  `36` | ![#color](../.github/images/color_c.svg) |
| Light gray    | `{s}` |  `37` | ![#color](../.github/images/color_s.svg) |
| Dark gray     | `{s-}`|  `90` | ![#color](../.github/images/color_sl.svg) |
| Light red     | `{r-}`|  `91` | ![#color](../.github/images/color_rl.svg) |
| Light green   | `{g-}`|  `92` | ![#color](../.github/images/color_gl.svg) |
| Light yellow  | `{y-}`|  `93` | ![#color](../.github/images/color_yl.svg) |
| Light blue    | `{b-}`|  `94` | ![#color](../.github/images/color_bl.svg) |
| Light magenta | `{m-}`|  `95` | ![#color](../.github/images/color_ml.svg) |
| Light cyan    | `{c-}`|  `96` | ![#color](../.github/images/color_cl.svg) |
| White         | `{w-}`|  `97` | ![#color](../.github/images/color_w.svg) |

##### Background

| Name          | Tag   | Code   | Color Preview |
|---------------|-------|--------|---------------|
| Black         | `{D}` |  `40`  | ![#color](../.github/images/color_d.svg) |
| Red           | `{R}` |  `41`  | ![#color](../.github/images/color_r.svg) |
| Green         | `{G}` |  `42`  | ![#color](../.github/images/color_g.svg) |
| Yellow        | `{Y}` |  `43`  | ![#color](../.github/images/color_y.svg) |
| Blue          | `{B}` |  `44`  | ![#color](../.github/images/color_b.svg) |
| Magenta       | `{M}` |  `45`  | ![#color](../.github/images/color_m.svg) |
| Cyan          | `{C}` |  `46`  | ![#color](../.github/images/color_c.svg) |
| Light gray    | `{S}` |  `47`  | ![#color](../.github/images/color_s.svg) |
| Dark gray     | `{S-}`|  `100` | ![#color](../.github/images/color_sl.svg) |
| Light red     | `{R-}`|  `101` | ![#color](../.github/images/color_rl.svg) |
| Light green   | `{G-}`|  `102` | ![#color](../.github/images/color_gl.svg) |
| Light yellow  | `{Y-}`|  `103` | ![#color](../.github/images/color_yl.svg) |
| Light blue    | `{B-}`|  `104` | ![#color](../.github/images/color_bl.svg) |
| Light magenta | `{M-}`|  `105` | ![#color](../.github/images/color_ml.svg) |
| Light cyan    | `{C-}`|  `106` | ![#color](../.github/images/color_cl.svg) |
| White         | `{W-}`|  `107` | ![#color](../.github/images/color_w.svg) |

#### 88/256 Colors

![#256colors](../.github/images/256_colors.png)

##### Foreground (Text)

Tag: `{#code}`

##### Background

Tag: `{%code}`

#### 24-bit Colors (_TrueColor_)

##### Foreground (Text)

Tag: `{#hex}`

##### Background

Tag: `{%hex}`

#### Named colors

Tag: `{?name}`

For more information about named colors see documentation for method `AddColor`.

#### Examples

```
{r*}Important!{!*} File deleted!{!}
 ┬             ┬─                ┬
 │             │                 │
 │             │                 └ Reset everything
 │             │
 │             └ Unset bold modificator
 │
 └ Bold, red text 
```

```
{rG*}OMG!{!*} Check your mail{!}
 ┬──      ┬─                  ┬
 │        │                   │
 │        │                   └ Reset everything
 │        │
 │        └ Unset bold modificator
 │
 └ Bold, red text with green background
```

```
{r}File {_}file.log{!_} deleted{!}
 ┬       ┬          ┬─          ┬
 │       │          │           │ 
 │       │          │           └ Reset everything
 │       │          │
 │       │          └ Unset underline modificator
 │       │
 │       └ Set underline modificator
 │
 └ Set text color to red
```

```
{*@y}Warning!{!@} Can't find user bob.{!}
 ┬──          ┬─                       ┬
 │            │                        │
 │            │                        └ Reset everything
 │            │
 │            └ Unset reverse modificator (keep yellow color)
 │
 └ Bold text with yellow background (due to reverse modificator)
```

```
{#213}{%240}Hi all!{!}
 ┬───  ┬───         ┬
 │     │            │
 │     │            └ Reset everything
 │     │
 │     └ Set background color to grey
 │
 └ Set text color to pink
```

```
{#ff1493}{%191970}Hi all!{!}
 ┬──────  ┬──────         ┬
 │        │               │
 │        │               └ Reset everything
 │        │
 │        └ Set background color to midnightblue
 │
 └ Set text color to deeppink
```

```
{?error}Can't find user "bob"{!}
 ┬─────                       ┬
 │                            │
 │                            └ Reset everything
 │
 │
 │
 └ Set color to named color "error"
```
