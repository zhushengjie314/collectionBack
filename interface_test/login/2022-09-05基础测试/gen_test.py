'''
Author: 朱圣杰
Date: 2022-09-05 14:21:38
LastEditors: 朱圣杰
LastEditTime: 2022-09-05 19:40:24
FilePath: /userCenter/gen_csv.py
Description: 生成测试数据

'''
import os
import json
import pandas as pd

inFile = "test.json"
outFile = "out.json"

with open(inFile,'r',encoding='utf8') as f:
    data = json.load(f)
    result = []
    parameters = data.keys()
    for id in data["id"]:
        for pass_ in data["pass"]:
            row = {}
            row["id"] = id["value"] if "value" in id else None
            row["pass"] = pass_["value"] if "value" in pass_ else None
            row["tag"] = id['tag'] + "," + pass_["tag"]
            row["want_result"] = False
            row["want_err_code"] = 1001
            result.append(row)
    print(result)
    with open(outFile,'w',encoding='utf8') as s:
        json.dump(result,s)