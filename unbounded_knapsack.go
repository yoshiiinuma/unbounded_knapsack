
package knapsack

import (
  "fmt"
)

const DEBUG = true

func showDP(DP [][]int) {
  N := len(DP)
  for i := 0; i < N; i++ {
    fmt.Printf("%v\n", DP[i])
  }
}

func Knapsack(K int, V []int, W []int) int {
  if DEBUG {
    fmt.Println("-------------------------------------------------")
    fmt.Printf(" K: %d, V: %v, W: %v\n", K, V, W)
    fmt.Println("-------------------------------------------------")
  }
  N := len(V) + 1
  DP := make([][]int, N)
  for i := 0; i < N; i++ {
    DP[i] = make([]int, K + 1)
  }
  return knapsack(DP, K, V, W)
}

func knapsack(DP [][]int, K int, V []int, W []int) int {
  N := len(DP)
  for i := 1; i < N; i++ {
    for j := 1; j < K + 1; j++ {
      DP[i][j] = DP[i-1][j]
      if j >= W[i-1] {
        if DP[i][j] < DP[i][j - W[i-1]] + V[i-1] {
          DP[i][j] = DP[i][j - W[i-1]] + V[i-1]
        }
      }
    }
  }
  if DEBUG { showDP(DP) }
  return DP[N-1][K]
}
