package model

import (
	"time"
	"google.golang.org/appengine/datastore"
	"context"
	"strings"
	"google.golang.org/appengine/log"
	"fmt"
)


const (
	HOST_ENTITY_NAME = "HostEntity"
	HOST_NAME = "HostName"
	READY = "Ready"
	REGISTERED_AT = "Registered_At"
)

type HostEntity struct {
	HostName string `json:"HostName"`
	Ready bool `json:"Ready"`
	RegisteredAt time.Time `json:"Registered_At"`
}

func DoesHostExist(ctx context.Context, hostName string) bool {
	key := datastore.NewKey(ctx, HOST_ENTITY_NAME, hostName, 0, nil)
	var host HostEntity
	err := datastore.Get(ctx, key, &host)
	log.Infof(ctx, "Check Host " + hostName)

	if err != nil {
		log.Infof(ctx, err.Error())
		if strings.HasPrefix(err.Error(), `datastore: no such entity`) {
			return false
		}
	}

	return true
}


func CreateHost(ctx context.Context, hostName string) (*HostEntity, error) {
	newHost := &HostEntity{
		hostName,
		false,
		time.Now(),
	}

	key := datastore.NewKey(ctx, HOST_ENTITY_NAME, hostName, 0, nil)
	insKey, err := datastore.Put(ctx, key, newHost)

	if err != nil {
		log.Errorf(ctx, "ERROR INSERTING Host: %v", err.Error())
		return nil, err
	} else {
		createdHost, err := getHostByKey(ctx, insKey.StringID())
		if err != nil {
			log.Errorf(ctx, "ERROR GETTING Host OUTPUT: %v", err.Error())
			return nil, err
		}

		return createdHost, nil
	}
}

func getHostByKey(ctx context.Context, key string) (*HostEntity, error) {
	hostKey := datastore.NewKey(ctx, HOST_ENTITY_NAME, key, 0, nil)
	var host HostEntity
	err := datastore.Get(ctx, hostKey, &host)

	if err != nil {
		if strings.HasPrefix(err.Error(), `datastore: no such entity`) {
			err = fmt.Errorf(`user '%v' not found`, host)
		}
		return nil, err
	}
	return &host, nil
}

func UpdateHost(ctx context.Context, host *HostEntity) bool {
	if DoesHostExist(ctx, host.HostName) {
		key := datastore.NewKey(ctx, HOST_ENTITY_NAME, host.HostName, 0, nil)
		_, err := datastore.Put(ctx, key, host)
		if err != nil {
			log.Errorf(ctx, "ERROR UPDATING USER: %v", err.Error())
			return false
		}
		return true
	}
	return false
}

func GetAllHosts(ctx context.Context, isReady bool) (*[]string, error) {
	query := datastore.NewQuery(HOST_ENTITY_NAME).KeysOnly()

	if isReady {
		query.Filter(READY + "=", "true")
	}

	var hosts[]string
	_, err := query.GetAll(ctx, &hosts)

	if err != nil {
		log.Errorf(ctx, "Error in getting hosts from db - " + err.Error())
		return nil, err
	}

	return &hosts, nil

}