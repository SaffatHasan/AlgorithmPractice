# Best Slices ![go](https://github.com/SaffatHasan/BestSlices/actions/workflows/go.yml/badge.svg) ![golangci-lint](https://github.com/SaffatHasan/BestSlices/actions/workflows/golangci-lint.yml/badge.svg)

## Problem Statement

You are given a list of pizza slices of `m` size each such
that there are no duplicates (all pizza sizes are unique).

Determine the largest sum of pizza slices you can get
where you and a friend take turns selecting a slice.

You can pick a slice at either end of the list but your
friend always picks the largest of the two ends.

```
input: [4, 5, 2, 3]
output: 8
explanation:
    You pick 3
    Your friend picks 4
    You pick 5
    Your friend picks 2
```