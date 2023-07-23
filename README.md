# Alem01 - Tetris-Optimizer

## Description
Tetris-Optimizer is an algorithmic task that aims to fit tetris pieces into the smallest square possible.

## Instruction
* The program requires a single parameter, which is a file containing a list of Tetriminos to assemble. The file must adhere to a specific format, where each Tetrimino must fit precisely in a 4x4 grid, and they are separated by a newline.
* If there are more than one commandline parameters given, the program will ask for its usage and terminate gracefully. There are at least one Tetrimino, but the upper bound is not set explicitly. However, I advise limiting the number of Tetriminos to 26 or fewer for experimentation and playing purposes.

## Validation process
* One piece is precisely 4 lines of 4 characters, each followed by a new line.
* A Tetrimino is a classic piece of Tetris composed of 4 blocks.
* Each character must be either a '#' character or '.' character.
* Each block of a Tetrimino must touch at least one other block on any of his 4 sides(up, down, left and right). This process is illustrated well with visualizations in the Beth Lock's article below.
* There are 19 valid pieces. If the input piece does not match one of them. Terminate the process.


## Algorithm
- The initial square dimension, denoted as dim, is calculated as dim = sqrt(n * 4), where n represents the number of Tetrimino pieces.

- The algorithm employed in this process follows a classical backtracking approach. For each calculated dimension, we recursively fit all combinations of Tetrimino pieces within the square board.
## Example
![Valid Input](https://01.alem.school/git/yeakbay/tetris-optimizer/workfiles/ex1.png?raw=true)

```bash
    go run main.go sample.txt
```
![Valid Result](https://01.alem.school/git/yeakbay/tetris-optimizer/workfiles/res1.png?raw=true)

## Test
In cases directory, there are tons of cases.
To run test.
```bash
    cd algorithm
    go test
```

## Tech
Go version 1.20
Used only standard library
## References
[Tetrimino background](https://en.wikipedia.org/wiki/Tetromino)
[42 School student Beth Locke's article](https://medium.com/@bethnenniger/fillit-solving-for-the-smallest-square-of-tetrominos-c6316004f909)
[42 School student Phat Truong's article](https://medium.com/@phtruong42/42-project-fillit-ffd0ce54223e)
[Sudoku solver' backtracking algorithm](https://www-geeksforgeeks-org.cdn.ampproject.org/c/s/www.geeksforgeeks.org/sudoku-backtracking-7/amp/)

# Author
[@yeakbay](https://01.alem.school/git/yeakbay)