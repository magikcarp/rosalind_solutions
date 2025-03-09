#!/usr/bin/env Rscript
# lia.r

nums <- lapply(strsplit(readLines(file("stdin")), " ")[[1]], as.numeric)
result <- pbinom(nums[[2]] - 1, 2 ^ nums[[1]], 0.25, lower.tail = FALSE)
writeLines(sprintf("%.3f", result))
