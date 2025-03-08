#!/usr/bin/env Rscript
# dna.R
# Prints nucleotide counts of a provided DNA sequence through stdin

seq <- readLines(file("stdin"))
nc <- table(strsplit(seq, "")[[1]]) # short for nucleotide counts
writeLines(sprintf("%d %d %d %d", nc[[1]], nc[[2]], nc[[3]], nc[[4]]))
