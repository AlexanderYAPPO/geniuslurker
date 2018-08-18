package geniuslurker

import (
	"io/ioutil"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type TelegramBotCommansTestSuite struct {
	suite.Suite
	Environment string
}

func (suite *TelegramBotCommansTestSuite) SetupTest() {
	suite.Environment = viper.GetString("environment")
	viper.Set("environment", "tests")
}

func (suite *TelegramBotCommansTestSuite) TearDownTest() {
	viper.Set("environment", suite.Environment)
}

func (suite *TelegramBotCommansTestSuite) TestSplitTextOnBlocks() {
	lyrics, _ := ioutil.ReadFile("../fixtures/long_lyrics_example")
	blocks := splitTextOnBlocks(string(lyrics))
	suite.Equal(2, len(blocks))
}

func TestSplitTextOnBlocksTestSuite(t *testing.T) {
	suite.Run(t, new(TelegramBotCommansTestSuite))
}
