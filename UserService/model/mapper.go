package model

func (user *User) ToUserDTO() UserDto {

	return UserDto{
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Username:     user.Username,
		EmailAddress: user.EmailAddress,
		Role:         string(user.Role),
		Address:      user.Address,
		Banned:       user.Banned,
	}
}

func (user *User) ToRepairmanDto() RepairmanDto {

	return RepairmanDto{
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Username:      user.Username,
		EmailAddress:  user.EmailAddress,
		Role:          string(user.Role),
		Address:       user.Address,
		Price:         user.Price,
		Occupation:    user.Occupation,
		Review:        user.Review,
		ReviewCounter: user.ReviewCounter,
	}
}
