package files

import (
	"bufio"
	"fmt"
	"os"
)

func FindReplaceFile(src, old, new string) (occ int, lines []int, err error) {
	file, err := os.Open(src)
	if err != nil {
		return occ, lines, err
	}
	defer file.Close()

	updatedText := ""

	scanner := bufio.NewScanner(file)
	linesCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		found, res, lineOcc := ProcessLine(line, old, new)

		if found {
			lines = append(lines, linesCount)
		}

		updatedText += res + "\n"
		occ += lineOcc

		linesCount++
	}

	if linesCount == 0 {
		err = fmt.Errorf("File %v is empty", src)
		return occ, lines, err
	}

	err = WriteFile("updated-"+src, updatedText)

	return occ, lines, err
}
