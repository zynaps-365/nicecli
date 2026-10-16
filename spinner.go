package nicecli

import (
	"fmt"
	"time"
)

func Spinner(message string, fn func() string) string {
	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	done := make(chan string)

	go func() {
		done <- fn()
	}()

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	frameIdx := 0
	frameShown := false

	for {
		select {
		case result := <-done:
			if frameShown {
				fmt.Print("\r\033[K")
			}
			return result
		case <-ticker.C:
			fmt.Printf("\r\033[35m%s\033[0m %s", frames[frameIdx], message)
			frameShown = true
			frameIdx = (frameIdx + 1) % len(frames)
		}
	}
}

func Success(format string, a ...interface{}) {
	fmt.Printf("\033[32m✔\033[0m %s\n", fmt.Sprintf(format, a...))
}
