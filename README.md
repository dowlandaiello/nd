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

Or, if you'd like to toggle between two tmux configs and reload tmux at the
same time (cmd is optional here, and just allows for a live reload):

```yml
(~/.config/nd/links.yml)
vim:
...
tmux:
  to: ~/.tmux.conf
  day:
    from: ~/.tmux.day.conf
    cmd: tmux source-file ~/.tmux.conf
  night:
    from: ~/.tmux.night.conf
    cmd: tmux source-file ~/.tmux.conf
```

Toggling between the two aforementioned vimrc files is as simple as `nd day` or
`nd night`!

## Installation

```zsh
go install github.com/dowlandaiello/nd
```
