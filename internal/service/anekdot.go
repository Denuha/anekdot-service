package service

import (
	"context"
	"errors"

	"github.com/Denuha/anekdot-service/internal/models"
	"github.com/Denuha/anekdot-service/internal/repository"
	"github.com/Denuha/anekdot-service/internal/service/parser"
)

type anekdot struct {
	anekDB repository.AnekdotDB
}

func (a *anekdot) ParseAnekdots(ctx context.Context, source string) (int, error) {
	var parserClient parser.Parser

	switch source {
	case "anekdotme":
		pAnekdotme := parser.NewParserAnekdotme()
		ps := parser.NewParserService(pAnekdotme)
		parserClient = ps.Parser
	default:
		return 0, errors.New("source is wrong")
	}

	anekdots, err := parserClient.ParseAnekdots()
	if err != nil {
		return 0, err
	}
	if len(anekdots) == 0 {
		return 0, errors.New("parse 0 anekdots")
	}

	err = a.anekDB.InsertAnekdotList(ctx, anekdots)
	if err != nil {
		return 0, err
	}

	return len(anekdots), nil
}

func (a *anekdot) GetRandomAnekdot(ctx context.Context) (*models.Anekdot, error) {
	anekdot, err := a.anekDB.GetRandomAnekdot(ctx)
	if err != nil {
		return nil, err
	}
	return anekdot, nil
}

func (a *anekdot) UpdateRating(ctx context.Context, anekdotID int, metod repository.MethodRaitng) error {
	err := a.anekDB.ChangeRating(ctx, anekdotID, metod)
	if err != nil {
		return err
	}
	return nil
}

func (a *anekdot) GetAnekdotByID(ctx context.Context, anekdotID int) (*models.Anekdot, error) {
	return a.anekDB.GetAnekdotByID(ctx, anekdotID)
}

func NewAnekdotService(repos *repository.Repositories) Anekdot {
	return &anekdot{
		anekDB: repos.AnekdotDB,
	}
}
