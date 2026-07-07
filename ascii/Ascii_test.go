package ascii

import (
	"os"
	"reflect"
	"testing"
)

func sampleBanner() map[rune][]string {
	return map[rune][]string{
		'A': {
			"A1", "A2", "A3", "A4",
			"A5", "A6", "A7", "A8",
		},
		'B': {
			"B1", "B2", "B3", "B4",
			"B5", "B6", "B7", "B8",
		},
		' ': {
			"  ", "  ", "  ", "  ",
			"  ", "  ", "  ", "  ",
		},
	}
}

func TestRenderLine(t *testing.T) {
	banner := sampleBanner()

	tests := []struct {
		name string
		text string
		want []string
	}{
		{
			name: "empty string",
			text: "",
			want: []string{"", "", "", "", "", "", "", ""},
		},
		{
			name: "single character",
			text: "A",
			want: []string{
				"A1", "A2", "A3", "A4",
				"A5", "A6", "A7", "A8",
			},
		},
		{
			name: "multiple characters",
			text: "AB",
			want: []string{
				"A1B1",
				"A2B2",
				"A3B3",
				"A4B4",
				"A5B5",
				"A6B6",
				"A7B7",
				"A8B8",
			},
		},
		{
			name: "text with spaces",
			text: "A B",
			want: []string{
				"A1  B1",
				"A2  B2",
				"A3  B3",
				"A4  B4",
				"A5  B5",
				"A6  B6",
				"A7  B7",
				"A8  B8",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RenderLine(tt.text, banner)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v\nwant %v", got, tt.want)
			}
		})
	}
}

func TestGenerateArt(t *testing.T) {
	banner := sampleBanner()

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "empty input",
			input: "",
			want:  "",
		},
		{
			name:  "single character",
			input: "A",
			want:  "A1\nA2\nA3\nA4\nA5\nA6\nA7\nA8\n",
		},
		{
			name:  "multiple characters",
			input: "AB",
			want:  "A1B1\nA2B2\nA3B3\nA4B4\nA5B5\nA6B6\nA7B7\nA8B8\n",
		},
		{
			name:  "multiple lines",
			input: "A\nB",
			want: "A1\nA2\nA3\nA4\nA5\nA6\nA7\nA8\n" +
				"B1\nB2\nB3\nB4\nB5\nB6\nB7\nB8\n",
		},
		{
			name:  "blank line",
			input: "A\n\nB",
			want: "A1\nA2\nA3\nA4\nA5\nA6\nA7\nA8\n\n" +
				"B1\nB2\nB3\nB4\nB5\nB6\nB7\nB8\n",
		},
		{
			name:  "windows newline",
			input: "A\r\nB",
			want: "A1\nA2\nA3\nA4\nA5\nA6\nA7\nA8\n" +
				"B1\nB2\nB3\nB4\nB5\nB6\nB7\nB8\n",
		},
		{
			name:  "only newline",
			input: "\n",
			want:  "\n\n",
		},
		{
			name:  "leading newline",
			input: "\nA",
			want: "\n" +
				"A1\nA2\nA3\nA4\nA5\nA6\nA7\nA8\n",
		},
		{
			name:  "trailing newline",
			input: "A\n",
			want:  "A1\nA2\nA3\nA4\nA5\nA6\nA7\nA8\n\n",
		},
		{
			name:  "invalid ascii",
			input: "A\tB",
			want:  "Invalid Ascii character",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateArt(tt.input, banner)

			if got != tt.want {
				t.Errorf("got:\n%q\nwant:\n%q", got, tt.want)
			}
		})
	}
}

func TestLoadBanner(t *testing.T) {
	tmp := t.TempDir()

	valid := tmp + "/banner.txt"
	empty := tmp + "/empty.txt"
	missing := tmp + "/missing.txt"

	content := ""

	content += "\n"
	content += "s1\ns2\ns3\ns4\ns5\ns6\ns7\ns8\n"

	content += "\n"
	content += "e1\ne2\ne3\ne4\ne5\ne6\ne7\ne8\n"

	if err := os.WriteFile(valid, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(empty, []byte(""), 0644); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		file    string
		wantErr bool
	}{
		{
			name:    "valid banner",
			file:    valid,
			wantErr: false,
		},
		{
			name:    "missing file",
			file:    missing,
			wantErr: true,
		},
		{
			name:    "empty file",
			file:    empty,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadBanner(tt.file)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected an error")
				}
				return
			}

			if err != nil {
				t.Fatal(err)
			}

			switch tt.name {
			case "valid banner":
				if len(got) != 2 {
					t.Fatalf("expected 2 characters, got %d", len(got))
				}

				if got[' '][0] != "s1" {
					t.Errorf("space character loaded incorrectly")
				}

				if got['!'][7] != "e8" {
					t.Errorf("! character loaded incorrectly")
				}

			case "empty file":
				if len(got) != 0 {
					t.Errorf("expected an empty banner map")
				}
			}
		})
	}
}
