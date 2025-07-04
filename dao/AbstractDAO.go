package dao

type AbstractDAO struct {
	FileName  string
	TableName string
}

type CustomQueryModel struct {
	Page    string
	Limit   string
	Keyword string
	Id      string
}
