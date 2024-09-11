package school

import privateschool "github.com/Eric-lab-star/school/internal/privateSchool"

func GetSchoolName() string {
	return "Lazy Dev School"
}

func GetPrivateSchoolName() string {
	private := privateschool.PrivateSchoolName()
	return private
}

func GetColor() string {
	return "Blue"
}
