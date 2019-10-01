// Copyright 2012-present Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package elastic

import (
	"encoding/json"
	"testing"
)

func TestScriptScoreQuery(t *testing.T) {
	q := NewScriptScoreQuery(NewTermQuery("user", "ki"), NewScript("doc['num1'.value > 1"))
	src, err := q.Source()
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(src)
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"script_score":{"query":{"term":{"user":"ki"}}},"script":{"source":"doc['num1'.value \u003e 1"}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func TestScriptScoreQueryWithMinScore(t *testing.T) {
	q := NewScriptScoreQuery(NewTermQuery("user", "ki"), NewScript("doc['num1'.value > 1")).MinScore(0.01)
	src, err := q.Source()
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(src)
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"script_score":{"query":{"term":{"user":"ki"}}},"script":{"source":"doc['num1'.value \u003e 1"},"min_score":0.01}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
