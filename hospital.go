package hospital

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Hospital struct {
	Id        int64
	Name      string
	Staffs    []*Staff
	Patients  []*Patient
	Addresses []*Address
}

type Staff struct {
	Id          int64
	HospitalId  int64
	FullName    string
	PhoneNumber string
}

type Patient struct {
	Id          int64
	HospitalId  int64
	FullName    string
	PatientInfo string
	PhoneNumber string
}

type Address struct {
	Id         int64
	HospitalId int64
	Region     string
	Street     string
}

type Response struct {
	Hospitals []*Hospital
}

const (
	PostgresHost     = "localhost"
	PostgresPort     = 5432
	PostgresUser     = "postgres"
	PostgresPassword = "bnnfav"
	PostgresDatabase = "hospital_support"
)

func Hosp() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", PostgresHost, PostgresPort, PostgresUser, PostgresPassword, PostgresDatabase)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error opening database connection: ", err)
		return
	}

	defer db.Close()

	// INSERT
	create(db)

	// SELECT
	// get(db)

	// UPDATE
	// update(db, 7)

	// DELETE
	// delete(db, 7)

}

func create(db *sql.DB) {

	hospital := []Hospital{
		{
			Name: "Akfa",
			Staffs: []*Staff{
				{
					FullName:    "Javohir",
					PhoneNumber: "999999999",
				},
				{
					FullName:    "Nilufar",
					PhoneNumber: "988888888",
				},
			},
			Patients: []*Patient{
				{
					FullName:    "Abdulloh Abdullayev",
					PatientInfo: "Ahvoli yaxshi",
					PhoneNumber: "333333333",
				},
				{
					FullName:    "Baxrom Kunduziy",
					PatientInfo: "Ahvoli o'rta",
					PhoneNumber: "344444444",
				},
				{
					FullName:    "Selena Gomes",
					PatientInfo: "Dipressiyada",
					PhoneNumber: "355555555",
				},
			},
			Addresses: []*Address{
				{
					Region: "Las-Vegas",
					Street: "Time square",
				},
				{
					Region: "Chilonzor",
					Street: "Metro",
				},
			},
		},
	}

	for _, hospital := range hospital {
		var hospitalId int64
		queryHospital := `
			INSERT INTO
			    hospital(name)
			VALUES($1)
				RETURNING id`

		err := db.QueryRow(queryHospital, hospital.Name).Scan(&hospitalId)
		if err != nil {
			fmt.Println("error inserting hospital: ", err)
			return
		}

		for _, staff := range hospital.Staffs {
			queryStaff := `
				INSERT INTO
					staff(hospital_id, full_name, phone_number)
				VALUES
					($1, $2, $3)`

			_, err = db.Exec(queryStaff, hospitalId, staff.FullName, staff.PhoneNumber)
			if err != nil {
				fmt.Println("error inserting staff: ", err)
				return
			}
		}

		for _, patient := range hospital.Patients {
			queryPatient := `
				INSERT INTO
					patients(hospital_id, full_name, patient_info, phone_number)
				VALUES
					($1, $2, $3, $4)`

			_, err = db.Exec(queryPatient, hospitalId, patient.FullName, patient.PatientInfo, patient.PhoneNumber)
			if err != nil {
				fmt.Println("error inserting patient: ", err)
				return
			}
		}

		for _, address := range hospital.Addresses {
			queryAddress := `
				INSERT INTO
					addresses(hospital_id, region, street)
				VALUES
					($1, $2, $3)`

			_, err = db.Exec(queryAddress, hospitalId, address.Region, address.Street)
			if err != nil {
				fmt.Println("error inserting address: ", err)
				return
			}
		}
	}

	fmt.Println("All things inserted")
}

