package postgres

import (
	"context"
	"develop/config"
	"develop/storage"
	"fmt"
	 "strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database"          //database is needed for migration
	_ "github.com/golang-migrate/migrate/v4/database/postgres" //postgres is used for database
	_ "github.com/golang-migrate/migrate/v4/source/file"       //file is needed for migration url
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

type Store struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, cfg config.Config)(storage.IStorage, error){
	url := fmt.Sprintf(
		`postgres://%s:%s@%s:%s/%s?sslmode=disable`,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
	)

	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil{
		fmt.Println("Error while parsing config!", err.Error())
		return Store{}, err
	}

	poolConfig.MaxConns = 100

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil{
		fmt.Println("Error while connecting to db!", err.Error())
		return Store{},err
	}

	//migration
	m, err := migrate.New("file://migrations",  url)
	if err != nil {
		fmt.Println("Error while migrating!",err.Error())
		return nil, err
	}

	err = m.Up()
	if err != nil {
		fmt.Println("Error while making db up!",err.Error())
		if !strings.Contains(err.Error(), "no change") {
			version, dirty, err := m.Version()
			if err != nil {
				fmt.Println("Error!")
				return nil, err
			}

			if dirty {
				version--
				err = m.Force(int(version))
				if err != nil {
					fmt.Println("Error while forcing db!")
					return nil, err
				}
			}
			return nil, err
		}
	}

	return Store{
		pool: pool,
	}, nil
}

func (s Store) Close() {
	s.pool.Close()
}

func (s Store) Book()storage.IBookStorage{
	return NewBookRepo(s.pool)
}