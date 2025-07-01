package repository

import (
	"database/sql"
	"emailn/internal/campaign/domain/entity"
	"emailn/internal/core/infra/database"
	"emailn/internal/shared/exception"
	sharedvo "emailn/internal/shared/valueobject"
	"errors"
	"fmt"
	"log"
)

type PostgresRepository struct {
	DbInterface database.DatabaseInterface
}

func (pr *PostgresRepository) Create(entity *entity.Campaign) (uid *sharedvo.UID, err error) {
	tx, err := database.Begin()
	if err != nil {
		err = exception.NewHttpInternalServerException(fmt.Sprintf("Error when begin transaction | Error: %s", err.Error()))
		return
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	if err = createContacts(entity.Contacts, pr.DbInterface, tx); err != nil {
		err = exception.NewHttpInternalServerException(fmt.Sprintf("Error when registering contact | Error: %s", err.Error()))
		return
	}

	query := "INSERT INTO campaigns(campaign_id, name, content, status, created_at) VALUES($1, $2, $3, $4, $5)"
	_, err = pr.DbInterface.Exec(tx, query, entity.ID, entity.Name, entity.Content, entity.Status, entity.CreatedAt)
	if err != nil {
		err = exception.NewHttpInternalServerException(fmt.Sprintf("Error when registering campaign | Error: %s", err.Error()))
		return
	}

	contactsRegistered := make(map[sharedvo.UID]struct{}, len(entity.Contacts))

	for _, contact := range entity.Contacts {
		if _, exists := contactsRegistered[*contact.ID]; exists {
			continue
		}

		query := "INSERT INTO campaigns_contacts(campaign_id, contact_id, created_at) VALUES($1,$2,$3)"
		_, err = pr.DbInterface.Exec(tx, query, entity.ID, contact.ID, entity.CreatedAt)
		if err != nil {
			err = exception.NewHttpInternalServerException(fmt.Sprintf("Error when registering contact | Error: %s", err.Error()))
			return
		}

		contactsRegistered[*contact.ID] = struct{}{}
	}

	return entity.ID, nil
}

func createContacts(contacts []*entity.Contact, db database.DatabaseInterface, tx *sql.Tx) error {
	for _, contact := range contacts {
		if contactId := getContactIdIfExists(contact, db); contactId != nil {
			contact.ID = contactId
			continue
		}

		query := "INSERT INTO contacts(contact_id, first_name, last_name, phone, email, created_at) VALUES($1,$2,$3,$4,$5,$6)"
		_, err := db.Exec(tx, query, contact.ID, contact.FirstName, contact.LastName, contact.Phone, contact.Email, contact.CreatedAt)
		if err != nil {
			return err
		}
	}

	return nil
}

func getContactIdIfExists(contact *entity.Contact, db database.DatabaseInterface) *sharedvo.UID {
	query := "SELECT contact_id FROM contacts WHERE phone = $1 AND email = $2"

	var id string
	row := db.QueryRow(query, contact.Phone, contact.Email)
	err := row.Scan(&id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}

		log.Printf("Erro ao buscar contato: %v", err)
		return nil
	}

	uid := sharedvo.UID(id)
	return &uid
}

func (pr *PostgresRepository) GetAll() ([]*entity.Campaign, error) {
	var campaigns []*entity.Campaign

}

func (pr *PostgresRepository) GetCampaignById(campaignId string) (*entity.Campaign, error) {

	return nil, nil
}
