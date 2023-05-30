package tree

import (
	"math"
	"testing"
)

func TestDFS(t *testing.T) {

	t.Run("empty tree", func(t *testing.T) {
		tree := GenerateTree(10)

		singleResult := tree.SingleThreadSearch()
		tree.RemoveVisitors()

		multiResult := tree.MultiTaskThreadSearch(4, int(math.Pow(2, 10)), 4)
		if singleResult != multiResult {
			t.Error("result are not equal")
		}
	})

	t.Run("not empty tree", func(t *testing.T) {
		tree := GenerateTree(10)
		tree.GenerateSearchPlace(7)

		singleResult := tree.SingleThreadSearch()
		tree.RemoveVisitors()

		multiResult := tree.MultiTaskThreadSearch(4, int(math.Pow(2, 10)), 4)
		if singleResult != multiResult {
			t.Error("result are not equal")
		}
	})
}