func get(db *sql.DB) {
	resp := Response{}

	queryHospital := `
		SELECT
			id, name
		FROM 
			hospital`
	rowHospital, err := db.Query(queryHospital)
	if err != nil {
		fmt.Println("error selecting hospital: ", err)
		return
	}

	for rowHospital.Next() {
		hospital := Hospital{}
		err = rowHospital.Scan(
			&hospital.Id,
			&hospital.Name,
		)
		if err != nil {
			fmt.Println("error scanning hospital: ", err)
			return
		}

		queryStaff := `
			SELECT
				id, hospital_id, full_name, phone_number
			FROM
				staff
			WHERE
			    hospital_id = $1`

		rowStaff, err := db.Query(queryStaff, hospital.Id)
		if err != nil {
			fmt.Println("error selecting staff: ", err)
			return
		}

		for rowStaff.Next() {
			staff := Staff{}
			err = rowStaff.Scan(
				&staff.Id,
				&staff.HospitalId,
				&staff.FullName,
				&staff.PhoneNumber,
			)
			if err != nil {
				fmt.Println("error scanning staff: ", err)
				return
			}

			hospital.Staffs = append(hospital.Staffs, &staff)
		}

		queryPatient := `
			SELECT
				id, hospital_id, full_name, patient_info, phone_number
			FROM
				patients
			WHERE hospital_id = $1`

		rowPatient, err := db.Query(queryPatient, hospital.Id)
		if err != nil {
			fmt.Println("error selecting patient: ", err)
			return
		}

		for rowPatient.Next() {
			patient := Patient{}
			err = rowPatient.Scan(
				&patient.Id,
				&patient.HospitalId,
				&patient.FullName,
				&patient.PatientInfo,
				&patient.PhoneNumber,
			)
			if err != nil {
				fmt.Println("error scanning patient: ", err)
				return
			}

			hospital.Patients = append(hospital.Patients, &patient)
		}

		queryAddress := `
			SELECT
				id, hospital_id, region, street
			FROM
				addresses
			WHERE hospital_id = $1`

		rowAddress, err := db.Query(queryAddress, hospital.Id)
		if err != nil {
			fmt.Println("error selecting address: ", err)
			return
		}

		for rowAddress.Next() {
			address := Address{}
			err = rowAddress.Scan(
				&address.Id,
				&address.HospitalId,
				&address.Region,
				&address.Street,
			)
			if err != nil {
				fmt.Println("error scanning address: ", err)
				return
			}

			hospital.Addresses = append(hospital.Addresses, &address)
		}

		resp.Hospitals = append(resp.Hospitals, &hospital)
	}
	c_hospital := 1
	for _, hospital := range resp.Hospitals {
		c_staff, c_patient, c_address := 1, 1, 1
		fmt.Printf("%d - hospital\n", c_hospital)
		fmt.Printf("ID: %d\n", hospital.Id)
		fmt.Printf("Name: %s\n", hospital.Name)
		fmt.Println("Staffs:")
		for _, staff := range hospital.Staffs {
			fmt.Printf("\t%d - staff\n", c_staff)
			fmt.Printf("\tID: %d\n", staff.Id)
			fmt.Printf("\tHospital ID: %d\n", staff.HospitalId)
			fmt.Printf("\tFull Name: %s\n", staff.FullName)
			fmt.Printf("\tPhone Number: %s\n\n", staff.PhoneNumber)
			c_staff++
		}
		fmt.Println("Patients:")
		for _, patient := range hospital.Patients {
			fmt.Printf("\t%d - patient\n", c_patient)
			fmt.Printf("\tID: %d\n", patient.Id)
			fmt.Printf("\tHospital ID: %d\n", patient.HospitalId)
			fmt.Printf("\tFull Name: %s\n", patient.FullName)
			fmt.Printf("\tPatient Info: %s\n", patient.PatientInfo)
			fmt.Printf("\tPhone Number: %s\n\n", patient.PhoneNumber)
			c_patient++
		}
		fmt.Println("Addresses:")
		for _, address := range hospital.Addresses {
			fmt.Printf("\t%d - address\n", c_address)
			fmt.Printf("\tID: %d\n", address.Id)
			fmt.Printf("\tHospital ID: %d\n", address.HospitalId)
			fmt.Printf("\tRegion: %s\n", address.Region)
			fmt.Printf("\tStreet: %s\n\n", address.Street)
			c_address++
		}
		c_hospital++
	}
}

