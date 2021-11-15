package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Results holds the different values per team
type Results struct {
	Matches int
	Wins    int
	Draws   int
	Losses  int
	Points  int
}

// Track stores the results of each Team
type Track map[string]*Results

//Tally reads the file that describes the playing teams and their respective success and writes the formated
// Results to a table
func Tally(reader io.Reader, writer io.Writer) error {
	TrackTournament := make(Track)
	scanner := bufio.NewScanner(reader)
	scanner.Split(CrunchSplitFunc)

	for scanner.Scan() {
		portion := scanner.Text()
		line := strings.Split(portion, ";")
		if line[0] == "" || line[0][0] == 35 {
			continue
		}
		if len(line) < 3 {
			return errors.New("incorrect input")
		}
		result, err := TallyResults(line[0], line[1], line[2], TrackTournament)
		if err != nil {
			return err
		} else {
			TrackTournament = result
		}
	}

	stringToBuffer := TallyString(TrackTournament)
	writer.Write([]byte(stringToBuffer))
	return nil
}

// TallyString takes the results, formates them and writes them to the buffer
func TallyString(results Track) (stringPrint string) {
	stringPrint = `Team                           | MP |  W |  D |  L |  P
`

	keys := make([]string, 0, len(results))
	for key := range results {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		pointsA := results[keys[i]].Points
		pointsB := results[keys[j]].Points
		if pointsA == pointsB {
			return keys[i] < keys[j]
		} else {
			return pointsA > pointsB
		}
	})

	for _, key := range keys {
		v := results[key]
		clear := fmt.Sprintf("%-31s", key)
		stringPrint += fmt.Sprintf("%s|  %d |  %d |  %d |  %d |  %d\n", clear, v.Matches, v.Wins, v.Draws, v.Losses, v.Points)
	}

	return stringPrint
}

// TallyResults takes in the TrackTournament map and assigns the correct results
func TallyResults(first string, second string, result string, results Track) (newResults Track, incorrectInput error) {

	holder := results
	if _, ok := holder[first]; !ok {
		holder[first] = &Results{Matches: 0, Wins: 0, Draws: 0, Losses: 0, Points: 0}
	}

	if _, ok := holder[second]; !ok {
		holder[second] = &Results{Matches: 0, Wins: 0, Draws: 0, Losses: 0, Points: 0}
	}

	switch result {
	case "win":
		{
			holder[first].Wins += 1
			holder[first].Points += 3
			holder[second].Losses += 1
		}
	case "loss":
		{
			holder[first].Losses += 1
			holder[second].Wins += 1
			holder[second].Points += 3
		}
	case "draw":
		{
			holder[first].Draws += 1
			holder[first].Points += 1
			holder[second].Draws += 1
			holder[second].Points += 1
		}
	default:
		{
			return nil, errors.New("incorrectInput")
		}
	}

	holder[first].Matches += 1
	holder[second].Matches += 1

	return holder, nil
}

// CrunchSplitFunc splits the input on every semicolon
func CrunchSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {

	// Return nothing if at end of file and no data passed
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// Find the index of the input of a newline followed by a
	// pound sign.
	if i := strings.Index(string(data), "\n"); i >= 0 {
		return i + 1, data[0:i], nil
	}

	// If at end of file with data return the data
	if atEOF {
		return len(data), data, nil
	}
	return
}
