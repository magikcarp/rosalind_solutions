#!/usr/bin/env sh
# Translates DNA to RNA
# Usage: ./rna.sh <file>
# Usage: <sequence> | ./rna.sh

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

RNA=$(echo $DNA | sed "s/T/U/g")
echo $RNA
