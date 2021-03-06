package category

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

// Repository struct for db connection
type Repository struct {
	Conn *sql.DB
}

// Category struct
type Category struct {
	ID          int64     `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Validate : interface for JSON backend validation
func (c *Category) Validate() error {

	if c.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	if len(c.Name) > 255 {
		return fmt.Errorf("name cannot be longer than 255 characters")
	}

	return nil

}

// SelectCategoryByID retrieves a product using its id
func (repository *Repository) SelectCategoryByID(id int64) (*Category, error) {
	row := repository.Conn.QueryRow("SELECT c.id, c.name, c.description, "+
		"c.created_at, c.updated_at FROM category c WHERE c.id=(?)", id)
	var name, description string
	var createdAt, updatedAt time.Time
	switch err := row.Scan(&id, &name, &description, &createdAt, &updatedAt); err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		category := Category{
			ID:          id,
			Name:        name,
			Description: description,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		}
		return &category, nil
	default:
		return nil, nil
	}
}

type filterName string

const filterByDateOfUpdate filterName = "updated_at"

// retrieveAllCategories stored in db
func (repository *Repository) retrieveAllCategories(filters map[filterName]interface{}) ([]*Category, error) {
	queryOrders := make([]string, 0)

	// Filtering categories by date
	if v, ok := filters[filterByDateOfUpdate]; ok {
		if vv, ok := v.(string); ok {
			switch vv {
			case "asc":
				queryOrders = append(queryOrders, "updated_at ASC LIMIT 3")
			case "desc":
				queryOrders = append(queryOrders, "updated_at DESC LIMIT 3")
			}
		}
	}

	queryFields := []string{
		"c.id", "c.name", "c.description", "c.created_at", "c.updated_at",
	}
	query := fmt.Sprintf("SELECT %s FROM category c", strings.Join(queryFields, ", "))

	if len(queryOrders) > 0 {
		query += fmt.Sprintf("\nORDER BY %s", strings.Join(queryOrders, ", "))
	}

	rows, err := repository.Conn.Query(query)

	if err != nil {
		return nil, err
	}

	var id int64
	var name, description string
	var createdAt, updatedAt time.Time
	var categories []*Category
	for rows.Next() {
		err := rows.Scan(&id, &name, &description, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &Category{
			ID:          id,
			Name:        name,
			Description: description,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		})
	}

	return categories, nil
}

// insertCategory posts a new category
func (repository *Repository) insertCategory(category *Category) error {
	stmt, err := repository.Conn.Prepare("INSERT INTO category(name, description, created_at," +
		" updated_at) VALUES(?,?,?,?)")

	if err != nil {
		return err
	}
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()
	res, errExec := stmt.Exec(category.Name, category.Description, category.CreatedAt, category.UpdatedAt)

	if errExec != nil {
		return errExec
	}

	lastInsertedID, errInsert := res.LastInsertId()

	if errInsert != nil {
		return errInsert
	}

	category.ID = lastInsertedID

	return nil
}

// updateCategory by ID
func (repository *Repository) updateCategory(category *Category, id int64) error {
	stmt, err := repository.Conn.Prepare("UPDATE category SET name=(?), description=(?), " +
		"updated_at=(?) WHERE id=(?)")
	if err != nil {
		return err
	}
	var createdAt time.Time
	row := repository.Conn.QueryRow("SELECT c.created_at FROM category c WHERE c.id=(?)", id)
	if err := row.Scan(&createdAt); err != nil {
		return err
	}
	category.CreatedAt = createdAt
	category.UpdatedAt = time.Now()

	_, errExec := stmt.Exec(category.Name, category.Description, category.UpdatedAt, id)

	if errExec != nil {
		return errExec
	}

	category.ID = id

	return nil
}

// deleteCategory by ID
func (repository *Repository) deleteCategory(id int64) (int64, error) {

	res, err := repository.Conn.Exec("DELETE FROM category WHERE id=(?)", id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
