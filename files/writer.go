package files

import (
	"bufio"
	"fmt"
	"os"
)

func WriteFile(dstFile, content string) (err error) {
	file, err := os.Create(dstFile)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	_, err = fmt.Fprint(writer, content)
	if err != nil {
		return err
	}

	return nil
}
