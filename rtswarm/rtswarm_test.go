package rtswarm_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/RTradeLtd/Temporal/rtswarm"
)

func TestRTSwarm(t *testing.T) {
	sm, err := rtswarm.NewSwarmManager()
	if err != nil {
		t.Fatal(err)
	}
	file, err := os.Create("/tmp/test_file")
	if err != nil {
		t.Fatal(err)
	}
	_, err = file.WriteString("wow such data")
	if err != nil {
		t.Fatal(err)
	}
	err = file.Close()
	if err != nil {
		t.Fatal(err)
	}
	file, err = os.Open("/tmp/test_file")
	if err != nil {
		t.Fatal(err)
	}
	resp, err := sm.UploadRaw(file, "", true)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}
