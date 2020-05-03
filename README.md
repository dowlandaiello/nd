# nd
Theme management for night and day.

## What is this?

Tired of manually switching between light and dark themes? Nd automates that
process with just one yml file. By default, nd looks for a `links.yml` file
in your home XDG config directory. More specifically, on macOS,
`~/Library/Preferences/nd/links.yml`, on Unix systems,
`~/.config/nd/links.yml`, and on Windows, `%LOCALAPPDATA%\nd\links.yml`.

For example, suppose you'd like to use different vimrc files depending on the
time of day:

```yml
(~/.config/nd/links.yml)
vim:
  to: ~/.vimrc
  day:
    from: ~/.vimrc.day
  night:
    from: ~/.vimrc.night
```

Toggling between the two aforementioned vimrc files is as simple as `nd day` or
`nd night`!

## Installation

```zsh
go install github.com/dowlandaiello/nd
```
