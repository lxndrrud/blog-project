import requests


class TestUserProfile:
    def launch(self):
        print("# Test User profile")
        self.test_OK_get_user_profile()
        self.test_NOT_FOUND_get_user_profile()

    def test_OK_get_user_profile(self):
        try:
            resp = requests.get("http://localhost/backend/users/profile?idUser=1")
            if resp.status_code != 200:
                raise Exception(f"Статус не равен 200: {resp.status_code}")
            json_ = resp.json()
            if not json_["user"] or not json_["posts"]:
                raise Exception(f'Неверная структура: {json_}')
            print(f"OK test_OK_get_user_profile")
        except Exception as e:
            print(f"FAIL test_OK_get_user_profile: {e}")

    def test_NOT_FOUND_get_user_profile(self):
        try:
            resp = requests.get("http://localhost/backend/users/profile?idUser=2")
            if resp.status_code != 404:
                raise Exception(f"Статус не равен 404: {resp.status_code}")
            
            print(f"OK test_NOT_FOUND_get_user_profile")
        except Exception as e:
            print(f"FAIL test_NOT_FOUND_get_user_profile: {e}")

class TestUserLogin:
    def launch(self):
        print("# Test User Login")
        self.test_OK_login_user()
        self.test_NOT_FOUND_login_user_wrong_login()
        self.test_NOT_FOUND_login_user_wrong_password()
        
    def test_OK_login_user(self):
        try:
            resp = requests.post("http://localhost/backend/users/login", {
                "login": "admin",
                "password": "123456"
            })
            if resp.status_code != 200:
                raise Exception(f"Статус не равен 200: {resp.status_code}")
            json_ = resp.json()
            if not "token" in json_:
                raise Exception(f"В теле ответа нет токена!")
            print(f"OK test_OK_login_user")
        except Exception as e:
            print(f"FAIL test_OK_login_user: {e}")

    def test_NOT_FOUND_login_user_wrong_login(self):
        try:
            resp = requests.post("http://localhost/backend/users/login", {
                "login": "wrong!",
                "password": "123456"
            })
            if resp.status_code != 404:
                raise Exception(f"Статус не равен 404: {resp.status_code}")
            json_ = resp.json()
            if not "message" in json_:
                raise Exception(f"В теле ответа нет сообщения об ошибке!")
            print(f"OK test_NOT_FOUND_login_user_wrong_login")
        except Exception as e:
            print(f"FAIL test_NOT_FOUND_login_user_wrong_login: {e}")

    def test_NOT_FOUND_login_user_wrong_password(self):
        try:
            resp = requests.post("http://localhost/backend/users/login", {
                "login": "admin",
                "password": "123456!"
            })
            if resp.status_code != 404:
                raise Exception(f"Статус не равен 404: {resp.status_code}")
            json_ = resp.json()
            if not "message" in json_:
                raise Exception(f"В теле ответа нет сообщения об ошибке!")
            print(f"OK test_NOT_FOUND_login_user_wrong_password")
        except Exception as e:
            print(f"FAIL test_NOT_FOUND_login_user_wrong_password: {e}")
        


def launch():
    print("\n--- User Controller ---\n")
    TestUserProfile().launch()
    TestUserLogin().launch()