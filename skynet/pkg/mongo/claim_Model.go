package mongo

import (
	root "skynet/pkg"

	"github.com/google/uuid"
)

type claimModel struct {
	claimDefinitionIdentifier string
	userIdentifier            string
}

type claimMethodModel struct {
	claimDefinitionIdentifier string
	attributesToType          map[string]string
}
