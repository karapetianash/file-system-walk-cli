package main

import (
	"os"
	"strings"

	"testing"
)

func TestFilterOut(t *testing.T) {
	testCases := []struct {
		name     string
		file     string
		ext      string
		minSize  int64
		expected bool
	}{
		{"FilterNoExtension", "testdata/dir.log", "", 0, false},
		{"FilterExtensionMatch", "testdata/dir.log", ".log", 0, false},
		{"FilterExtensionNoMatch", "testdata/dir.log", ".sh", 0, true},
		{"FilterExtensionSizeMatch", "testdata/dir.log", ".log", 10, false},
		{"FilterExtensionSizeNoMatch", "testdata/dir.log", ".log", 20, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			info, err := os.Stat(tc.file)
			if err != nil {
				t.Fatal(err)
			}

			extList := strings.Fields(tc.ext)
			f := filterOut(tc.file, extList, tc.minSize, info)

			if f != tc.expected {
				t.Errorf("Expected '%t', got '%t' instead\n", tc.expected, f)
			}
		})
	}
}

func TestListContainsExt(t *testing.T) {
	testCases := []struct {
		name     string
		list     []string
		ext      string
		expected bool
	}{
		{"ListContain", []string{".log", ".rar", ".png", ".jpeg", "xlsb", ".txt"}, ".log", true},
		{"ListNoContain", []string{".log", ".rar", ".png", ".jpeg", "xlsb", ".txt"}, ".zip", false},
		{"ListEmptyNoContain", []string{}, ".zip", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := listContainsExt(tc.list, tc.ext)

			if tc.expected != res {
				t.Errorf("Expected %t, got %t instead.", tc.expected, res)
			}
		})
	}
}
