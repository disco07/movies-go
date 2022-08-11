package main

type Student struct {
	ID         int64      `db:"id"`
	Lastname   string     `db:"lastname"`
	Firstname  string     `db:"firstname"`
	Email      string     `db:"email"`
	Convention Convention `db:"convention"`
}

type Convention struct {
	ID     int64  `db:"id"`
	Name   string `db:"name"`
	NbHour int
}

type Attestation struct {
	ID         int64      `db:"id"`
	Student    Student    `db:"student"`
	Convention Convention `db:"convention"`
	Message    string     `db:"message"`
}
