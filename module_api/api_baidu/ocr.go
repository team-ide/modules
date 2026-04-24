// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:module_api/baidu/ocr.coos] 在 [2026-04-24 17:44] 生成

package api_baidu

type OcrAccurateBasic struct {
	LogId          int64                   `json:"log_id"`
	Direction      int32                   `json:"direction"`
	WordsResultNum int32                   `json:"words_result_num"`
	WordsResult    []*OcrAccurateBasicWord `json:"words_result"`
}

type OcrAccurateBasicWord struct {
	Words string `json:"words"`
}

func OcrAccurateBasicByImage64(image64 string) (res *OcrAccurateBasic, err error) {
	return
}
