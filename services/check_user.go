package services

import (
	"ADMgmtSystem/library"
	"ADMgmtSystem/models"
	"ADMgmtSystem/models/db"
	"crypto/sha256"
	"fmt"
)

func CheckUser(userAccount, userPassword string) error {
	library.Log.Info("Start function is 'CheckUser'")

	var users models.Users

	var user db.UsersRepository = &db.User{}

	library.Log.Trace("Search database by user account")
	users, err := user.GetUserByAccount(userAccount)

	if err != nil {
		library.Log.Errorf("%v", err)

		return err
	}

	library.Log.Trace("compare database and user account")
	if users.Account != userAccount {
		library.Log.Errorf("No account")

		return fmt.Errorf("%s", "No account!")
	}

	library.Log.Trace("Check user error count")
	if users.ErrorCount >= 5 {
		library.Log.Errorf("More then 5 errors")

		return fmt.Errorf("%s", "More then 5 errors!")
	}

	library.Log.Trace("compare database and user password")
	userPasswordSha256 := fmt.Sprintf("%x", sha256.Sum256([]byte(userPassword+users.PasswordSalt)))
	if userPasswordSha256 != users.Password {
		library.Log.Errorf("No match password")

		err := user.UpdateUserErrorById(users.ID)
		if err != nil {
			library.Log.Errorf("%v", err)
			return err
		}
		return fmt.Errorf("%s", "No match password!")
	}

	library.Log.Info("Finish function is 'CheckUser'")
	return nil
}
