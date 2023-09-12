package users

import "gorm.io/gorm"

type Repository interface {
	Save(user Users) (Users, error)
	FindByEmail(email string) (Users, error)
	FindByID(id int) (Users, error)
	Update(user Users) (Users, error)
}

type repository struct{
	db *gorm.DB
}	

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}


func (r *repository) Save(user Users) (Users, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (Users, error)  {
	var user Users
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByID(ID int) (Users, error)  {
	var user Users
	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Update(user Users) (Users, error)  {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}