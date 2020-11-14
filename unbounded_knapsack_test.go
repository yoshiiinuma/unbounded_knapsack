
package knapsack

import (
  "testing"
  . "github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
  var K int
  var W []int
  var V []int

  K = 10
  V = []int{1, 30}
  W = []int{1, 5}
  Equal(t, 60, Knapsack(K, V, W))

  K = 9
  V = []int{1, 30}
  W = []int{1, 5}
  Equal(t, 34, Knapsack(K, V, W))

  K = 10
  V = []int{1, 4, 6}
  W = []int{1, 3, 4}
  Equal(t, 14, Knapsack(K, V, W))

  K = 8
  V = []int{10, 40, 50, 70}
  W = []int{1, 3, 4, 5}
  Equal(t, 110, Knapsack(K, V, W))

  K = 7
  V = []int{4, 5, 3}
  W = []int{3, 4, 2}
  Equal(t, 10, Knapsack(K, V, W))
}
