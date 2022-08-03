package atcoder

import (
	"io/ioutil"
	"strings"
	"testing"
)

// TODO: update REVEL_SESSION
func TestGenAtCoderContestTemplates(t *testing.T) {
	const contestID = "abc260"
	// todo username := os.Getenv("ATCODER_USERNAME")
	//   password := os.Getenv("ATCODER_PASSWORD")
	if err := GenAtCoderContestTemplates(contestID); err != nil {
		t.Fatal(err)
	}
}

// 仅限结束的比赛
// https://atcoder.jp/contests/abc161/tasks/abc161_f
// https://atcoder.jp/contests/abc161/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc161_f&orderBy=source_length
func TestGenAtCoderProblemTemplate(t *testing.T) {
	raw, err := ioutil.ReadFile("data.txt")
	if err != nil {
		t.Fatal(err)
	}
	problemURL := strings.TrimSpace(string(raw))
	if err := GenAtCoderProblemTemplate(problemURL); err != nil {
		t.Fatal(err)
	}
}
