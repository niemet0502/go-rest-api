package services

import (
	"go-server/entities"
	"go-server/repositories"
)

type GameService struct {
	reposity *repositories.GameRepository
}

func CreateNewGameService() *GameService {
	return &GameService{reposity: repositories.NewGameRepository()}
}

func (gs *GameService) Create(newGame *entities.Game) error {
	return gs.reposity.CreateGame(newGame)
}

func (gs *GameService) FindAll() []entities.Game {
	return gs.reposity.GetGames()
}

func (gs *GameService) GetById(id int) (entities.Game, error) {
	return gs.reposity.GetGameById(id)
}
