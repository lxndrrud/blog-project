import requests
from typing import Optional

class Auth:
    def getAdminToken(self) -> Optional[str]:
        try:
            resp = requests.post("http://localhost/backend/users/login", json={
                "login": "admin",
                "password": "123456"
            })
            json_ = resp.json()
            if resp.status_code != 200:
                raise Exception(f"Статус не равен 200: {resp.status_code}")
            if not json_['token']:
                raise Exception("В ответе нет токена!")
            return json_['token']
        except Exception as e:
            print(f"FAIL getAdminToken: {e}")
            return ""

    def getBlogerToken(self) -> Optional[str]:
        try:
            resp = requests.post("http://localhost/backend/users/login", json={
                "login": "bloger",
                "password": "123456"
            })
            json_ = resp.json()
            if resp.status_code != 200:
                raise Exception(f"Статус не равен 200: {resp.status_code}")
            if not json_['token']:
                raise Exception("В ответе нет токена!")
            return json_['token']
        except Exception as e:
            print(f"FAIL getAdminToken: {e}")
            return ""