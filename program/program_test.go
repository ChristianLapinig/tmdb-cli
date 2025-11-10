package program

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/ChristianLapinig/tmdb-cli/constants"
)

func TestProgram(t *testing.T) {
	t.Run("exits when access token is missing", func(t *testing.T) {
		if os.Getenv("TEST_PROGRAM_FATAL") == "1" {
			Program("")
			return
		}

		cmd := exec.Command(os.Args[0], "-test.run=TestProgram")
		cmd.Env = append(os.Environ(), "TEST_PROGRAM_FATAL=1")
		actual, err := cmd.CombinedOutput()

		if exitErr, ok := err.(*exec.ExitError); !ok || exitErr == nil {
			t.Fatalf("FAILE: expected process to exit with no error go %v", exitErr)
		}

		expected := fmt.Sprintf("%s\n", constants.AccessTokenRequired)
		fmt.Println(string(actual))
		if string(actual) != expected {
			t.Fatalf("FAILED: expected %s, got %s", expected, actual)
		}
	})
}
