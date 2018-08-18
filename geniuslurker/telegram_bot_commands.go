package geniuslurker

import (
	"strconv"
	"strings"

	"github.com/bot-api/telegram"
	"github.com/bot-api/telegram/telebot"
	"golang.org/x/net/context"
)

const maxTelegramMessageLength = 4096

// SearchCommand requests geniuslurker for search results from Genius
func SearchCommand(ctx context.Context, arg string) error {
	api := telebot.GetAPI(ctx)
	update := telebot.GetUpdate(ctx)
	chatID := update.Chat().ID

	searchResults := GetFetcherClient().Search(arg)

	redisClient := GetRedisClient()
	redisKey := "search:" + strconv.FormatInt(chatID, 10)
	exists := redisClient.Exists(redisKey)
	if exists {
		//cleanup previous values
		redisClient.Del(redisKey)
	}
	for _, searchResult := range searchResults {
		redisClient.SearchResultsRPushJSON(redisKey, searchResult)
	}

	message := "Results: \n"
	for index, searchResult := range searchResults {
		message += strconv.Itoa(index) + ": " + searchResult.FullTitle + "\n"
	}
	_, err := api.SendMessage(ctx,
		telegram.NewMessagef(update.Chat().ID,
			message,
		))
	return err
}

// GetLyricsCommand gets lyrics from genius lurker
func GetLyricsCommand(ctx context.Context, arg string) error {
	api := telebot.GetAPI(ctx)
	update := telebot.GetUpdate(ctx)
	chatID := update.Chat().ID

	redisClient := GetRedisClient()
	redisKey := "search:" + strconv.FormatInt(chatID, 10)
	size := redisClient.LLen(redisKey)

	index, err := strconv.ParseInt(arg, 10, 64)
	if err != nil || index < 0 || index > size {
		InfoLogger.Println("Incorrect input in chat: "+strconv.FormatInt(chatID, 10), err)
		_, err = api.SendMessage(ctx,
			telegram.NewMessagef(update.Chat().ID,
				"Incorrect input. Lyrics are not yet search or index is not in the boundaries.",
			))
		return err
	}

	searchResult := redisClient.SearchResultsIndexJSON(redisKey, index)

	lyrics := GetFetcherClient().GetLyrics(searchResult)

	lyricsBlocks := splitTextOnBlocks(lyrics)
	for _, block := range lyricsBlocks {
		_, err = api.SendMessage(ctx,
			telegram.NewMessagef(update.Chat().ID,
				block,
			))
		if err != nil {
			ErrorLogger.Panicln("Failed to send message", err)
		}
	}
	return err
}

func splitTextOnBlocks(originalText string) []string {
	if len(originalText) <= maxTelegramMessageLength {
		return []string{originalText}
	}
	var resultBlocks []string
	verses := strings.Split(originalText, "\n\n")
	left := 0
	right := 0
	currentBlockLength := 0
	blockLengthAfterAppending := 0
	for ; right < len(verses); right++ {
		currentVerse := verses[right]
		if len(currentVerse) > maxTelegramMessageLength {
			ErrorLogger.Panicln("The length of a block exceeds the maxmium acceptable length")
		}
		if currentBlockLength == 0 {
			blockLengthAfterAppending = len(currentVerse)
		} else {
			blockLengthAfterAppending = currentBlockLength + len(currentVerse) + 2
		}
		if blockLengthAfterAppending > maxTelegramMessageLength {
			resultBlocks = append(resultBlocks, strings.Join(verses[left:right], "\n\n"))
			left = right
			currentBlockLength = len(currentVerse)
		} else {
			currentBlockLength += len(currentVerse)
		}
	}
	resultBlocks = append(resultBlocks, strings.Join(verses[left:], "\n\n"))
	return resultBlocks
}
