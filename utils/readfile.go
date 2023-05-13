package utils

import "io/ioutil"

func ReadFile(filename string) (string, error) {
    content, err := ioutil.ReadFile(filename)
    if err != nil {
        return "", err
    }
    return string(content), nil
}