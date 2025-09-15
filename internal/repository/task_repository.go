package repository

import (
	"database/sql"
	"fmt"
	"task-manager/internal/model"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

// save task method
func (repository *TaskRepository) Save(task model.Task) (model.Task, error) {
	sqlStatement := `INSERT INTO tasks (title, description, status) VALUES (?, ?, ?)`
	
	result, err := repository.db.Exec(sqlStatement, task.Title, task.Description, "pending")
	if err != nil {
		return task, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return task, err
	}

	task.ID = int(lastInsertID)
	return task, nil
}

// get all method
func (repository *TaskRepository) FindAll()([]model.Task, error) {
	sqlStatement := `SELECT id, title, description, status, created_at, updated_at FROM tasks`

	rows, err := repository.db.Query(sqlStatement) // query digunakan khusus untuk query SELECT 
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []model.Task // variabel kosong untuk menyimpan tasks yang di looping
	
	// looping untuk cek data yang ada di database kemudian di append ke keranjang atau variabel tasks
	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	// kembalikan hasil variable tasks atau keranjang yang sudah berisi data dari hasil lopping
	return tasks, nil
}

// delete method
func (repository *TaskRepository) DeleteByID (id int) error {
	sqlStatement := `DELETE FROM tasks WHERE id = ?`

	result, err := repository.db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}

	rowsAffcted, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffcted == 0 {
		return fmt.Errorf("task dengan ID %d tidak ditemukan", id)
	}

	return nil
}

// search by ID
func (repository *TaskRepository) FindByID(id int) (model.Task, error) {
	var task model.Task // variabel penampung

	sqlStatement := `SELECT id, tittle, description, status, created_at, updated_at FROM tasks WHERE id= ?`

	row := repository.db.QueryRow(sqlStatement, id) // gunakan queryrow karena untuk mengembalikan satu objek "row"
	
	err := row.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return task, err
	}
	
	return task, nil

}