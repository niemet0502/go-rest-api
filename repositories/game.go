package repositories

import (
	"errors"
	"go-server/database"
	"go-server/entities"
	"gorm.io/gorm"
	"log"
)

type GameRepository struct {
	DB *gorm.DB
}

func NewGameRepository() *GameRepository {
	return &GameRepository{DB: database.GetDB()}
}

func (gr *GameRepository) GetGames() []entities.Game {
	var games []entities.Game
	if err := gr.DB.Find(&games).Error; err != nil {
		log.Fatal(err)
	}

	return games
}

func (gr *GameRepository) CreateGame(newGame *entities.Game) error {
	return gr.DB.Create(newGame).Error
}

func (gr *GameRepository) GetGameById(id int) (entities.Game, error) {
	var game entities.Game
	result := gr.DB.First(&game, id)

	if result.Error != nil {
		log.Println("Game not found")
		return entities.Game{}, errors.New("Game not found")
	}
	return game, nil
}
