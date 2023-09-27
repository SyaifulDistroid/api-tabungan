package helper

import (
	"api-tabungan/domain/tabungan/model"
	"fmt"
	"strconv"
	"time"
)

func GenerateNoRekening(request model.CreateRekeningRequest) (noRekening int64) {
	var time = time.Now()
	noRek := fmt.Sprintf("%v%v%v%v%v", time.Minute(), len(request.Nama), time.Day(), time.Second(), time.Hour())
	intNoRek, _ := strconv.Atoi(noRek)
	noRekening = int64(intNoRek)
	return
}

func StrToInt64(s string) (i int64) {
	str, _ := strconv.Atoi(s)
	return int64(str)
}
