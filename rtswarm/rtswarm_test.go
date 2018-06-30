package rtswarm_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/RTradeLtd/Temporal/rtswarm"
)

func TestRTSwarmUpload(t *testing.T) {
	sm, err := rtswarm.NewSwarmManager()
	if err != nil {
		t.Fatal(err)
	}
	resp, err := sm.Upload("/tmp/test_file", "", true)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}

func TestRTSwarmRaw(t *testing.T) {
	fmt.Println(1)
	sm, err := rtswarm.NewSwarmManager()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(2)
	file, err := os.Create("/tmp/test_file")
	defer file.Close()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(3)
	_, err = file.WriteString("wow such data")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(4)
	file, err = os.Open("/tmp/test_file")
	defer file.Close()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(5)
	resp, err := sm.UploadRaw(file, "", false)
	if err != nil {
		t.Fatal(err)
	}
	manifest, encrypted, err := sm.DownloadManifest(resp)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(manifest)
	fmt.Println(encrypted)
}
