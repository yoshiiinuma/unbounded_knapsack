# Unbounded Knapsack

Given a set of different N items, each one with an associated value and weight, determine the maximum value of items that can be packed into a knapsack of capacity K.
The values and weights of N items are given as integer arrays; V { v1, v2, ..., vn }, W { w1, w2, ..., wn }, respectively. Here we are allowed to use unlimited number of instances of an item.

## Constraints

```
0 < N <= 100
0 < V[i], W[i] <= 100
0 < K <= 10000
```

## Examples

```
INPUT:

  N: 2
  K: 10 
  V: [1, 30]
  W: [1, 5]

OUTPUT:

  60
```

```
INPUT:

  N: 2
  K: 9 
  V: [1, 30]
  W: [1, 5]

OUTPUT:

  34
```

```
INPUT:

  N: 4
  K: 8
  V: [10, 40, 50, 70]
  W: [1, 3, 4, 5]

OUTPUT:

  110
```

## Approach

This problem can be solved with Dynamic Programming.

Suppose DP[i][j] represents the maximum value that can be held by a knapsack of capacity j, then: 

```
i) When j < W[i]

  DP[i+1][j] = DP[i][j]

ii) When j >= W[i]

  DP[i+1][j] = MAX(DP[i][j], D[i][j - W[i]] + V[i])
```

## Complexity Analysis

* Time Complexity: O(N*K)
* Space Complexity: O(N*K)
