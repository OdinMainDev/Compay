package esfile

import (
	"compay/core/config"
	"compay/core/eslog"
	"context"

	"cloud.google.com/go/firestore"
	cloud "cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type ESFile struct {
	client  *firestore.Client
	storage *cloud.Client
	ctx     context.Context
}

type imageStructure struct {
	ImageName string `json:"imageName"`
	URL       string `json:"url"`
}

func CreateInstance() *ESFile {
	if config.CdnUploaderDisabled {
		return nil
	}
	sa := option.WithCredentialsFile("files/serviceAccountKey.json")
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		eslog.PanicLog("INIT FIRESTORE FAIL ", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		eslog.PanicLog("INIT FIRESTORE FAIL ", err)
	}

	storage, err := cloud.NewClient(ctx, sa)
	if err != nil {
		eslog.PanicLog("INIT FIRESTORE FAIL ", err)
	}

	eslog.Log("Firestore initialized")
	return &ESFile{
		client:  client,
		storage: storage,
		ctx:     ctx,
	}
}

func (that *ESFile) GetStorage() *cloud.Client {
	return that.storage
}
