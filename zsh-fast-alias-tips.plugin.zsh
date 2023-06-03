__fast_alias_tips__PLUGIN_DIR=${0:a:h}

if [[ -L ${0:a} ]]; then
  __fast_alias_tips__PLUGIN_DIR=$(readlink ${0:a})
  __fast_alias_tips__PLUGIN_DIR=${__fast_alias_tips__PLUGIN_DIR:h}
fi

: ${ZSH_FAST_ALIAS_TIPS_PREFIX:="💡 $(tput bold)"}
: ${ZSH_FAST_ALIAS_TIPS_SUFFIX:="$(tput sgr0)"}

__fast_alias_tips_preexec() {
    local CMD="$1"
    local CMD_EXPANDED="$2"

    local first="$(cut -d' ' -f1 <<<"$CMD")"

    local suggested="$(alias | exec ${__fast_alias_tips__PLUGIN_DIR}/zsh-alias-matcher "$CMD_EXPANDED")"
    if [[ "$suggested" == '' ]]; then
        return
    fi

    local suggested_first="$(cut -d' ' -f1 <<<"$suggested")"
    if [[ "$suggested_first" == "$first" ]]; then
        return
    fi

    if [[ ${#suggested} -gt ${#CMD} ]]; then
        return
    fi

    echo "${ZSH_FAST_ALIAS_TIPS_PREFIX}${suggested}${ZSH_FAST_ALIAS_TIPS_SUFFIX}"
}

autoload -Uz add-zsh-hook
add-zsh-hook preexec  __fast_alias_tips_preexec
