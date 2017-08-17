package api

import (
  "fmt"
  "strings"
  "sync"
  "net/http"
  "io/ioutil"
  "github.com/jacki3lene/goean/config"
)

func List(hotelIds []int, checkIn string, checkOut string, response chan string, wg *sync.WaitGroup) {
  defer wg.Done()

  var stringIds = make([]string, len(hotelIds))
  for i, id := range(hotelIds) {
    stringIds[i] = fmt.Sprintf("%d", id)
  }

  //fmt.Println("Processing:", stringIds)

  var url = fmt.Sprintf("%v/list?apiKey=%v&arrivalDate=%v&cid=55505&departureDate=%v&hotelIdList=%v&maxRatePlanCount=50&numberOfResults=600&options=HOTEL_SUMMARY,ROOM_RATE_DETAILS&sort=OVERALL_VALUE&supplierType=E&useCache=true&minorRev=%v", config.Conf.Url, config.Conf.ApiKey, checkIn, checkOut, strings.Join(stringIds, ","), config.Conf.MinorRev)
  fmt.Println("Url:", url)

  resp, _ := http.Get(url)

  defer resp.Body.Close()
  body, _ := ioutil.ReadAll(resp.Body)

  response <- string(body)
}