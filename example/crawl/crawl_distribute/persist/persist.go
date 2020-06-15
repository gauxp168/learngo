package persist

import "simqo.com/mygospace/learngo/example/crawl/engine"

type ItemService struct {
	Client *elastic.Client
}

func (s *ItemService) Save(item engine.Item, result *string) error  {
	persist.Save()
}

