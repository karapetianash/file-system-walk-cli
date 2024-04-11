package main

import (
	"os"
	"strings"
	"time"

	"testing"
)

func TestFilterOut(t *testing.T) {
	testCases := []struct {
		name     string
		file     string
		ext      string
		minSize  int64
		modSince string
		expected bool
	}{
		{"FilterNoExtension", "testdata/dir.log", "", 0, "", false},
		{"FilterExtensionMatch", "testdata/dir.log", ".log", 0, "", false},
		{"FilterExtensionNoMatch", "testdata/dir.log", ".sh", 0, "", true},
		{"FilterExtensionSizeMatch", "testdata/dir.log", ".log", 10, "", false},
		{"FilterExtensionSizeNoMatch", "testdata/dir.log", ".log", 20, "", true},
		{"FilterExtensionTimeMatch", "testdata/dir.log", ".log", 0, "01 Apr 24 12:00 +0300", false},
		{"FilterExtensionTimeNoMatch", "testdata/dir.log", ".log", 0, "12 Apr 24 12:00 +0300", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			info, err := os.Stat(tc.file)
			if err != nil {
				t.Fatal(err)
			}

			extList := strings.Fields(tc.ext)

			var afterRFC822Z time.Time
			if tc.modSince != "" {
				afterRFC822Z, err = time.Parse(time.RFC822Z, tc.modSince)
				if err != nil {
					t.Fatal(err)
				}

			}

			f := filterOut(tc.file, extList, tc.minSize, afterRFC822Z, info)

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
