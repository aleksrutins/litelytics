#!/usr/bin/env bash
source .env
reverse_array(){
  local -n _source_array_ref="${1}"
  local -n _destination_array_ref="${2}"

  for ((_index=${#_source_array_ref[@]}-1; _index>=0; _index--)); do
    _destination_array_ref+=("${_source_array_ref[$_index]}")
  done
}


migrations=(migrations/*)
actual_migrations=()

filename=""
if [ "$1" == "up" ]; then
    filename="up.sql"
    actual_migrations=(${migrations[@]})
elif [ "$1" == "down" ]; then
    filename="down.sql"
    reverse_array 'migrations' 'actual_migrations'
elif [ "$1" == "revert" ]; then
    filename="down.sql"
    reverse_array 'migrations' 'actual_migrations'
    actual_migrations=($actual_migrations)
else
    printf "\e[41m ERROR \e[0m Please provide either \e[1mup\e[0m, \e[1mrevert\e[0m or \e[1mdown\e[0m as a command.\n"
    exit 1
fi

for migration in ${actual_migrations[@]}; do
    printf "\e[s\e[46m RUNNING \e[0m $migration\n"
    if psql $DATABASE_URL -f "$migration/$filename"; then
        printf "\e[u\e[2K"
        printf "\e[42m SUCCESS "
    else
        printf "\e[41m ERROR "
    fi
    printf "\e[0m $migration\n"
done