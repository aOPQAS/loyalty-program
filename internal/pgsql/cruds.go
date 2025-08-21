package pgsql

import (
	"fmt"
	"microservice/pkg/models"
	"time"

	"github.com/google/uuid"
)

func (c *Client) GetProgram(programType string, name string, active bool) ([]models.Program, error) {
	dbSession := c.GetSession()

	var resp []models.Program
	stmt := dbSession.Select("*").From("programs")

	if programType != "" {
		stmt = stmt.Where("type = ?", programType)
	}

	if name != "" {
		stmt = stmt.Where("name = ?", name)
	}

	if active {
		stmt = stmt.Where("active = ?", active)
	}

	_, err := stmt.Load(&resp)
	if err != nil {
		return nil, fmt.Errorf("failed to get program: %w", err)
	}

	return resp, nil
}

func (c *Client) GetProgramBYID(id string) (*models.Program, error) {
	dbSession := c.GetSession()
	var resp models.Program
	stmt := dbSession.Select("*").From("programs")

	if id != "" {
		if u, err := uuid.Parse(id); err == nil {
			stmt = stmt.Where("id = ?", u)
		} else {
			stmt = stmt.Where("int_id = ?", id)
		}
	}

	_, err := stmt.Load(&resp)
	if err != nil {
		return nil, fmt.Errorf("failed to get program: %w", err)
	}

	return &resp, nil
}

func (c *Client) CreateProgram(p models.Program) (string, error) {
	dbSession := c.GetSession()

	createdID := uuid.New().String()
	now := float64(time.Now().Unix())

	stmt := dbSession.InsertInto("programs").Columns(
		"id", "type", "name", "image", "fixed_price", "total_services_cost", "discount_percent",
		"valid_until", "terms", "created_at", "updated_at", "active",
	).Values(
		createdID, p.Type, p.Name, p.Image, p.FixedPrice, p.TotalServicesCost, p.DiscountPercent,
		p.ValidUntil, p.Terms, now, now, p.Active,
	)

	if _, err := stmt.Exec(); err != nil {
		return "", fmt.Errorf("failed to create program: %w", err)
	}

	return createdID, nil
}

func (c *Client) UpdateProgram(p models.Program) error {
	dbSession := c.GetSession()

	stmt := dbSession.Update("programs").SetMap(map[string]interface{}{
		"type":                p.Type,
		"name":                p.Name,
		"image":               p.Image,
		"fixed_price":         p.FixedPrice,
		"total_services_cost": p.TotalServicesCost,
		"discount_percent":    p.DiscountPercent,
		"valid_until":         p.ValidUntil,
		"terms":               p.Terms,
		"updated_at":          float64(time.Now().Unix()),
		"active":              p.Active,
	}).Where("id = ?", p.ID)

	if _, err := stmt.Exec(); err != nil {
		return fmt.Errorf("failed to update program: %w", err)
	}

	return nil
}

func (c *Client) DeleteProgram(id string) error {
	dbSession := c.GetSession()

	stmt := dbSession.DeleteFrom("programs").Where("id = ?", id)

	if _, err := stmt.Exec(); err != nil {
		return fmt.Errorf("failed to delete program: %w", err)
	}

	return nil
}

func (c *Client) GetServices(name string) ([]models.Service, error) {
	dbSession := c.GetSession()

	var resp []models.Service
	stmt := dbSession.Select("*").From("services")

	if name != "" {
		stmt = stmt.Where("name = ?", name)
	}

	_, err := stmt.Load(&resp)
	if err != nil {
		return nil, fmt.Errorf("failed to get services: %w", err)
	}

	return resp, nil
}

func (c *Client) CreateServices(serv models.Service) (string, error) {
	dbSession := c.GetSession()

	createdID := uuid.New().String()

	stmt := dbSession.InsertInto("services").Columns(
		"service_id", "name", "tarif", "duration",
	).Values(
		createdID, serv.Name, serv.Tarif, serv.Duration,
	)

	if _, err := stmt.Exec(); err != nil {
		return "", fmt.Errorf("failed to create services: %w", err)
	}

	return createdID, nil
}

func (c *Client) UpdateServices(serv models.Service) error {
	dbSession := c.GetSession()

	stmt := dbSession.Update("services").SetMap(map[string]interface{}{
		"name":     serv.Name,
		"tarif":    serv.Tarif,
		"duration": serv.Duration,
	}).Where("service_id = ?", serv.ServiceID)

	if _, err := stmt.Exec(); err != nil {
		return fmt.Errorf("failed to update services: %w", err)
	}

	return nil
}

func (c *Client) DeleteServices(id string) error {
	dbSession := c.GetSession()

	stmt := dbSession.DeleteFrom("services").Where("service_id = ?", id)

	if _, err := stmt.Exec(); err != nil {
		return fmt.Errorf("failed to delete services: %w", err)
	}

	return nil
}

func (c *Client) GetProgramServices(programID string) ([]models.ProgramService, error) {
	dbSession := c.GetSession()

	var resp []models.ProgramService
	stmt := dbSession.Select("*").From("program_services")

	if programID != "" {
		stmt = stmt.Where("program_id = ?", programID)
	}

	_, err := stmt.Load(&resp)
	if err != nil {
		return nil, fmt.Errorf("failed to to get program services: %w", err)
	}

	return resp, nil
}

func (c *Client) CreateProgramService(ps models.ProgramService) error {
	dbSession := c.GetSession()

	stmt := dbSession.InsertInto("program_services").Columns(
		"program_id", "service_id",
	).Values(
		ps.ProgramID, ps.ServiceID,
	)

	if _, err := stmt.Exec(); err != nil {
		return fmt.Errorf("failed to create program services: %w", err)
	}

	return nil
}

// program_services — это таблица-связка, поэтому обычно не делают update

func (c *Client) DeleteProgramServices(programID, serviceID string) error {
	dbSession := c.GetSession()

	stmt := dbSession.DeleteFrom("program_services").Where("program_id = ?", programID).Where("service_id = ?", serviceID)

	if _, err := stmt.Exec(); err != nil {
		return fmt.Errorf("failed to delete program services: %w", err)
	}

	return nil
}
