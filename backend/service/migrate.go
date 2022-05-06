package service

import "github.com/go-emix/fortune/backend/service/system"

func Migrate() error {
	return system.Migrate()
}
