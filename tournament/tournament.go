package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type team struct {
	name   string
	played int
	won    int
	drawn  int
	lost   int
	points int
}

// Tally the results of a small football competition.
func Tally(r io.Reader, w io.Writer) error {
	score := map[string]team{}

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		match := strings.Split(line, ";")

		if len(match) != 3 {
			return fmt.Errorf("invalid row: %s", line)
		}

		first := score[match[0]]
		second := score[match[1]]
		first.name = match[0]
		second.name = match[1]
		first.played++
		second.played++
		switch match[2] {
		case "win":
			first.points += 3
			first.won++
			second.lost++
		case "draw":
			first.points++
			second.points++
			first.drawn++
			second.drawn++
		case "loss":
			first.lost++
			second.won++
			second.points += 3
		default:
			return errors.New("invalid result")
		}
		score[match[0]] = first
		score[match[1]] = second
	}
	teams := make([]team, 0, len(score))
	for _, k := range score {
		teams = append(teams, k)
	}

	sort.Slice(teams,
		func(i int, k int) bool {
			if teams[k].points == teams[i].points {
				return teams[i].name < teams[k].name
			}
			return teams[k].points < teams[i].points
		})

	fmt.Fprintln(w, "Team                           | MP |  W |  D |  L |  P")
	for _, t := range teams {
		fmt.Fprintf(w, "%-31s| %2d | %2d | %2d | %2d | %2d\n",
			t.name,
			t.played,
			t.won,
			t.drawn,
			t.lost,
			t.points)
	}
	return nil
}
