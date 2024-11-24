package repository

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/mcfiet/goDo/draw/model"
	userModel "github.com/mcfiet/goDo/user/model"
	"gorm.io/gorm"
)

type DrawRepository interface {
	GetAllDraws() ([]model.DrawResult, error)
	CreateDraw(draw *model.DrawResult) error
	GetDrawByGiverId(id uuid.UUID) (model.DrawResult, error)
}

type drawRepository struct {
	db *gorm.DB
}

func NewDrawRepository(db *gorm.DB) DrawRepository {
	return &drawRepository{db}
}

func (repo *drawRepository) GetRandomUser() (*userModel.User, error) {
	var user *userModel.User
	err := repo.db.Order("RANDOM()").First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *drawRepository) GetRandomUserExcluding(giverId uuid.UUID, excludedIds []uuid.UUID) (*userModel.User, error) {
	// TODO: Implement the GetRandomUserExcluding function
	var user userModel.User

	// Sicherstellen, dass excludedIds nicht leer ist
	if len(excludedIds) == 0 {
		excludedIds = append(excludedIds, uuid.Nil)
	}

	// Raw SQL-Abfrage mit UUID Array und korrekt umgewandeltem excludedIds
	query := `
		SELECT * FROM users 
		WHERE id != ? 
		AND id != ALL (?::uuid[])  -- Sicherstellen, dass excludedIds als UUID-Array behandelt werden
		AND id NOT IN (SELECT receiver_id FROM draw_results)
		ORDER BY RANDOM() LIMIT 1
	`

	// Übergebe excludedIds als uuid-Array, nicht als Text
	err := repo.db.Raw(query, giverId, excludedIds).Scan(&user).Error
	if err != nil {
		return nil, fmt.Errorf("Fehler beim Abrufen eines zufälligen Benutzers: %w", err)
	}

	// Überprüfen, ob ein Benutzer gefunden wurde
	if user.ID == uuid.Nil {
		return nil, fmt.Errorf("Kein Benutzer gefunden, der den Kriterien entspricht")
	}

	return &user, nil
}

func (repo *drawRepository) CheckIfDrawExists(giverId uuid.UUID, recieverId uuid.UUID) bool {
	var count int64
	err := repo.db.Model(&model.DrawResult{}).Where("giver_id = ? AND receiver_id = ?", giverId, recieverId).Count(&count).Error
	if err != nil {
		log.Fatal(err)
	}
	return count > 0
}

func (repo *drawRepository) CheckIfDrawExistsByGiverId(id uuid.UUID) bool {
	var count int64
	err := repo.db.Model(&model.DrawResult{}).Where("giver_id = ?", id).Count(&count).Error
	if err != nil {
		log.Fatal(err)
	}
	return count > 0
}

func (repo *drawRepository) GetAllDraws() ([]model.DrawResult, error) {
	var draws []model.DrawResult
	result := repo.db.Find(&draws)

	return draws, result.Error
}

func (repo *drawRepository) CreateDraw(draw *model.DrawResult) error {
	if repo.CheckIfDrawExistsByGiverId(draw.GiverId) {
		return fmt.Errorf("Du hast schon jemanden gezogen!")
	}

	var excludedIds []uuid.UUID
	draws, _ := repo.GetAllDraws()
	for _, draw := range draws {
		excludedIds = append(excludedIds, draw.ReceiverId)
	}

	maxRetries := 100

	for attempts := 0; attempts < maxRetries; attempts++ {

		reciever, err := repo.GetRandomUserExcluding(draw.GiverId, excludedIds)
		if err != nil {
			return err
		}
		if draw.GiverId == reciever.ID {
			continue
		}

		if repo.CheckIfDrawExists(draw.GiverId, reciever.ID) {
			continue
		}

		draw.ReceiverId = reciever.ID
		err = repo.db.Create(draw).Error
		if err != nil {
			return err
		}

		return nil
	}
	return fmt.Errorf("Auslosen fehlgeschlagen, versuche es noch einmal")
}

func (repo *drawRepository) GetDrawByGiverId(id uuid.UUID) (model.DrawResult, error) {
	var draw model.DrawResult

	result := repo.db.Where("giver_id = ?", id).First(&draw)
	return draw, result.Error
}

func (repo *drawRepository) UpdateDrawById(draw model.DrawResult) error {
	result := repo.db.Model(&draw).Updates(draw)

	return result.Error
}
