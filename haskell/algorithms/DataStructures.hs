module DataStructures where

import           System.Random (getStdRandom, randomR)
import Data.Maybe (listToMaybe)
import           Data.Monoid ((<>))

-- List

-- Fisher Yates shuffle
-- Resources:
-- https://spin.atomicobject.com/2014/08/11/fisher-yates-shuffle-randomization-algorithm/
-- https://bost.ocks.org/mike/shuffle/
-- http://blog.codinghorror.com/the-danger-of-naivete/
-- Time O(n^2) :(( TODO: switch to array for O(n)
-- Space O(n)
shuffle :: [a] -> IO [a]
shuffle [] = return []
shuffle as = do
  rand <- getStdRandom $ randomR (0, length as - 1) -- [0, n)
  let (left, a:right) = splitAt rand as
  (a:) <$> shuffle (left <> right)

-- Reverse a list
-- Time O(n^2) :(( TODO: switch to where helper fn w/ [] accumulator for O(n)
-- Space O(n)
reverse' :: [a] -> [a]
reverse' [] = []
reverse' (a:as) = reverse' as <> [a]

-- Banker's Queue
data Queue a = Queue [a] [a]

emptyQ :: Queue a
emptyQ = Queue [] []

sizeQ :: Queue a -> Int
sizeQ (Queue xs ys) = foldl (\s l -> s + length l) 0 [xs, ys]

isEmptyQ :: Queue a -> Bool
isEmptyQ (Queue [] []) = True
isEmptyQ _ = False

enqueueQ :: a -> Queue a -> Queue a
enqueueQ a (Queue xs ys) = Queue xs $ [a] <> ys

dequeueQ :: Queue a -> (Maybe a, Queue a)
dequeueQ (Queue [] []) = (Nothing, emptyQ)
dequeueQ (Queue (x:xs) ys) = (Just x, Queue xs ys)
dequeueQ (Queue [] ys) = dequeueQ $ Queue (reverse' ys) []

data Stack a = Stack [a]

emptyS :: Stack a
emptyS = Stack []

isEmptyS :: Stack a -> Bool
isEmptyS (Stack []) = True
isEmptyS _ = False

popS :: Stack a -> (Maybe a, Stack a)
popS (Stack []) = (Nothing, emptyS)
popS (Stack (a:rest)) = (Just a, Stack rest)

pushS :: a -> Stack a -> Stack a
pushS a (Stack as) = Stack $ [a] <> as

data Tree a = Empty | Node a (Tree a) (Tree a)
