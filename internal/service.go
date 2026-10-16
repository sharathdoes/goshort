package internal

import "errors"

// import "context"

func createShortCut(key string, redirect string) error {
	exists, err := getShortcut(key)
	if err != nil {
		return err
	}
	if exists != "" {
		updateShortCut(key, redirect)
	} else {
		err = createShortcut(key, redirect)
		if err != nil {
			return err
		}
	}
	return nil
}

func updateShortCut(key string, redirect string) error {
	if redirect == "" {
		return errors.New("key doesnt exist, cant update")
	}
	err:= updateShortcut(key, redirect)
	return err
}

func deleteShortCut(  key string ) error {
	err:= deleteShortcut(key)
	return err
}

func getRedirectBykey(key string) (string, error) {
	redirect, err := getShortcut(key)
	return redirect, err
}
