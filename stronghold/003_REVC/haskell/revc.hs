module Main where

complement :: Char -> Char
complement n
  | n == 'A' = 'T'
  | n == 'C' = 'G'
  | n == 'G' = 'C'
  | n == 'T' = 'A'
  | otherwise = n

reverse_complement :: [Char] -> [Char]
reverse_complement seq = map complement $ reverse seq

main :: IO ()
main = do
  seq <- getLine
  let out = reverse_complement seq
  putStrLn out
  
