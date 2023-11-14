package mazesolver

import (
	"testing"

	"github.com/joshdevelopsIRL/jeer"
)

func TestMazeSolver(t *testing.T) {
    maze := [][]string{
        {"x", "x", "x", "x", "x", "x", "x", "x", "x", "x", " ", "x", },
        {"x", " ", " ", " ", " ", " ", " ", " ", " ", "x", " ", "x", },
        {"x", " ", " ", " ", " ", " ", " ", " ", " ", "x", " ", "x", },
        {"x", " ", "x", "x", "x", "x", "x", "x", "x", "x", " ", "x", },
        {"x", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "x", },
        {"x", " ", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", },
    }
    start := Point{10, 0}
    end := Point{1, 5}
    path := []Point{
        start, {10, 1}, {10, 2}, {10, 3}, {10, 4}, {9, 4}, {8, 4},
        {7, 4}, {6, 4}, {5, 4}, {4, 4}, {3, 4}, {2, 4}, {1, 4}, end,
    }

    result := solve(maze, "x", start, end)

    jeer.Test[Point](t).IsList().Actual(result...).Expected(path...).Run("test 1")

    // ===========================================================================

    maze = [][]string{
        {"#", "#", "#", "#", "#", "#", "#", "#", "#"},
        {"#", " ", "#", " ", "#", " ", "#", " ", "#"},
        {" ", " ", "#", " ", "#", " ", " ", " ", "#"},
        {"#", " ", "#", " ", " ", " ", "#", " ", "#"},
        {"#", " ", "#", " ", "#", " ", "#", " ", "#"},
        {"#", " ", " ", " ", "#", " ", "#", " ", "#"},
        {"#", " ", "#", " ", "#", " ", "#", " ", " "},
        {"#", " ", "#", " ", "#", " ", "#", " ", "#"},
        {"#", "#", "#", "#", "#", "#", "#", "#", "#"},
    }
    start = Point{0, 2}
    end = Point{8, 6}
    path = []Point{
        start, {1, 2}, {1, 3}, {1, 4}, {1, 5}, {2, 5}, {3, 5},
        {3, 4}, {3, 3}, {4, 3}, {5, 3}, {5, 2}, {6, 2}, {7, 2},
        {7, 3}, {7, 4}, {7, 5}, {7, 6}, end,
    }

    result = solve(maze, "#", start, end)

    jeer.Test[Point](t).IsList().Actual(result...).Expected(path...).Run("test 2")
}