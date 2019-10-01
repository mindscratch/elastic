// Copyright 2012-present Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package elastic

import "errors"

// ScriptScoreQuery uses a script to provide a custom score for returned documents.
//
// For details, see
// https://www.elastic.co/guide/en/elasticsearch/reference/7.x/query-dsl-script-score-query.html
type ScriptScoreQuery struct {
	query    Query
	script   *Script
	minScore *float64
}

// NewScriptScoreQuery creates and initializes a new ScriptScoreQuery.
func NewScriptScoreQuery(query Query, script *Script) *ScriptScoreQuery {
	return &ScriptScoreQuery{
		query:  query,
		script: script,
	}
}

// MinScore sets the minimum score
func (q *ScriptScoreQuery) MinScore(minScore float64) *ScriptScoreQuery {
	q.minScore = &minScore
	return q
}

// Source returns JSON for the query.
func (q *ScriptScoreQuery) Source() (interface{}, error) {
	if q.query == nil {
		return nil, errors.New("ScriptScoreQuery expected a query")
	}
	if q.script == nil {
		return nil, errors.New("ScriptScoreQuery expected a script")
	}
	source := make(map[string]interface{})
	query := make(map[string]interface{})
	source["script_score"] = query

	src, err := q.query.Source()
	if err != nil {
		return nil, err
	}
	query["query"] = src

	src, err = q.script.Source()
	if err != nil {
		return nil, err
	}
	query["script"] = src

	if q.minScore != nil {
		query["min_score"] = *q.minScore
	}
	return source, nil
}
