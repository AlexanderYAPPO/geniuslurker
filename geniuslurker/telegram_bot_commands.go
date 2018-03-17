package geniuslurker

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/bot-api/telegram"
	"github.com/bot-api/telegram/telebot"
	"golang.org/x/net/context"
)

// SearchCommand requests geniuslurker for search results from Genius
func SearchCommand(ctx context.Context, arg string) error {

	api := telebot.GetAPI(ctx)
	update := telebot.GetUpdate(ctx)
	chatID := update.Chat().ID

	searchResults := GetGeniusLurkerFetcherClient().Search(arg)

	redisClient := GetRedisClient()
	redisKey := "search:" + strconv.FormatInt(chatID, 10)
	exists, err := redisClient.Exists(redisKey).Result()
	if err != nil {
		fmt.Println("whoops:", err)
		panic(err)
	}
	if exists != 0 {
		//cleanup previous values
		_, err := redisClient.Del(redisKey).Result()
		if err != nil {
			fmt.Println("whoops:", err)
			panic(err)
		}
	}
	for _, searchResult := range searchResults {
		bSearchResult, _ := json.Marshal(searchResult)
		_, err = redisClient.RPush(redisKey, bSearchResult).Result()
		if err != nil {
			fmt.Println("whoops:", err)
			panic(err)
		}
	}

	message := "Results: \n"
	for index, searchResult := range searchResults {
		message += strconv.Itoa(index) + ": " + searchResult.FullTitle + "\n"
	}
	_, err = api.SendMessage(ctx,
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
	size, err := redisClient.LLen(redisKey).Result()
	if err != nil {
		fmt.Println("whoops:", err)
		panic(err)
	}

	index, err := strconv.ParseInt(arg, 10, 64)
	if err != nil || index < 0 || index > size {
		fmt.Println("Incorrect input in chat: "+strconv.FormatInt(chatID, 10), err)
		_, err = api.SendMessage(ctx,
			telegram.NewMessagef(update.Chat().ID,
				"Incorrect input. Lyrics are not yet search or index is not in the boundaries.",
			))
		return err
	}

	searchResultB, err := redisClient.LIndex(redisKey, index).Bytes()
	if err != nil {
		fmt.Println("whoops:", err)
		panic(err)
	}
	var searchResult SearchResult
	json.Unmarshal(searchResultB, &searchResult)

	lyrics := GetGeniusLurkerFetcherClient().GetLyrics(searchResult)

	_, err = api.SendMessage(ctx,
		telegram.NewMessagef(update.Chat().ID,
			lyrics,
		))
	return err
}
