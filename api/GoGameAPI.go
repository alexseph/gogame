package api

import (
	"net/http"
	"log"
	"io/ioutil"
	"github.com/alexseph/gogame/constants"
	"time"
	"strconv"
	"github.com/alexseph/gogame/model"
	"encoding/json"
)

//GetNextPS4GamesDates obtains the release date of a number of x PS4 games
func GetNextPS4GamesDates(quantity int) string {
	currentTime := time.Now().Unix()
	urlGetDates := "https://api-2445582011268.apicast.io/release_dates/?fields=*&filter[platform][eq]=48&order=date:asc&filter[date][gt]=" + strconv.Itoa(int(currentTime))

	jsonReleased := makeGetRequest(urlGetDates)
	if jsonReleased != "" {
		var releases []model.GameRelease
		parse_error_releases := json.Unmarshal([]byte(jsonReleased),&releases)
		if parse_error_releases != nil {
			log.Fatal(parse_error_releases)
		} else if len(releases) > 0 {
			retorno := make([]model.GameInfo, 0)
			for _ , release := range releases {
				if release.ID != 0 && release.DataString != "" {
					//Faz uma nova requisicao, so que pegando os dados do jogo
					urlGetGame := "https://api-2445582011268.apicast.io/games/"+strconv.Itoa(release.ID)+"?fields=*"
					jsonGame := makeGetRequest(urlGetGame)
					if jsonGame != "" {
						var game []model.GameData
						parse_error_game := json.Unmarshal([]byte(jsonGame),&game)
						if parse_error_game != nil {
							log.Fatal(parse_error_game)
						} else if len(game) > 0 && game[0].Name != "" {
							currentGame := model.GameInfo{
								Name: game[0].Name,
								ReleaseDate: release.DataString,
							}
							retorno = append(retorno, currentGame)
						}
					}
				}
			}

			if len(retorno) > 0 {
				jsonRetorno, parse_error_retorno := json.Marshal(retorno)
				if parse_error_retorno != nil {
					log.Fatal(parse_error_releases)
				} else {
					return string(jsonRetorno)
				}
			}
		}
	}
	return "NO VALUE"
}

//makeGetRequest make a get request
func makeGetRequest(url string) string {
	if url != "" {
		request, error_request := http.NewRequest("GET", url, nil)
		if error_request != nil {
			log.Fatal(error_request)
		} else {
			request.Header.Set(constants.IGBD_USER_KEY_PARAM, constants.IGBD_USER_KEY_VALUE)
			request.Header.Set(constants.IGBD_ACCEPT_PARAM, constants.IGBD_ACCEPT_VALUE)

			client := &http.Client{}
			response, error_response := client.Do(request)
			if error_response != nil {
				log.Fatal(error_response)
			} else {
				html, error_read := ioutil.ReadAll(response.Body)
				if error_read != nil {
					log.Fatal(error_read)
				} else {
					return string(html)
				}
			}
		}
	}
	return ""
}