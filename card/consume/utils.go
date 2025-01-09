package consume

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/twiglab/crm/card/orm/ent"
)

// import (
// 	"context"
// 	"fmt"
// 	"strconv"
// 	"strings"

// 	"github.com/twiglab/crm/card/orm/ent"
// )

// genCode 生成paycode
//
//	@param code
//	@param balance
//	@return string
func genCode(code string, balance int64) string {
	// TODO get version from config
	version := "1"
	pc := fmt.Sprintf("~!%s|%s|%d", version, code, balance)
	return pc
}

// // genPayCode 生成paycode
// //
// //	@param ctx
// //	@param client
// //	@param balance
// //	@return string
// //	@return error
// func genPayCode(ctx context.Context, client *ent.Client, balance int64, cardCode string) (string, error) {
// 	code, err := addRecord(ctx, client, cardCode)
// 	if err != nil {
// 		return "", err
// 	}
// 	pc := genCode(code, balance)
// 	err = updateRecordPaycode(ctx, client, cardCode, pc)
// 	return pc, err
// }

// 拆分paycode
func splitPayCode(payCode string) (string, string, int64, error) {
	if !strings.HasPrefix(payCode, "~!") {
		return "", "", 0, fmt.Errorf("invalid paycode: %s")
	}
	x := strings.TrimPrefix(payCode, "~!")
	arr := strings.Split(x, "|")
	version := arr[0]
	code := arr[1]
	balance, err := strconv.ParseInt(arr[2], 10, 64)
	if err != nil {
		return "", "", 0, err
	}
	return version, code, balance, nil
}

// calcBalance 计算余额
//
//	@param dbBalance 数据库余额
//	@param changeList 变动列表
//	@return bool
func calcBalance(dbBalance int64, changeList []*ent.ChargeRecord) int64 {

	for _, c := range changeList {
		dbBalance -= c.Deduct
	}

	return dbBalance
}
