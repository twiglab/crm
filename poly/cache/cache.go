package cache

import (
	"context"
	"errors"

	"github.com/allegro/bigcache/v3"
	"github.com/twiglab/crm/poly/orm/ent"
	"github.com/twiglab/crm/poly/orm/ent/poly"
)

type PolyCache struct {
	local  bigcache.BigCache
	client *ent.Client
}

func (c *PolyCache) Get(key string) (*PolyItem, error) {
	var (
		err error
		bs  []byte
		pi  *PolyItem
	)

	if bs, err = c.local.Get(key); err == nil {
		pi = new(PolyItem)
		if _, err = pi.UnmarshalMsg(bs); err != nil {
			return nil, err
		}
		return pi, nil
	}

	if errors.Is(err, bigcache.ErrEntryNotFound) {
		return c.load(key)
	}

	return nil, err
}

func (c *PolyCache) Set(key string, pi *PolyItem) error {
	bs, err := pi.MarshalMsg(nil)
	if err != nil {
		return err
	}
	return c.local.Set(key, bs)
}

func (c *PolyCache) load(key string) (*PolyItem, error) {
	q := c.client.Poly.Query()
	q.Where(poly.CodeEQ(key))
	p, err := q.Only(context.Background())
	if err != nil {
		return nil, err
	}

	pi := &PolyItem{Code: p.Code}
	err = c.Set(pi.Code, pi)
	return pi, err
}
