package adapter

import (
	"bufio"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"strings"
)

func WriteStringData(w *bufio.Writer, str string, closed *bool) {
	str = strings.TrimPrefix(str, "data: ")
	str = strings.TrimSuffix(str, "\r")
	str = strings.TrimSuffix(str, "\n")
	_, err := fmt.Fprintf(w, "data: %s\n\n", str)
	if err != nil {
		return
	}
	err = w.Flush()
	if err != nil {
		log.Errorf("Error while flushing: %v. Closing http connection.", err)
		if err.Error() == "connection closed" {
			*closed = true
		}
		return
	}
}
