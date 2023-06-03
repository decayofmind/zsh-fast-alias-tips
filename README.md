# zsh-fast-alias-tips

A Zsh plugin to help remembering those shell aliases you once defined.

Ported from [djui/alias-tips](https://github.com/djui/alias-tips). Written in Zsh and Go, so __10x__ faster!


⚠️  __This is maintainable [fork](#differences) of [sei40kr/zsh-fast-alias-tips](https://github.com/sei40kr/zsh-fast-alias-tips).__ ⚠️


## 🖥️ Example

```
$ alias gst='git status'

$ git status
💡  gst
On branch master
Your branch is up to date with 'origin/master'.

nothing to commit, working tree clean
```

## 📦 Installation

### [zinit](https://github.com/zdharma-continuum/zinit)

From GitHub Releases (__RECOMMENDED__):

```sh
zinit ice from'gh-r' as'program'
zinit light decayofmind/zsh-fast-alias-tips
```

Compile from sources (assume you have proper version of Golang installed. __NOT__ recommended):

```sh
zinit ice atclone"make build" atpull"%atclone"
zinit light decayofmind/zsh-fast-alias-tips
```

## ⚙️  Configuration

| Variable                       | Default value         | Description                                     |
| :--                            | :-------------------- | :---------------------------------------------- |
| `ZSH_FAST_ALIAS_TIPS_EXCLUDES` | ` `                   | List of aliases to exclude (separated by space) |
| `ZSH_FAST_ALIAS_TIPS_PREFIX`   | `"💡 $(tput bold)"`   | The prefix of the Tips                          |
| `ZSH_FAST_ALIAS_TIPS_SUFFIX`   | `"$(tput sgr0)"`      | The suffix of the Tips                          |

## Differences from sei40kr/zsh-fast-alias-tips <a id='differences'></a>

I've been using the original plugin from [@sei40kr](https://github.com/tsu1980) for many years.
However, since 2020 the project is not receiving much attention from the author, with issues and PRs not being addressed.

So I decided to fork the project to address some of the issues and PRs, which resulted even in refactoring and restructuring.

Here's the full list of changes between this fork and the original project:

* Zsh plugin is released together with binary, so second repo is not needed.
* Golang upgraded to 1.20.
* It's possible to compile the plugin's binary on installation.
* Configuration variables are prefixed with plugin name (`ZSH_FAST_ALIAS_TIPS_`).
* `def-mathcer` renamed to `zsh-alias-matcher` to provide more context.
* The plugin doesn't not require binary to be in `PATH`.
* Repo is structured as a common Golang package (with `cmd` and `internal`).
* Fixed `Makefile`.
* Added ci job to GitHub workflow, which will run [golangci-lint](https://github.com/golangci/golangci-lint) and tests.
* Do not show tip if input command is shorter ([#27](https://github.com/sei40kr/zsh-fast-alias-tips/pull/23)).
* Build binary for Apple Silicon ([#26](https://github.com/sei40kr/zsh-fast-alias-tips/pull/26)).
* Zsh plugin file renamed to match the plugin name ([#25](https://github.com/sei40kr/zsh-fast-alias-tips/pull/25)).
* It's possible to exclude some unwanted aliases with `ZSH_FAST_ALIAS_TIPS_EXCLUDES` ([#24](https://github.com/sei40kr/zsh-fast-alias-tips/issues/24)).
* Some code refactoring to improve readability. For example - `Abbr` renamed to `Expanded`.
