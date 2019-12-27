package others

import (
	"encoding/json"
	"strings"
)

/*
* 题目: 将所给的路径字符串数组转化为 json 字符串
*
* 背景: 从 oss 取出的文件列表格式为字符串数组, 但是前端需要转换成 map(对于前端而言即json),
* 固有此需求.
*
* 示例
* files:
*		public/test.md
*		public/readme.md
*		readme.md
*		test/
*
* result:
*		{
*			public:{
*				readme.md:readme.md,
*				test.md:test.md
*			},
*			readme.md:readme.md,
*			test:{}
*		}
 */

// 解决思想: 分两个方面, 第一是递归处理单条路径, 第二是处理多条路径

func ls2Json(paths []string) string {
	var result = make(map[string]interface{})
	for _, path := range paths {
		parsePath(path, "/", result)
	}
	buf, _ := json.Marshal(result)
	return string(buf)
}

func parsePath(path string, delimiter string, result map[string]interface{}) map[string]interface{} {
	if path == "" {
		return nil
	}
	dirs := strings.Split(path, delimiter)
	parseFile(dirs, result)
	return result
}

func parseFile(dirs []string, result map[string]interface{}) map[string]interface{} {
	var key = dirs[0]
	switch len(dirs) {
	case 1:
		if key == "" {
			return nil
		}
		result[key] = key
	case 0:
		return nil
	default:
		// 检测是否添加过该key,且当key为文件时复写为map
		var temp map[string]interface{}
		if v, ok := result[key]; ok {
			var isMap bool
			temp, isMap = v.(map[string]interface{})
			if !isMap {
				temp = make(map[string]interface{})
				result[key] = temp
			}
		} else {
			temp = make(map[string]interface{})
			result[key] = temp
		}
		parseFile(dirs[1:], temp)
	}
	return result
}
