package main

import (
	"testing"
)

func TestExtractingIPs(t *testing.T) {
	logContents := `2015-01-11 00:52:06 127.0.0.1 Unable to set default name
2015-08-10 00:52:08 127.0.0.1 Creating embedded database
2015-09-15 00:53:01 207.121.21.21 update user
`

	expected := `127.0.0.1
127.0.0.1
207.121.21.21
`
	test(t, expected, logContents, 1)
}

func TestExtractingDateAndTime(t *testing.T) {
	logContents := `2015-01-11 00:52:06 127.0.0.1 Unable to set default name
2015-08-10 00:52:08 127.0.0.1 Creating embedded database
2015-09-15 00:53:01 207.121.21.21 update user
`

	expected := `2015-01-11 00:52:06
2015-08-10 00:52:08
2015-09-15 00:53:01
`
	test(t, expected, logContents, 0)
}

func TestExtractingText(t *testing.T) {
	logContents := `2015-01-11 00:52:06 127.0.0.1 Unable to set default name
2015-08-10 00:52:08 127.0.0.1 Creating embedded database
2015-09-15 00:53:01 207.121.21.21 update user
`

	expected := `Unable to set default name
Creating embedded database
update user
`
	test(t, expected, logContents, 2)
}

func test(t *testing.T, expected, logContents string, column uint8) {
	found := ExtractColumn(logContents, column)

	if found != expected {
		t.Errorf("Expected\n---\n%s\n---\nbut found\n---\n%s\n---\n", expected, found)
	}
}
