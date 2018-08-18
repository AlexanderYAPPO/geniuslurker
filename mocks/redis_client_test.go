package mocks

import (
	"testing"

	"github.com/AlexanderYAPPO/geniuslurker/datastructers"
	"github.com/stretchr/testify/suite"
)

type RedisClientMockTestSuite struct {
	suite.Suite
	*RedisClientMock
}

func (suite *RedisClientMockTestSuite) SetupTest() {
	suite.RedisClientMock = NewRedisClient()
}

func TestRedisClientMockTestSuite(t *testing.T) {
	suite.Run(t, new(RedisClientMockTestSuite))
}

func (suite *RedisClientMockTestSuite) TestExistsPassingMissingKey() {
	suite.False(suite.RedisClientMock.Exists("test"))
}

func (suite *RedisClientMockTestSuite) TestSingleItemAdded() {
	st := datastructers.SearchResult{
		FullTitle: "test_1",
		URL:       "http://example.com",
	}
	suite.RedisClientMock.SearchResultsRPushJSON("test", st)
	suite.True(suite.RedisClientMock.Exists("test"))
	suite.Equal(int64(1), suite.RedisClientMock.LLen("test"))
}

func (suite *RedisClientMockTestSuite) TestTwoItemsSingleKeyAdded() {
	st := datastructers.SearchResult{
		FullTitle: "test_1",
		URL:       "http://example.com",
	}
	suite.RedisClientMock.SearchResultsRPushJSON("test", st)
	st = datastructers.SearchResult{
		FullTitle: "test_2",
		URL:       "http://example.com",
	}
	suite.RedisClientMock.SearchResultsRPushJSON("test", st)
	suite.Equal(int64(2), suite.RedisClientMock.LLen("test"))
}

func (suite *RedisClientMockTestSuite) TestFetchingByIndex() {
	st := datastructers.SearchResult{
		FullTitle: "test_1",
		URL:       "http://example.com",
	}
	suite.RedisClientMock.SearchResultsRPushJSON("test", st)
	st = datastructers.SearchResult{
		FullTitle: "test_2",
		URL:       "http://example.com",
	}
	suite.RedisClientMock.SearchResultsRPushJSON("test", st)
	suite.Equal("test_1", suite.RedisClientMock.SearchResultsIndexJSON("test", 0).FullTitle)
	suite.Equal("test_2", suite.RedisClientMock.SearchResultsIndexJSON("test", 1).FullTitle)
}

func (suite *RedisClientMockTestSuite) TestDeletingElement() {
	st := datastructers.SearchResult{
		FullTitle: "test_1",
		URL:       "http://example.com",
	}
	suite.RedisClientMock.SearchResultsRPushJSON("test", st)
	suite.RedisClientMock.Del("test")
	suite.False(suite.RedisClientMock.Exists("test"))
}

func (suite *RedisClientMockTestSuite) TestIndexOutOfRange() {
	st := datastructers.SearchResult{
		FullTitle: "test_1",
		URL:       "http://example.com",
	}
	suite.RedisClientMock.SearchResultsRPushJSON("test", st)
	suite.Panics(func() { suite.RedisClientMock.SearchResultsIndexJSON("test", int64(2)) })
}
