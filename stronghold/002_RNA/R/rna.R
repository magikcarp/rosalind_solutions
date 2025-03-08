#!/usr/bin/env Rscript
# rna.R
# Translates DNA sequence
seq <- readLines(file("stdin"))
seq <- gsub("T", "U", toupper(gsub("%s+", "", seq)))
writeLines(seq)
