package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mengjay315/lottery/db/dbtypes"
	"github.com/mengjay315/lottery/db/zcpg"
	db "github.com/mengjay315/lottery/db/zcpg/database"
	"github.com/mengjay315/lottery/model"
	"log"
	"net/http"
	"strconv"
	"time"
)

func SignIn(c *gin.Context) {
	name := c.Query("name")
	depart := c.Query("depart")
	memo := c.Query("memo")

	inTime := time.Now().Unix()

	// Able=0代表该人没有投过票，后续只要投过一次票后，该值变为1
	person := &dbtypes.PersonBasic{
		Name: name,
		Memo: memo,
		Time: inTime,
		Able: 0,
	}

	program := &dbtypes.ProgramBasic{
		Depart: depart,
		Number: 0,
	}

	// 先判断是否签过，看persons 表中是否有这个人，
	// 有的话，直接返回错误信息，已签到

	personRes, err := model.GetPerson(name)
	if err != nil {
		log.Fatal(err)
	}

	if personRes.Name == "" {
		_, err := zcpg.InsertPerson(db.DB, person)
		if err != nil {
			log.Fatalf("insert person error %v", err)
		}

		//先查询有没有这个部门，有的话就不用存了
		queryProg, err := model.GetProgram(depart)
		if err != nil {
			log.Fatal(err)
		}

		if queryProg.Depart == "" {
			_, err = zcpg.InsertProgram(db.DB, program)
			if err != nil {
				log.Fatalf("insert program error %v", err)
			}

			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"memo":    memo,
				"message": "Sign_in successfully!",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"memo":    memo,
				"message": "Sign_in successfully!",
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"memo":    memo,
			"message": "You have already signIn, return your memo",
		})
	}
}

func SetVoteID(c *gin.Context) {
	id := c.Query("vid") // vid =1, id=1

	vid, _ := strconv.Atoi(id)

	voteID := dbtypes.VoteId{Vid: vid}

	// 更新vote ID
	//db := database.InitDB()
	_, err := zcpg.UpdateVoteID(db.DB, &voteID)
	if err != nil {
		log.Fatalf("update vidID error %v", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Could vote in time",
	})
}

func Vote(c *gin.Context) {
	// 先判断是否有投票的信号
	// 获取投票信号，为0，不可投，为1可投

	vidRes, err := model.GetVid()
	if err != nil {
		log.Fatal(err)
	}

	if vidRes.Vid == 1 {
		// 先判断该名字有没有投过
		// 一个名字只能投一次
		name := c.Query("name")
		depart := c.Query("depart")

		// 查询person表的able字段为0或1
		person, err := model.GetPerson(name)
		if err != nil {
			log.Fatal(err)
		}

		if person.Able == 0 {
			// 投票操作
			program, err := model.GetProgram(depart)
			if err != nil {
				log.Fatal(err)
			}

			program.Number = program.Number + 1

			// 更新depart number
			//db := database.InitDB()
			_, err = zcpg.UpdateProgram(db.DB, &program)
			if err != nil {
				log.Fatalf("update program error %v", err)
			}

			// 更新person able = 1
			person.Able = 1

			_, err = zcpg.UpdatePerson(db.DB, &person)
			if err != nil {
				log.Fatalf("update person error %v", err)
			}

			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "Vote successfully!",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "Sorry, everyone can only vote once",
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "Sorry, It's not time to vote",
		})
	}
}

func GetVoteRes(c *gin.Context) {
	//c.Header("Access-Control-Allow-Origin", "*")
	voteRes, err := model.GetVoteNums()
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"voteRes": voteRes,
		"message": "Query vote res successfully",
	})
}

func GetPersons(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	persons, err := model.GetAllPerson()
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":       200,
		"allPersons": persons,
		"message":    "Query persons res successfully",
	})
}

//-----------------------------------------
//// websocket实现
//var upGrader = websocket.Upgrader{
//	CheckOrigin: func(r *http.Request) bool {
//		return true
//	},
//}

func GetSignInRes(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	res, err := model.GetSignRes()
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":      200,
		"signInRes": res,
		"message":   "Query signIn res successfully",
	})

	//ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	//if err != nil {
	//	return
	//}

	//defer ws.Close()

	//for {
	//	// 读取ws中的数据
	//	mt, message, err := ws.ReadMessage()
	//	if err != nil {
	//		// 客户端关闭连接时也会进入
	//		log.Fatal(err)
	//		break
	//	}
	//
	//	res, err := model.GetSignRes()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	if string(message) == "signin_res" {
	//		message, _ = json.Marshal(res)
	//	}
	//
	//	// 写入ws数据，
	//	err = ws.WriteMessage(mt, message)
	//	// 返回json字符串
	//	signInRes:= gin.H{"res": res}
	//	err = ws.WriteJSON(signInRes)
	//	if err != nil {
	//		log.Fatal(err)
	//		break
	//	}
	//}

	//res, err := model.GetSignRes()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//
	//
	//
	//
	//c.JSON(http.StatusOK, gin.H{
	//	"code": 200,
	//	"signInRes": res,
	//	"message": "Query signIn res successfully",
	//})
}
