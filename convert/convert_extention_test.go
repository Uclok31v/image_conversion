package convert

import (
	"testing"
)

func Test_Execute(t *testing.T) {
	cases := []struct {
		name          string
		directoryPath string
		before        string
		after         string
		expected      int
	}{
		{name: "jpg->png", directoryPath: "../testdata/jpg", before: "jpg", after: "png", expected: 0},
		{name: "png->jpg", directoryPath: "../testdata/png", before: "png", after: "jpg", expected: 0},
		{name: "Fail to encode", directoryPath: "../testdata/jpg", before: "jpg", after: "gif", expected: 1},
	}

	for _, c := range cases {
		t.Helper()
		t.Run(c.name, func(t *testing.T) {
			cli := Cli{c.directoryPath, &c.before, &c.after}
			actual := cli.Execute()
			if c.expected != actual {
				t.Errorf("want %d, actual %d", c.expected, actual)
			}
		})
	}
}

func Test_Convert(t *testing.T) {
	cases := []struct {
		name     string
		filePath string
		after    string
		expected int
	}{
		{name: "jpg->png", filePath: "../testdata/jpg/gopher.jpg", after: "png", expected: 0},
		{name: "png->jpg", filePath: "../testdata/png/gopher.png", after: "jpg", expected: 0},
		{name: "Fail to no file", filePath: "", after: "png", expected: 1},
		{name: "Fail to decode", filePath: "../testdata/empty", after: "png", expected: 1},
		{name: "Fail to encode", filePath: "../testdata/jpg/gopher.jpg", after: "gif", expected: 1},
	}

	for _, c := range cases {
		t.Helper()
		t.Run(c.name, func(t *testing.T) {
			err := convert(c.filePath, c.after)
			var actual int
			if err == nil {
				actual = 0
			} else {
				actual = 1
			}
			if c.expected != actual {
				t.Errorf("want %d, actual %d", c.expected, actual)
			}
		})
	}
}
