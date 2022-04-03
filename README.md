# learnGo
[![Go Report Card](https://goreportcard.com/badge/github.com/samiam2013/learnGo)](https://goreportcard.com/report/github.com/samiam2013/learnGo) [![License: CC BY-SA 4.0](https://img.shields.io/badge/License-CC_BY--SA_4.0-lightgrey.svg)](https://creativecommons.org/licenses/by-sa/4.0/)

[Simple exercises from adriann.github.io/programming_problems.html](https://adriann.github.io/programming_problems.html) (CC-BY-SA)

There are no external dependencies, this is all pure [Go](https://www.golang.org)

There's three folders with the exercises so far: [Elementary](elementary/), [ListsAndStrings](listsandstrings/) and [Intermediate](intermediate/) Each folder has it's own readme with a listing of the exercies and a status for each one in the form of emojis

Oh, and it has [a heuristic sudoku solver](offscript/sudoku.go) in /offscript

![gopherLearn](offscript/witch-learning.svg)

[[from egonelbre/gophers/](https://github.com/egonelbre/gophers/)]

## Next Groups of Exercises from the Set (not yet started)
### Advanced Exercises

1. Given two strings, write a program that efficiently finds the longest common subsequence.
2. Given an array with numbers, write a program that efficiently answers queries of the form: “Which is the nearest larger value for the number at position i?”, where distance is the difference in array indices. For example in the array [1,4,3,2,5,7], the nearest larger value for 4 is 5. After linear time preprocessing you should be able to answer queries in constant time.
3. Given two strings, write a program that outputs the shortest sequence of character insertions and deletions that turn one string into the other.
4. Write a function that multiplies two matrices together. Make it as efficient as you can and compare the performance to a polished linear algebra library for your language. You might want to read about Strassen’s algorithm and the effects CPU caches have. Try out different matrix layouts and see what happens.
5. Implement a van Emde Boas tree. Compare it with your previous search tree implementations.
6. Given a set of d-dimensional rectangular boxes, write a program that computes the volume of their union. Start with 2D and work your way up.

### GUI

1. Write a program that displays a bouncing ball.
2. Write a Memory game.
3. Write a Tetris clone

### Open Ended

1. Write a program that plays Hangman as good as possible. For example you can use a large dictionary like this and select the letter that excludes most words that are still possible solutions. Try to make the program as efficient as possible, i.e. don’t scan the whole dictionary in every turn.
2. Write a program that plays Rock, Paper, Scissors better than random against a human. Try to exploit that humans are very bad at generating random numbers.
3. Write a program that plays Battle Ship against human opponents. It takes coordinates as input and outputs whether that was a hit or not and its own shot’s coordinates.
