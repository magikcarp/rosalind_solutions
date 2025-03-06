module Main where

import Data.Map (Map, fromListWith, elems)

data Nucleotide = A | C | G | T deriving (Eq, Ord, Show)

isInvalidNucleotide :: Char -> Bool
isInvalidNucleotide x = x `notElem` "ACGT"

toNucleotide :: Char -> Nucleotide
toNucleotide x
  | x == 'A' = A
  | x == 'C' = C
  | x == 'G' = G
  | x == 'T' = T

strToDNA :: String -> [Nucleotide]
strToDNA = map toNucleotide

countNucs :: [Nucleotide] -> Map Nucleotide Int
countNucs seq = fromListWith (+) [(c, 1) | c <- seq]

main :: IO ()
main = do
    seq <- getLine
    let out = elems $ countNucs $ strToDNA seq
    putStrLn $ unwords (map show out)
