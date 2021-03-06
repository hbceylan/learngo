// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"io"
	"io/ioutil"
	"os"
)

func fromFile(f *os.File) (list, error) {
	return fromReader(f)
}

func fromReader(r io.Reader) (list, error) {	
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var store list

	if err := json.Unmarshal(data, &store); err != nil {
		return nil, err
	}

	return store, nil
}
