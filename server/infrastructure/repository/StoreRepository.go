package repository

import (
	"context"
	"database/sql"
	"server/core/entity"
	"server/core/infra/repository"
	"server/db/models"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var _ repository.IStoreRepository = &StoreRepository{}

type StoreRepository struct {
	db *sql.DB
}

func NewStoreRepository() *StoreRepository {
	db := InitDB()

	return &StoreRepository{
		db: db,
	}
}

func (pr *StoreRepository) Save(updateStore *entity.Store) error {
	store := models.Store{
		ID:              updateStore.ID.String(),
		Name:            updateStore.Name,
		Address:         updateStore.Address,
		Tel:             updateStore.Tel,
		Parking:         updateStore.Parking,
		AccessInfo:      updateStore.AccessInfo,
		IsActive:        updateStore.IsActive,
		QRCode:          updateStore.QRCode.String(),
		UnLimitedQRCode: updateStore.UnLimitedQRCode.String(),
	}
	err := store.Upsert(context.Background(), pr.db, true, []string{"id"}, boil.Infer(), boil.Infer())
	return err
}

func (pr *StoreRepository) Delete(storeID uuid.UUID) error {
	deleteStore, err := models.FindStore(context.Background(), pr.db, storeID.String())
	if err != nil {
		return err
	}
	_, err = deleteStore.Delete(context.Background(), pr.db)
	return err
}
