multiples :: Int -> Int
multiples x
  | x < 1000 && (mod x 3 == 0 || mod x 5 == 0)
    = x + multiples(x+1)
  | x < 1000 = multiples(x+1)
  | otherwise = 0
