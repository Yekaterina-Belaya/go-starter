package models

type Vacancy struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Salary      string `json:"salary"`
}

// Пример функции для получения вакансий
func GetAllVacancies() []Vacancy {
	// Это пример, в реальном проекте здесь будет запрос к базе данных
	return []Vacancy{
		{ID: 1, Title: "Frontend Developer", Description: "Build cool websites", Salary: "50000"},
		{ID: 2, Title: "Backend Developer", Description: "Build APIs", Salary: "60000"},
	}
}

// Функция для создания вакансии
func CreateVacancy(vacancy Vacancy) {
	// В реальном проекте сюда нужно добавить логику для сохранения вакансии в базу данных
}
