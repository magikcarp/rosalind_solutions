module Main where

rabbitReproduction :: Int -> Int -> Int
rabbitReproduction 1 _ = 1
rabbitReproduction 2 _ = 1
rabbitReproduction n k =
  (rabbitReproduction (n-1) k) + ((rabbitReproduction (n-2) k) * k)

main :: IO ()
main = do
  putStrLn $ show $ rabbitReproduction 30 4
