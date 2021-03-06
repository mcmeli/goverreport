package report

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var results []Summary
var cover1 Summary
var cover2 Summary

func init() {
	cover1 = Summary{
		Name:          "file1",
		BlockCoverage: 0.5, StmtCoverage: 0.6,
		MissingBlocks: 10, MissingStmts: 8}

	cover2 = Summary{
		Name:          "file2",
		BlockCoverage: 0.6, StmtCoverage: 0.5,
		MissingBlocks: 8, MissingStmts: 10}

	results = []Summary{cover1, cover2}
}

func TestSortByFileName(t *testing.T) {
	assert.NoError(t, sortResults(results, "filename", "asc"))
	assert.Equal(t, results, []Summary{cover1, cover2})
}

func TestSortByBlockCoverage(t *testing.T) {
	assert.NoError(t, sortResults(results, "block", "desc"))
	assert.Equal(t, results, []Summary{cover2, cover1})

}

func TestSortByStmtCoverage(t *testing.T) {
	assert.NoError(t, sortResults(results, "stmt", "desc"))
	assert.Equal(t, results, []Summary{cover1, cover2})
}

func TestSortByMissingBlocks(t *testing.T) {
	assert.NoError(t, sortResults(results, "missing-blocks", "asc"))
	assert.Equal(t, results, []Summary{cover2, cover1})
}

func TestSortByMissingStmts(t *testing.T) {
	assert.NoError(t, sortResults(results, "missing-stmts", "asc"))
	assert.Equal(t, results, []Summary{cover1, cover2})
}

func TestInvalidParameters(t *testing.T) {
	assert.Error(t, sortResults(results, "xxx", "asc"))
	assert.Error(t, sortResults(results, "block", "yyy"))
}

func TestReport(t *testing.T) {
	assert := assert.New(t)
	report, err := GenerateReport("../sample_coverage.out", "", []string{}, "block", "desc")
	assert.NoError(err)
	assert.InDelta(75.6, report.Total.BlockCoverage, 0.1)
	assert.InDelta(80, report.Total.StmtCoverage, 0.1)
	assert.Equal(1801, report.Total.Stmts)
	assert.Equal(1097, report.Total.Blocks)
}

func TestInvalidCoverProfile(t *testing.T) {
	_, err := GenerateReport("../xxx.out", "", []string{}, "block", "desc")
	assert.Error(t, err)
}
