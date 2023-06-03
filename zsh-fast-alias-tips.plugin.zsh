: ${ZSH_FAST_ALIAS_TIPS_PREFIX:="💡 $(tput bold)"}
: ${ZSH_FAST_ALIAS_TIPS_SUFFIX:="$(tput sgr0)"}

__fast_alias_tips_preexec() {
    local cmd="$1"
    local cmd_expanded="$2"

    local first="$(cut -d' ' -f1 <<<"$cmd")"

    local suggested="$(alias | "zsh-alias-matcher" "$cmd_expanded")"
    if [[ "$suggested" == '' ]]; then
        return
    fi

    local suggested_first="$(cut -d' ' -f1 <<<"$suggested")"
    if [[ "$suggested_first" == "$first" ]]; then
        return
    fi

    echo "${ZSH_FAST_ALIAS_TIPS_PREFIX}${suggested}${ZSH_FAST_ALIAS_TIPS_SUFFIX}"
}

autoload -Uz add-zsh-hook
add-zsh-hook preexec  __fast_alias_tips_preexec
