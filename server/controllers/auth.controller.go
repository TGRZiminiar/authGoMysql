package controllers

import (
	"template-go-auth-mysql/database"
	"template-go-auth-mysql/utils"
	"time"

	"template-go-auth-mysql/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Name            string `validate:"required" json:"name"`
	Email           string `validate:"required" json:"email"`
	Password        string `validate:"required" json:"password"`
	ConfirmPassword string `validate:"required" json:"confirmPassword"`
}

func Register(c *fiber.Ctx) error {

	var data RegisterRequest

	if err := c.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if err := utils.Validate.Struct(data); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	
	var userId int64 = utils.GenerateuuidBigInt()
	user := models.User{
		UserName: data.Name,
		Email:    data.Email,
		Password: hashedPassword,
		UserId:   userId,
	}

	result := database.DB.Create(&user)
	// result := database.DB.Exec("INSERT INTO users (userId, userName, email, password, updatedAt) VALUES (?, ?, ?, ?, ?)", userId, user.UserName, user.Email, user.Password, time.Now())
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }
	if result.Error != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "This Email Is Already Exist",
		})
	}

	token, err := utils.GenerateToken(userId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	cookie := fiber.Cookie{
		Name:     "accessToken",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	// var newUser models.UserModel;
	// database.DB.Select("userName, userId, role, userImage, coin, gender email").Where("userId = ?", userId).Find(&newUser);

	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{
		"message": "ลงทะเบียนสำเร็จแล้ว",
	})

}

type LoginRequest struct {
	Email				string		`validate:"required" json:"email"`;
	Password			string		`validate:"required" json:"password"`;
}

func Login(c *fiber.Ctx) error{

	var data LoginRequest;

	if err := c.BodyParser(&data); err != nil{
		return fiber.NewError(fiber.StatusInternalServerError, err.Error());
	}

	if err := utils.Validate.Struct(data); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error());
	}
	
	var user models.User;
	
	err := database.DB.Select("userId, email, password").Where("email = ?", data.Email).Find(&user).Error;

	if err != nil {
    	c.Status(fiber.StatusInternalServerError);
    	return c.Send([]byte(err.Error()));
	}

	
	// rows, err := database.DB.Raw("SELECT * FROM users WHERE email = ? LIMIT 1", data["email"]).Rows();
	// if err != nil {
	// 	c.Status(fiber.StatusInternalServerError)
	// 	return c.Send([]byte(err.Error()))
	// }
	// defer rows.Close();
	
	
	// // var users []models.User;
	

	// for rows.Next() {
	// 	// var user models.User;
    //     if err := rows.Scan(&user.UserId, &user.Username, &user.Password, &user.Email); err != nil {
    //         log.Fatal(err);
    //     }
    //     // fmt.Printf("%s, %s, %s, %s\n", user.UserID, user.Username, user.Email, user.Password);

	// 	// users = append(users, user);

    // }
	// // fmt.Println("THIS IS USER =>",user.Password);
	
	if(user.Password == nil){
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Email Is Not Exist",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data.Password)); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
			"data": user.Password,
		})
	}

	token, err := utils.GenerateToken(user.UserId);
	if (err != nil){
		return fiber.NewError(fiber.StatusInternalServerError, err.Error());
	}

	cookie := fiber.Cookie{
		Name:     "accessToken",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie);

	return c.JSON(fiber.Map{
		"message": "เข้าสู่ระบบสำเร็จแล้ว",
	})


}


func CurrentUser(c *fiber.Ctx) error{
	
	var jwtCookie string = c.Cookies("accessToken")      


	if(len(jwtCookie) > 2){
		userId, err := utils.GetUserIdFromToken(jwtCookie);
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Something Wronge With Your Token");
		}

		var userData models.User;
	
		database.DB.Select("userName, userId, role, userImage, coin, gender, email, userImage").Where("userId = ?", userId).Find(&userData);

		// database.DB.Unscoped().Select("userId ,email").Where("userId = ?", userId).Find(&userData);
		// fmt.Println(userData)
		// fmt.Printf("t1: %T\n", userData.UserId)



		// rows, err := database.DB.Raw("SELECT userId, userName, email, role, coin, gender, userImage FROM novel.users WHERE userId = ? ", userId).Rows();
		// if err != nil {
		// 	c.Status(fiber.StatusInternalServerError)
		// 	return c.Send([]byte(err.Error()))
		// }
		// defer rows.Close();
		
		
		// // var users []models.User;
		// // var users []models.User;

		// for rows.Next() {
		// 	// var user models.User;
		// 	if err := rows.Scan(&userData.UserId, &userData.UserName, &userData.Email, &userData.Role, &userData.Coin, &userData.Gender, &userData.UserImage); err != nil {
		// 		log.Fatal(err);
		// 	}

		// 	// fmt.Printf("%d, %s, %s, %s\n", user.UserId, user.UserName, user.Email, user.Password);

		// 	// users = append(users, user);

    	// }

		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"user": userData,
		});

	}else {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"user": nil,
		});
	}
};