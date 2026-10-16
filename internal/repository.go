package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func writeToFile(data Data) error {
	write,err:=json.Marshal(data)
	if err!=nil {
		return err
	}
	fmt.Println("writing")
	err=os.WriteFile("/Users/shgd/sharathdoes/go/goshort/internal/db.json", write , 0664)
	if err!=nil {
		return err
	}
	return nil
}

func getFromFile() (Data, error) {
	data,err:=os.ReadFile("/Users/shgd/sharathdoes/go/goshort/internal/db.json")
	if err!=nil {
		return Data{},err
	}
		var read Data
	if err := json.Unmarshal(data, &read); err != nil {
		return Data{}, err
	}

	return read, nil

}

func createShortcut( key string, url string) error {
	shortcut:=URL{key,url}
	data,err:=getFromFile()
	if err!=nil {
		return err
	}
	data.URLS=append(data.URLS, shortcut)
	fmt.Println("writing this", data)
	return writeToFile(data)
}

func getShortcut(key string ) (string, error) {
	data,err:=getFromFile()
	if err!=nil {
		return "", err
	}
	for _, url := range data.URLS {
		if url.Key == key {
			return url.Redirect, nil
		}
	}
	return "", nil
}

func updateShortcut(key string , new_redirect string) ( error) {
	data,err:=getFromFile()
	if err!=nil {
		return  err
	}
	for i := range data.URLS {
		if data.URLS[i].Key == key {
			data.URLS[i].Redirect = new_redirect
			return writeToFile(data)
		}
	}
	return nil
}

func deleteShortcut(key string ) ( error) {
	data,err:=getFromFile()
	if err!=nil {
		return  err
	}
	found := false
	var newData Data
		for i := range data.URLS {
		if data.URLS[i].Key != key {
			newData.URLS = append(newData.URLS, data.URLS[i])
		} else {
			found = true
		}
	}

	if !found {
		return errors.New("key doesn't exist")
	}

	return writeToFile(newData)
}