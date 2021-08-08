# Sudoku

This is my take of [sudoku-solver](https://leetcode.com/problems/sudoku-solver/)

## Design

This uses a standard DFS with backtracking. Constant space (input board is modified).

DISCLAIMER: this solution was fixed to n=9. It would take a bit of refactoring to support other sizes.