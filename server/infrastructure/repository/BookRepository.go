package repository

import (
	"context"
	"database/sql"

	"server/core/entity"
	repository "server/core/infra/repository"
	"server/db/models"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var _ repository.IBookRepository = &BookRepository{}

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository() *BookRepository {
	db := InitDB()

	return &BookRepository{
		db: db,
	}
}

func (pq *BookRepository) Save(entity *entity.Booking) error {
	guest := models.BookGuestDatum{
		ID:            uuid.New().String(),
		FirstName:     entity.GuestData.FirstName,
		LastName:      entity.GuestData.LastName,
		FirstNameKana: entity.GuestData.FirstNameKana,
		LastNameKana:  entity.GuestData.LastNameKana,
		CompanyName:   null.StringFromPtr(entity.GuestData.CompanyName),
		BirthDate:     entity.GuestData.BirthDate,
		ZipCode:       null.StringFromPtr(entity.GuestData.ZipCode),
		Prefecture:    entity.GuestData.Prefecture.ToInt(),
		City:          null.StringFromPtr(entity.GuestData.City),
		Address:       null.StringFromPtr(entity.GuestData.Address),
		Tel:           null.StringFromPtr(entity.GuestData.Tel),
		Mail:          entity.GuestData.Mail,
	}

	plan := models.BookPlan{
		ID:              uuid.New().String(),
		PlanID:          entity.BookPlan.ID,
		Title:           entity.BookPlan.Title,
		Price:           int(entity.BookPlan.Price),
		ImageURL:        entity.BookPlan.ImageURL,
		RoomType:        int(entity.BookPlan.RoomType),
		MealTypeMorning: entity.BookPlan.MealType.Morning,
		MealTypeDinner:  entity.BookPlan.MealType.Dinner,
		SmokeType:       int(entity.BookPlan.SmokeType),
		Overview:        entity.BookPlan.OverView,
		StoreID:         entity.BookPlan.StoreID.String(),
	}

	book := models.UserBook{
		ID:              entity.ID.String(),
		TLBookingNumber: entity.TlBookingNumber,
		StayFrom:        entity.StayFrom,
		StayTo:          entity.StayTo,
		Adult:           int(entity.Adult),
		Child:           int(entity.Child),
		RoomCount:       int(entity.RoomCount),
		CheckInTime:     entity.CheckInTime.String(),
		TotalCost:       int(entity.TotalCost),
		GuestDataID:     guest.ID,
		BookPlanID:      plan.ID,
		BookUserID:      entity.BookUserID.String(),
		Note:            null.StringFrom(entity.Note),
	}
	ctx := context.Background()
	tran, err := pq.db.BeginTx(ctx, nil)
	defer func() {
		if err != nil {
			_ = tran.Rollback()
		} else {
			_ = tran.Commit()
		}
	}()

	if err != nil {
		return err
	}

	err = plan.Upsert(ctx, pq.db, true, []string{"id"}, boil.Infer(), boil.Infer())
	if err != nil {
		return err
	}

	err = guest.Upsert(ctx, pq.db, true, []string{"id"}, boil.Infer(), boil.Infer())
	if err != nil {
		return err
	}

	err = book.Upsert(ctx, pq.db, true, []string{"id"}, boil.Infer(), boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func (pq *BookRepository) Delete(bookID uuid.UUID) error {
	book, err := models.FindUserBook(context.Background(), pq.db, bookID.String())
	if err != nil {
		return err
	}
	_, err = book.Delete(context.Background(), pq.db)
	if err != nil {
		return err
	}
	return nil
}
