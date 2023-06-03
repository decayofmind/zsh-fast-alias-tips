# zsh-fast-alias-tips

Helps you remembering the aliases you defined once.

sei40kr/fast-alias-tips-bin

Written in zsh and Go. Ported from [djui/alias-tips](https://github.com/djui/alias-tips).

## Example

```
$ alias gst='git status'

$ git status
💡  gst
On branch master
Your branch is up to date with 'origin/master'.

nothing to commit, working tree clean
```

## Install

### Install with [zinit](https://github.com/zdharma/zinit)

```sh
zinit ice from'gh-r' as'program'
zinit light sei40kr/fast-alias-tips-bin
zinit light sei40kr/zsh-fast-alias-tips
```

## Configuration

| Variable                     | Default value       | Description           |
| :--                          | :--                 | :--                   |
| `ZSH_FAST_ALIAS_TIPS_PREFIX` | `"💡 $(tput bold)"` | The prefix of the Tips |
| `ZSH_FAST_ALIAS_TIPS_SUFFIX` | `"$(tput sgr0)"`    | The suffix of the Tips |
