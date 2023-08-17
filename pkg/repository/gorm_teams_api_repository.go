package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/RBMarketplace/di-api-teams/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TODO, better error handling

type GormTeamRepository struct {
	db *gorm.DB
}

func NewGormTeamRepository() *GormTeamRepository {
	host := os.Getenv("DATABASE_URL")
	fmt.Println(host)
	dsn := fmt.Sprintf("host=%s user=postgres dbname=postgres port=5433 sslmode=disable TimeZone=Asia/Shanghai", host)
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Map the Team struct to the teams table in the database
	//   - If the table doesn't exist, it will create it
	//	 - If the table does exist, it will add any missing fields
	//   - If the table does exist, it will NOT remove any extra fields
	err = db.AutoMigrate(&domain.Team{})
	if err != nil {
		log.Fatalf("Failed to migrate the schema. %v", err)
	}

	return &GormTeamRepository{
		db: db,
	}
}

func (r *GormTeamRepository) GetTeam(id string) (domain.Team, error) {
	var team domain.Team
	if err := r.db.Where("id = ?", id).Find(&team).Error; err != nil {
		return domain.Team{}, err
	}
	return team, nil
}

func (r *GormTeamRepository) GetTeams() ([]domain.Team, error) {
	var teams []domain.Team
	if err := r.db.Find(&teams).Error; err != nil {
		return nil, err
	}
	return teams, nil
}

func (r *GormTeamRepository) AddTeam(team domain.Team) error {
	return r.db.Create(team).Error
}

func (r *GormTeamRepository) UpdateTeam(team domain.Team) error {
	return r.db.Save(team).Error
}

func (r *GormTeamRepository) RemoveTeam(id string) error {
	return r.db.Where("id = ?", id).Delete(&domain.Team{}).Error
}

func (r *GormTeamRepository) DatabaseAvailable() (bool, error) {
	sqlDB, err := r.db.DB()
	if err != nil {
		return false, fmt.Errorf("DB was not connectable: %v", err)
	}
	err = sqlDB.Ping()
	if err != nil {
		return false, fmt.Errorf("DB was not reachable: %v", err)
	}
	return true, nil
}
