package dir

import (
	"os"
	"testing"
	"time"
)

func TestLastModifiedFileReturnsCorrectOrder(t *testing.T) {
	const testDir = "tmp_cases"

	cases := []struct {
		files []string
		want  string
	}{
		{[]string{"foo.txt", "bar.txt", "baz.txt"}, "baz.txt"},
		{[]string{"baz.txt", "bar.txt", "foo.txt"}, "foo.txt"},
	}

	for _, c := range cases {
		// setup
		_ = os.Mkdir(testDir, 0777)
		for _, filename := range c.files {
			_, _ = os.Create(testDir + "/" + filename)
			time.Sleep(1 * time.Second)
		}
		// test
		got := LastModifiedFile(testDir + "/*.txt")
		if got != testDir+"/"+c.want {
			t.Errorf("Wanted %q, got %q", c.want, got)
		}
		// cleanup
		err := os.RemoveAll(testDir)
		if err != nil {
			t.Errorf("%q", err)
		}
	}

}
