package cilent

import (
	"crawler/distributed/config"
	"crawler/distributed/rpcsupport"
	"crawler/engine"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item saver: #%d: %v", itemCount, item.Payload)
			itemCount++

			result := ""
			err := client.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				log.Printf("Item saver: error saving item %v: %v:", item, err)
			}
		}
	}()
	return out, nil
}
