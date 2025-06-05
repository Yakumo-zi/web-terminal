package apiserver

import (
	"context"
	_ "github.com/Yakumo-zi/web-terminal/ent/runtime"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/domain"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/repo"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/service"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssetDelete(t *testing.T) {
	svc := service.NewService()
	assetRepo := svc.Repo.Assets()
	defer svc.Repo.Close()
	assets, count, err := assetRepo.List(context.Background(), &repo.ListOptions{Offset: 0, Limit: 20})
	if err != nil {
		t.Errorf("assetRepo.List() error = %v", err)
	}
	assert.Equal(t, count, 2)
	for idx, asset := range assets {
		t.Logf("asset[%d]: %+v", idx, asset)
	}
	err = assetRepo.Delete(context.Background(), assets[0].Id)
	if err != nil {
		t.Errorf("assetRepo.Delete() error = %v", err)
	}
	assets, count, err = assetRepo.List(context.Background(), &repo.ListOptions{Offset: 0, Limit: 20})
	if err != nil {
		t.Errorf("assetRepo.List() error = %v", err)
	}
	assert.Equal(t, count, 1)
}

func TestAssetCreate(t *testing.T) {
	svc := service.NewService()
	assetRepo := svc.Repo.Assets()
	id := uuid.New()
	err := assetRepo.Create(context.Background(), &domain.Asset{
		Name: "test",
		Id:   id,
		Ip:   "127.0.0.1",
		Port: 1234,
		Type: "host",
	})
	if err != nil {
		t.Errorf("assetRepo.Create() error = %v", err)
	}
	asset, err := assetRepo.Get(context.Background(), id)
	if err != nil {
		t.Errorf("assetRepo.Get() error = %v", err)
	}
	assert.Equal(t, asset.Name, "test")
	assert.Equal(t, asset.Ip, "127.0.0.1")
	assert.Equal(t, asset.Port, 1234)
	assert.Equal(t, asset.Type, "host")

	asset.Name = "modify"
	err = assetRepo.Update(context.Background(), asset)
	if err != nil {
		t.Errorf("assetRepo.Update() error = %v", err)
	}
	asset, err = assetRepo.Get(context.Background(), id)
	if err != nil {
		t.Errorf("assetRepo.Get() error = %v", err)
	}
	assert.Equal(t, asset.Name, "modify")
	assert.Equal(t, asset.Ip, "127.0.0.1")
	assert.Equal(t, asset.Port, 1234)
	assert.Equal(t, asset.Type, "host")

	err = assetRepo.Create(context.Background(), &domain.Asset{
		Name: "test",
		Id:   uuid.New(),
		Ip:   "127.0.0.2",
		Port: 1234,
		Type: "host",
	})
	if err != nil {
		t.Errorf("assetRepo.Create() error = %v", err)
	}
	_, count, err := assetRepo.List(context.Background(), &repo.ListOptions{Offset: 0, Limit: 20})
	if err != nil {
		t.Errorf("assetRepo.List() error = %v", err)
	}
	assert.Equal(t, count, 2)
	svc.Repo.Close()
}
