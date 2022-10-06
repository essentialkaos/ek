### `fmtc` color tags

#### Modificators

| Name      | Tag   | Code |
|-----------|-------|------|
| Reset     | `{!}` | `0`  |
| Bold      | `{*}` | `1`  |
| Dim       | `{^}` | `2`  |
| Underline | `{_}` | `4`  |
| Blink     | `{~}` | `5`  |
| Reverse   | `{@}` | `7`  |

#### 8/16 Colors

##### Foreground (Text)

| Name          | Tag   | Code  | Color Preview |
|---------------|-------|-------|---------------|
| Black         | `{d}` |  `30` | ![#color](https://via.placeholder.com/100x16/000000/000000?text=+) |
| Red           | `{r}` |  `31` | ![#color](https://via.placeholder.com/100x16/CC0000/000000?text=+) |
| Green         | `{g}` |  `32` | ![#color](https://via.placeholder.com/100x16/4D9A05/000000?text=+) |
| Yellow        | `{y}` |  `33` | ![#color](https://via.placeholder.com/100x16/C4A000/000000?text=+) |
| Blue          | `{b}` |  `34` | ![#color](https://via.placeholder.com/100x16/3465A4/000000?text=+) |
| Magenta       | `{m}` |  `35` | ![#color](https://via.placeholder.com/100x16/754F7B/000000?text=+) |
| Cyan          | `{c}` |  `36` | ![#color](https://via.placeholder.com/100x16/069899/000000?text=+) |
| Light gray    | `{s}` |  `37` | ![#color](https://via.placeholder.com/100x16/D3D7CE/000000?text=+) |
| Dark gray     | `{s-}`|  `90` | ![#color](https://via.placeholder.com/100x16/555752/000000?text=+) |
| Light red     | `{r-}`|  `91` | ![#color](https://via.placeholder.com/100x16/EE2828/000000?text=+) |
| Light green   | `{g-}`|  `92` | ![#color](https://via.placeholder.com/100x16/8AE234/000000?text=+) |
| Light yellow  | `{y-}`|  `93` | ![#color](https://via.placeholder.com/100x16/FCE94F/000000?text=+) |
| Light blue    | `{b-}`|  `94` | ![#color](https://via.placeholder.com/100x16/729FCE/000000?text=+) |
| Light magenta | `{m-}`|  `95` | ![#color](https://via.placeholder.com/100x16/AD7EA8/000000?text=+) |
| Light cyan    | `{c-}`|  `96` | ![#color](https://via.placeholder.com/100x16/34E1E1/000000?text=+) |
| White         | `{w-}`|  `97` | ![#color](https://via.placeholder.com/100x16/EEEEEC/000000?text=+) |

##### Background

| Name          | Tag   | Code   | Color Preview |
|---------------|-------|--------|---------------|
| Black         | `{D}` |  `40`  | ![#color](https://via.placeholder.com/100x16/000000/000000?text=+) |
| Red           | `{R}` |  `41`  | ![#color](https://via.placeholder.com/100x16/CC0000/000000?text=+) |
| Green         | `{G}` |  `42`  | ![#color](https://via.placeholder.com/100x16/4D9A05/000000?text=+) |
| Yellow        | `{Y}` |  `43`  | ![#color](https://via.placeholder.com/100x16/C4A000/000000?text=+) |
| Blue          | `{B}` |  `44`  | ![#color](https://via.placeholder.com/100x16/3465A4/000000?text=+) |
| Magenta       | `{M}` |  `45`  | ![#color](https://via.placeholder.com/100x16/754F7B/000000?text=+) |
| Cyan          | `{C}` |  `46`  | ![#color](https://via.placeholder.com/100x16/069899/000000?text=+) |
| Light gray    | `{S}` |  `47`  | ![#color](https://via.placeholder.com/100x16/D3D7CE/000000?text=+) |
| Dark gray     | `{S-}`|  `100` | ![#color](https://via.placeholder.com/100x16/555752/000000?text=+) |
| Light red     | `{R-}`|  `101` | ![#color](https://via.placeholder.com/100x16/EE2828/000000?text=+) |
| Light green   | `{G-}`|  `102` | ![#color](https://via.placeholder.com/100x16/8AE234/000000?text=+) |
| Light yellow  | `{Y-}`|  `103` | ![#color](https://via.placeholder.com/100x16/FCE94F/000000?text=+) |
| Light blue    | `{B-}`|  `104` | ![#color](https://via.placeholder.com/100x16/729FCE/000000?text=+) |
| Light magenta | `{M-}`|  `105` | ![#color](https://via.placeholder.com/100x16/AD7EA8/000000?text=+) |
| Light cyan    | `{C-}`|  `106` | ![#color](https://via.placeholder.com/100x16/34E1E1/000000?text=+) |
| White         | `{W-}`|  `107` | ![#color](https://via.placeholder.com/100x16/EEEEEC/000000?text=+) |

#### 88/256 Colors

##### Foreground (Text)

Tag: `{#code}`

![#colors](../.github/images/256_colors_fg.png)

_Image from [FLOZz' MISC](https://misc.flogisoft.com/bash/tip_colors_and_formatting) website_

##### Background

Tag: `{%code}`

![#colors](../.github/images/256_colors_bg.png)

_Image from [FLOZz' MISC](https://misc.flogisoft.com/bash/tip_colors_and_formatting) website_

#### 24-bit Colors (_TrueColor_)

##### Foreground (Text)

Tag: `{#hex}`

##### Background

Tag: `{%hex}`

#### Named colors

Tag: `{?name}`

For more information about named colors see documentation for method `NameColor`.

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
{rG*}OMG!{!} Check your mail{!}
 ┬──                         ┬
 │                           │
 │                           └ Reset everything
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