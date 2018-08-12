. ~/dotfiles/common.conf

# プロンプトにgitのブランチ名を表示する
. $HOME/dotfiles/.git-prompt.sh

# プロンプトの表示を変更
export PS1='\[\033[01;33m\][\t]\[\033[00m\] \[\033[01;32m\]\u\[\033[00m\]:\[\033[01;34m\]\w\[\033[00m\] \[\033[01;35m\]$(__git_ps1 "(%s)")\[\033[00m\]\$ '

