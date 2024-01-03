package storage

import "errors"

var (
  ErrUrlExists = errors.New("url exists")
  ErrUrlNotFound = errors.New("url not found")
)
