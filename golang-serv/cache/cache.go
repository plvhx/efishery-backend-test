package cache

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type CacheData struct {
	Data map[string]interface{} `json:"data"`
}

type Cache struct {
	File string
}

func NewCache() *Cache {
	return &Cache{}
}

func (c *Cache) GetFile() string {
	return c.File
}

func (c *Cache) SetFile(file string) {
	c.File = file
}

func (c *Cache) Exists(key string) bool {
	file, err := os.Open(c.GetFile())

	if err != nil {
		return false
	}

	defer file.Close()

	buf, err := ioutil.ReadAll(file)

	if err != nil {
		return false
	}

	cacheData := &CacheData{}

	err = json.Unmarshal(buf, &cacheData)

	if err != nil {
		return false
	}

	if _, ok := cacheData.Data[key]; !ok {
		return false
	}

	return true
}

func (c *Cache) Put(key string, v interface{}) {
	if c.Exists(key) {
		return
	}

	file, err := os.OpenFile(c.GetFile(), os.O_RDWR, 0666)

	if err != nil {
		return
	}

	defer file.Close()

	buf, err := ioutil.ReadAll(file)

	if err != nil {
		return
	}

	cacheData := &CacheData{}

	err = json.Unmarshal(buf, &cacheData)

	if err != nil {
		return
	}

	cacheData.Data[key] = v

	buf, err = json.Marshal(&cacheData)

	if err != nil {
		return
	}

	_, err = file.Seek(0, io.SeekStart)

	if err != nil {
		return
	}

	_, err = file.Write(buf)

	if err != nil {
		return
	}
}

func (c *Cache) Update(key string, v interface{}) {
	if !c.Exists(key) {
		return
	}

	file, err := os.OpenFile(c.GetFile(), os.O_RDWR, 0666)

	if err != nil {
		return
	}

	defer file.Close()

	buf, err := ioutil.ReadAll(file)

	if err != nil {
		return
	}

	cacheData := &CacheData{}

	err = json.Unmarshal(buf, &cacheData)

	if err != nil {
		return
	}

	cacheData.Data[key] = v

	buf, err = json.Marshal(cacheData)

	if err != nil {
		return
	}

	_, err = file.Seek(0, io.SeekStart)

	if err != nil {
		return
	}

	_, err = file.Write(buf)

	if err != nil {
		return
	}
}

func (c *Cache) Get(key string) interface{} {
	file, err := os.Open(c.GetFile())

	if err != nil {
		return nil
	}

	defer file.Close()

	buf, err := ioutil.ReadAll(file)

	if err != nil {
		return nil
	}

	cacheData := &CacheData{}

	err = json.Unmarshal(buf, &cacheData)

	if err != nil {
		return nil
	}

	if val, ok := cacheData.Data[key]; ok {
		return val
	}

	return nil
}

func (c *Cache) All() *CacheData {
	file, err := os.Open(c.GetFile())

	if err != nil {
		return nil
	}

	defer file.Close()

	buf, err := ioutil.ReadAll(file)

	if err != nil {
		return nil
	}

	cacheData := &CacheData{}

	err = json.Unmarshal(buf, &cacheData)

	if err != nil {
		return nil
	}

	return cacheData
}

func AuthCacheFactory() (*Cache, error) {
	dir, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	dir = fmt.Sprintf("%s/../cache/auth.json", dir)

	cache := NewCache()
	cache.SetFile(dir)

	return cache, nil
}

func CurrencyCacheFactory() (*Cache, error) {
	dir, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	dir = fmt.Sprintf("%s/../cache/currency.json", dir)

	cache := NewCache()
	cache.SetFile(dir)

	return cache, nil
}

func DataCacheFactory() (*Cache, error) {
	dir, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	dir = fmt.Sprintf("%s/../cache/data.json", dir)

	cache := NewCache()
	cache.SetFile(dir)

	return cache, nil
}
