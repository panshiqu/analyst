package cache

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/panshiqu/analyst/define"
	"github.com/panshiqu/analyst/utils"
)

var alertsMutex sync.Mutex
var alerts map[string]*define.Alert

func init() {
	alerts = make(map[string]*define.Alert)
}

func PrintAlerts() {
	alertsMutex.Lock()
	defer alertsMutex.Unlock()

	for k, v := range alerts {
		fmt.Printf("%s %#v\n", k, v)
	}
}

func HandleAlert(id int64, name string, address string, price float64, notify string) error {
	s := fmt.Sprintf("%d%s%s%g%s", id, name, address, price, notify)

	alertsMutex.Lock()
	defer alertsMutex.Unlock()

	if _, ok := alerts[s]; ok {
		delete(alerts, s)
		return nil
	}

	alerts[s] = &define.Alert{
		ID:      id,
		Name:    name,
		Address: address,
		Price:   price,
		Notify:  notify,
	}
	return nil
}

func NotifyAlerts(address string, price float64) error {
	alertsMutex.Lock()
	defer alertsMutex.Unlock()

	for k, v := range alerts {
		if v.Address != address {
			continue
		}

		var text string
		switch {
		case v.Price >= 0:
			if price < v.Price {
				continue
			}
			text = fmt.Sprintf("👆%s price: %g", utils.Address2Symbol(address), price)
		case v.Price < 0:
			if price > -v.Price {
				continue
			}
			text = fmt.Sprintf("👇%s price: %g", utils.Address2Symbol(address), price)
		}

		msg := tgbotapi.NewMessage(v.ID, text)
		msg.DisableNotification = utils.IsDisableNotification(v.Notify)

		if _, err := Bot.Send(msg); err != nil {
			return utils.Wrap(err)
		}

		delete(alerts, k)
	}

	return nil
}

func NotifyCurrentAlerts(name string, address string) error {
	alertsMutex.Lock()
	defer alertsMutex.Unlock()

	for _, v := range alerts {
		if v.Name != name {
			continue
		}

		if address != "" && v.Address != address {
			continue
		}

		msg := tgbotapi.NewMessage(v.ID, fmt.Sprintf("set %s %g %s", utils.Address2Symbol(v.Address), v.Price, v.Notify))

		if _, err := Bot.Send(msg); err != nil {
			return utils.Wrap(err)
		}
	}

	return nil
}

const alertsFile = "/tmp/alerts.json"

func SaveAlerts() error {
	alertsMutex.Lock()
	defer alertsMutex.Unlock()

	data, err := json.Marshal(alerts)
	if err != nil {
		return utils.Wrap(err)
	}

	if err := ioutil.WriteFile(alertsFile, data, 0644); err != nil {
		return utils.Wrap(err)
	}

	return nil
}

func LoadAlerts() error {
	if _, err := os.Stat(alertsFile); os.IsNotExist(err) {
		return nil
	}

	alertsMutex.Lock()
	defer alertsMutex.Unlock()

	return utils.ReadJSON(alertsFile, &alerts)
}
