package transaction

import (
	"github.com/gofiber/fiber/v2"
	"github.com/levigross/grequests"
	"github.com/tidwall/gjson"
	"go_be_dev/auth"
	"time"
)

func PostClockIn(c *fiber.Ctx) error {
	authorization := c.Get("Authorization")

	tokenString := authorization[7:]
	claims, err := auth.GetUserInfoFromToken(tokenString)
	if err != nil {
		// Handle error (e.g., token invalid or expired)
		return c.Status(500).JSON(fiber.Map{"status": 0, "message": fiber.StatusInternalServerError})
	}

	username := claims["username"].(string)

	getUserInfoByUsername := GetUserInfoByUsername(username)

	if len(getUserInfoByUsername) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": 0, "message": fiber.StatusNotFound})
	} else {
		userIP := c.IP() // Get user's IP address from the request

		// Call the IP geolocation service
		url := "http://ip-api.com/json/" + userIP
		resp, err := grequests.Get(url, nil)
		if err != nil {
			return err
		}

		// Parse the JSON response
		var result gjson.Result
		if err := resp.JSON(&result); err != nil {
			return err
		}
		location, err := time.LoadLocation("Asia/Jakarta")
		if err != nil {
			return err
		}

		timeNow := time.Now().In(location).Format("2006-01-02 15:04:05")

		latitude := result.Get("lat").Float()
		longitude := result.Get("lon").Float()
		userId := getUserInfoByUsername[0].Id

		data := DataTransaction{
			UserId:    userId,
			IpAddress: userIP,
			Latitude:  latitude,
			Longitude: longitude,
			ClockIn:   timeNow,
		}

		transactionClockIn := TransactionClockIn(data)
		if transactionClockIn != nil {
			return c.Status(400).JSON(fiber.Map{
				"status":  false,
				"message": transactionClockIn.Error(),
			})
		} else {
			return c.Status(200).JSON(fiber.Map{
				"status":  true,
				"message": "Clock in succes",
			})
		}
	}
	return nil
}

func PostClockOut(c *fiber.Ctx) error {
	authorization := c.Get("Authorization")

	tokenString := authorization[7:]
	claims, err := auth.GetUserInfoFromToken(tokenString)
	if err != nil {
		// Handle error (e.g., token invalid or expired)
		return c.Status(500).JSON(fiber.Map{"status": 0, "message": fiber.StatusInternalServerError})
	}

	username := claims["username"].(string)

	getUserInfoByUsername := GetUserInfoByUsername(username)

	if len(getUserInfoByUsername) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": 0, "message": fiber.StatusNotFound})
	} else {
		userIP := c.IP() // Get user's IP address from the request

		// Call the IP geolocation service
		url := "http://ip-api.com/json/" + userIP
		resp, err := grequests.Get(url, nil)
		if err != nil {
			return err
		}

		// Parse the JSON response
		var result gjson.Result
		if err := resp.JSON(&result); err != nil {
			return err
		}
		location, err := time.LoadLocation("Asia/Jakarta")
		if err != nil {
			return err
		}

		timeNow := time.Now().In(location).Format("2006-01-02 15:04:05")

		latitude := result.Get("lat").Float()
		longitude := result.Get("lon").Float()
		userId := getUserInfoByUsername[0].Id

		data := DataTransaction2{
			UserId:    userId,
			IpAddress: userIP,
			Latitude:  latitude,
			Longitude: longitude,
			ClockOut:  timeNow,
		}

		transactionClockIn := TransactionClockOut(data)
		if transactionClockIn != nil {
			return c.Status(400).JSON(fiber.Map{
				"status":  false,
				"message": transactionClockIn.Error(),
			})
		} else {
			return c.Status(200).JSON(fiber.Map{
				"status":  true,
				"message": "Clock out succes",
			})
		}
	}
	return nil
}
