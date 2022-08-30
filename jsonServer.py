#!/usr/bin/env python3
# -*- coding: utf-8 -*-

#参考：https://qiita.com/tkj/items/1338ad081038fa64cef8

import http.server as s
import json

class MyHandler(s.BaseHTTPRequestHandler):
    def do_POST(self):

        # リクエスト取得
        content_len  = int(self.headers.get("content-length"))
        body = json.loads(self.rfile.read(content_len).decode('utf-8'))
        print(body)

        # レスポンス処理
        body["test"]="response"
        self.send_response(200)
        self.send_header('Content-type', 'application/json;charset=utf-8')
        self.end_headers()
        self.wfile.write(body)

# サーバ起動
host = '0.0.0.0'
port = 3333
httpd = s.HTTPServer((host, port), MyHandler)
httpd.serve_forever()