#!/usr/bin/env bash

reverse_complement () {
  local SEQ=$1
  echo $(echo $SEQ | rev | tr 'ACGT' 'TGCA')
}

DNA=""

if [ -p /dev/stdin ]; then
  while IFS= read line; do
    DNA="${DNA}${line/\n}"
  done
else
  if [ -f "$1" ]; then
    DNA=$(cat $1)
  else
    echo "[ERROR] No input provided." >&2
    exit 1
  fi
fi

echo $(reverse_complement $DNA)
