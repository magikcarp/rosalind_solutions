#!/usr/bin/env sh
# Counts the number of A, C, G, T in a sequence
# Usage: ./dna.sh <file>
# Usage: <sequence> | ./dna.sh

SEQ=""

if [ -p /dev/stdin ]; then
  while IFS= read line; do
    SEQ="${SEQ}${line/\n}"
  done
else
  if [ -f "$1" ]; then
    SEQ=$(cat $1)
  else
    echo "[ERROR] No input provided." >&2
    exit 1
  fi
fi

A_COUNT=$(echo $SEQ | grep -o A | wc -l | tr -d ' ')
C_COUNT=$(echo $SEQ | grep -o C | wc -l | tr -d ' ')
G_COUNT=$(echo $SEQ | grep -o G | wc -l | tr -d ' ')
T_COUNT=$(echo $SEQ | grep -o T | wc -l | tr -d ' ')

echo "$A_COUNT $C_COUNT $G_COUNT $T_COUNT"
