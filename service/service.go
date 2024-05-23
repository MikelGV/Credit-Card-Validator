package main

import "context"

type Service interface {
    getLunhCheck(context.Context) (bool, error)
}

func getLunhCheck() (bool) {
    return false
}
