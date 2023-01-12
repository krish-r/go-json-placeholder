package main

var users = Users{
	{
		Id:       1,
		Name:     "User 1",
		Username: "User_1",
		Email:    "user-1@example.com",
		Address: UserAddress{
			Street:  "imaginary street",
			Suite:   "Apt. 001",
			City:    "imaginary city",
			Zipcode: "12345-1111",
			Geo: Geo{
				Lat: "-99.9999",
				Lng: "99.9999",
			},
		},
		Phone:   "9-999-999-9999 x99999",
		Website: "example.com",
		Company: UserCompany{
			Name:        "Not-A-Real-Company",
			CatchPhrase: "As the name suggests, Not-A-Real-Company",
			Bs:          "Not A Real Company",
		},
	},
	{
		Id:       2,
		Name:     "User 2",
		Username: "User_2",
		Email:    "user-02@example.com",
		Address: UserAddress{
			Street:  "imaginary street",
			Suite:   "Apt. 002",
			City:    "imaginary city",
			Zipcode: "12345-1111",
			Geo: Geo{
				Lat: "-99.9999",
				Lng: "99.9999",
			},
		},
		Phone:   "9-999-999-9999 x88888",
		Website: "example.com",
		Company: UserCompany{
			Name:        "Not-A-Real-Company",
			CatchPhrase: "As the name suggests, Not-A-Real-Company",
			Bs:          "Not A Real Company",
		},
	},
	{
		Id:       3,
		Name:     "User 3",
		Username: "User_3",
		Email:    "user-03@example.com",
		Address: UserAddress{
			Street:  "imaginary street",
			Suite:   "Apt. 003",
			City:    "imaginary city",
			Zipcode: "12345-1111",
			Geo: Geo{
				Lat: "-99.9999",
				Lng: "99.9999",
			},
		},
		Phone:   "9-999-999-9999 x77777",
		Website: "example.com",
		Company: UserCompany{
			Name:        "Not-A-Real-Company",
			CatchPhrase: "As the name suggests, Not-A-Real-Company",
			Bs:          "Not A Real Company",
		},
	},
}

var posts = Posts{
	{
		UserId: 1,
		Id:     1,
		Title:  "post1-title1",
		Body:   "post1-body1",
	},
	{
		UserId: 1,
		Id:     2,
		Title:  "post2-title1",
		Body:   "post2-body1",
	},
	{
		UserId: 2,
		Id:     3,
		Title:  "post3-title1",
		Body:   "post3-body1",
	},
	{
		UserId: 2,
		Id:     4,
		Title:  "post4-title1",
		Body:   "post4-body1",
	},
	{
		UserId: 2,
		Id:     5,
		Title:  "post5-title1",
		Body:   "post5-body1",
	},
	{
		UserId: 2,
		Id:     6,
		Title:  "post6-title1",
		Body:   "post6-body1",
	},
	{
		UserId: 3,
		Id:     7,
		Title:  "post7-title1",
		Body:   "post7-body1",
	},
	{
		UserId: 3,
		Id:     8,
		Title:  "post8-title1",
		Body:   "post8-body1",
	},
}

var comments = Comments{
	{
		PostId: 1,
		Id:     1,
		Name:   "subscriber-1",
		Email:  "subscriber1@example.com",
		Body:   "post1-comment1",
	},
	{
		PostId: 1,
		Id:     2,
		Name:   "subscriber-2",
		Email:  "subscriber2@example.com",
		Body:   "post1-comment2",
	},
	{
		PostId: 2,
		Id:     3,
		Name:   "subscriber-1",
		Email:  "subscriber1@example.com",
		Body:   "post2-comment1",
	},
	{
		PostId: 2,
		Id:     4,
		Name:   "subscriber-3",
		Email:  "subscriber3@example.com",
		Body:   "post2-comment2",
	},
	{
		PostId: 3,
		Id:     5,
		Name:   "subscriber-2",
		Email:  "subscriber2@example.com",
		Body:   "post3-comment1",
	},
	{
		PostId: 4,
		Id:     6,
		Name:   "subscriber-4",
		Email:  "subscriber4@example.com",
		Body:   "post4-comment1",
	},
	{
		PostId: 5,
		Id:     7,
		Name:   "subscriber-3",
		Email:  "subscriber3@example.com",
		Body:   "post5-comment1",
	},
	{
		PostId: 6,
		Id:     8,
		Name:   "subscriber-4",
		Email:  "subscriber4@example.com",
		Body:   "post6-comment1",
	},
	{
		PostId: 6,
		Id:     9,
		Name:   "subscriber-2",
		Email:  "subscriber2@example.com",
		Body:   "post6-comment1",
	},
	{
		PostId: 7,
		Id:     10,
		Name:   "subscriber-1",
		Email:  "subscriber1@example.com",
		Body:   "post7-comment1",
	},
	{
		PostId: 8,
		Id:     11,
		Name:   "subscriber-1",
		Email:  "subscriber1@example.com",
		Body:   "post8-comment1",
	},
	{
		PostId: 8,
		Id:     12,
		Name:   "subscriber-2",
		Email:  "subscriber2@example.com",
		Body:   "post8-comment2",
	},
}
