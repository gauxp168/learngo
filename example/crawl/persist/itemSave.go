package persist

import "log"

func ItemSave() chan interface{} {
	client,err := elastic.NewClient()
	if err != nil {
		panic(err)
	}
	out := make(chan interface{})
	go func() {
		itemcount := 0
		for  {
			item := <-out
			log.Printf("Item Saver :Got$%d %v", itemcount, item)
			save(client, item)
			itemcount++
		}
	}()
	return out
}