func update(db *sql.DB, hospital_id int64) {
	new := []Hospital{
		{
			Name: "new_name",
			Staffs: []*Staff{
				{
					FullName:    "new staff 1",
					PhoneNumber: "000000000",
				},
				{
					FullName:    "new staff 2",
					PhoneNumber: "000000001",
				},
			},
			Patients: []*Patient{
				{
					FullName:    "new patient 1",
					PatientInfo: "Ahvoli yomooooon",
					PhoneNumber: "000000000222222",
				},
				{
					FullName:    "new patient 2",
					PatientInfo: "ahvoli yomooooon",
					PhoneNumber: "000000000022222",
				},
				{
					FullName:    "new patient 3",
					PatientInfo: "ahvoli yomooooon",
					PhoneNumber: "00000000022222",
				},
			},
			Addresses: []*Address{
				{
					Region: "new address region 1",
					Street: "new address street 1",
				},
				{
					Region: "new address region 2",
					Street: "new address street 2",
				},
			},
		},
	}

	for _, hospital := range new {
		queryHospital := `
			    UPDATE 
					hospital
				SET 
					name = $1
				WHERE 
					id = $2`

		_, err := db.Exec(queryHospital, hospital.Name, hospital_id)
		if err != nil {
			fmt.Println("error updating hospital: ", err)
			return
		}

		for _, staff := range hospital.Staffs {
			queryStaff := `
				UPDATE
					staff
				SET
					full_name = $1, phone_number = $2
				WHERE
					hospital_id = $3`

			_, err = db.Exec(queryStaff, staff.FullName, staff.PhoneNumber, hospital_id)
			if err != nil {
				fmt.Println("error updating staff: ", err)
				return
			}
		}

		for _, patient := range hospital.Patients {
			queryPatient := `
				UPDATE
					patients
				SET
					full_name = $1, patient_info = $2, phone_number = $3
				WHERE
					hospital_id = $4`

			_, err = db.Exec(queryPatient, patient.FullName, patient.PatientInfo, patient.PhoneNumber, hospital_id)
			if err != nil {
				fmt.Println("error updating patient: ", err)
				return
			}
		}

		for _, address := range hospital.Addresses {
			queryAddress := `
				UPDATE
					addresses
				SET
					region = $1, street = $2
				WHERE
					hospital_id = $3`

			_, err = db.Exec(queryAddress, address.Region, address.Street, hospital_id)
			if err != nil {
				fmt.Println("error updating address: ", err)
				return
			}
		}
	}

	fmt.Println("All things updated")
}

func delete(db *sql.DB, hospital_id int64) {
	queryStaff := `
        DELETE
        FROM
			staff
		WHERE
			hospital_id = $1`

	_, err := db.Exec(queryStaff, hospital_id)
	if err != nil {
		fmt.Println("error deleting staff: ", err)
		return
	}

	queryPatient := `
        DELETE
        FROM
			patients
		WHERE
			hospital_id = $1`

	_, err = db.Exec(queryPatient, hospital_id)
	if err != nil {
		fmt.Println("error deleting patient: ", err)
		return
	}

	queryAddress := `
        DELETE
        FROM
			addresses
		WHERE
			hospital_id = $1`

	_, err = db.Exec(queryAddress, hospital_id)
	if err != nil {
		fmt.Println("error deleting address: ", err)
		return
	}
	
	queryHospital := `
        DELETE
        FROM
			hospital
		WHERE
			id = $1`

	_, err = db.Exec(queryHospital, hospital_id)
	if err != nil {
		fmt.Println("error deleting hospital: ", err)
		return
	}

	fmt.Println("Hospital deleted")
}
