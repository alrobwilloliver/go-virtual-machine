package handledb

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlStore struct {
	*sql.DB
}

type VirtualMachines []Machine

type Machine struct {
	ID              int
	OperatingSystem string
	Owner           string
}

func (db *MySqlStore) AddVirtualMachine(m Machine) {
	_, err := db.Exec("INSERT INTO Machine (operatingSystem, owner) VALUES (?,?)", m.OperatingSystem, m.Owner)
	if err != nil {
		log.Fatal(err)
	}
}

func getVirtualMachineByOS(db *sql.DB, os string) (VirtualMachines, error) {
	rows, err := db.Query("SELECT * FROM Machine WHERE operatingSystem = (?)", os)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	virtualms := VirtualMachines{}
	for rows.Next() {
		vmachine := Machine{}
		err := rows.Scan(&vmachine.OperatingSystem, &vmachine.ID, &vmachine.Owner)
		if err != nil {
			virtualms = append(virtualms, vmachine)
		}
	}

	err = rows.Err()
	return virtualms, err
}

func (db *MySqlStore) GetVirtualMachineById(ctx context.Context, id int) (Machine, error) {
	row := db.QueryRowContext(ctx, "SELECT * from Machine WHERE id = (?)", id)
	vm := Machine{}
	err := row.Scan(&vm.ID, &vm.Owner, &vm.OperatingSystem)
	if err != nil {
		log.Fatal(err)
	}
	return vm, nil
}

func Connect() (*MySqlStore, error) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/VirtualMachine?parseTime=true")
	if err != nil {
		log.Fatal("Error occurred connecting to the database \n")
	}

	return &MySqlStore{
		DB: db,
	}, nil
}
