import requests
import auth

class TestGettingActualPosts:
    def launch(self):
        print("# Test Getting Actual Posts")
        self.test_OK_actual_posts()

    def test_OK_actual_posts(self):
        try:
            resp = requests.get("http://localhost/backend/posts/actual")
            if resp.status_code != 200:
                raise Exception(f"Статус не равен 200: {resp.status_code}")
            json_ = resp.json()
            if len(json_) != 1:
                raise Exception(f'Неверная длина массива: {len(json_)}')
            print("OK test_OK_actual_posts")
        except Exception as e:
            print(f"FAIL test_OK_actual_posts: {e}")

class TestCreatePost:
    def launch(self):
        print("# Test Create Post")
        self.test_FAIL_unexisting_token()
        self.test_FAIL_post_without_title()
        self.test_FAIL_post_without_text()
        self.test_FAIL_post_with_incorrect_timeEnd()
        self.test_OK_admin_creates_post()
        self.test_OK_bloger_creates_post()
        self.test_OK_admin_creates_post_with_timeEnd()

        

    # * Without timeStart and timeEnd
    def test_OK_admin_creates_post(self):
        try:
            token = auth.Auth().getAdminToken()
            if not token:
                raise Exception("Токен не получен!")
            resp = requests.post("http://localhost/backend/posts/new", {
                "title": "test title",
                "annotation": "test annotation",
                "text": "test text",
            }, headers= {
                "auth-token": token
            })
            if resp.status_code != 201:
                raise Exception(f"Статус не равен 201: {resp.status_code}")
            print("OK test_OK_admin_creates_post")
        except Exception as e:
            print(f"FAIL test_OK_admin_creates_post: {e}")

    # * Without timeStart and timeEnd
    def test_OK_bloger_creates_post(self):
        try:
            token = auth.Auth().getBlogerToken()
            if not token:
                raise Exception("Токен не получен!")
            resp = requests.post("http://localhost/backend/posts/new", {
                "title": "test title",
                "annotation": "test annotation",
                "text": "test text",
            }, headers= {
                "auth-token": token
            })
            if resp.status_code != 201:
                raise Exception(f"Статус не равен 201: {resp.status_code}")
            print("OK test_OK_bloger_creates_post")
        except Exception as e:
            print(f"FAIL test_OK_bloger_creates_post: {e}")

    # * With timeEnd and without timeStart
    def test_OK_admin_creates_post_with_timeEnd(self):
        try:
            token = auth.Auth().getAdminToken()
            if not token:
                raise Exception("Токен не получен!")
            resp = requests.post("http://localhost/backend/posts/new", {
                "title": "test title",
                "annotation": "test annotation",
                "text": "test text",
                "timeEnd": "2023-01-15 01:35:00"
            }, headers= {
                "auth-token": token
            })
            if resp.status_code != 201:
                raise Exception(f"Статус не равен 201: {resp.status_code} - {resp.json()} ")
            print("OK test_OK_admin_creates_post_with_timeEnd")
        except Exception as e:
            print(f"FAIL test_OK_admin_creates_post_with_timeEnd: {e}")

    def test_FAIL_unexisting_token(self):
        try:
            token = "wrong_unexisting_token!"
            resp = requests.post("http://localhost/backend/posts/new", {
                "title": "test title",
                "annotation": "test annotation",
                "text": "test text",
            }, headers= {
                "auth-token": token
            })
            if resp.status_code != 404:
                raise Exception(f"Статус не равен 404: {resp.status_code}")
            print("OK test_FAIL_unexisting_token")
        except Exception as e:
            print(f"FAIL test_FAIL_unexisting_token: {e}")

    def test_FAIL_post_without_title(self):
        try:
            token = auth.Auth().getAdminToken()
            if not token:
                raise Exception('Токен не получен!')
            resp = requests.post("http://localhost/backend/posts/new", {
                #! "title": "test title",
                "annotation": "test annotation",
                "text": "test text",
            }, headers= {
                "auth-token": token
            })
            if resp.status_code != 400:
                raise Exception(f"Статус не равен 400: {resp.status_code}")
            print("OK test_FAIL_post_without_title")
        except Exception as e:
            print(f"FAIL test_FAIL_post_without_title: {e}")

    def test_FAIL_post_without_text(self):
        try:
            token = auth.Auth().getAdminToken()
            if not token:
                raise Exception('Токен не получен!')
            resp = requests.post("http://localhost/backend/posts/new", {
                "title": "test title",
                "annotation": "test annotation",
                #! "text": "test text",
            }, headers= {
                "auth-token": token
            })
            if resp.status_code != 400:
                raise Exception(f"Статус не равен 400: {resp.status_code}")
            print("OK test_FAIL_post_without_text")
        except Exception as e:
            print(f"FAIL test_FAIL_post_without_text: {e}")

    def test_FAIL_post_without_annotation(self):
        try:
            token = auth.Auth().getAdminToken()
            if not token:
                raise Exception('Токен не получен!')
            resp = requests.post("http://localhost/backend/posts/new", {
                "title": "test title",
                #! "annotation": "test annotation",
                "text": "test text",
            }, headers= {
                "auth-token": token
            })
            if resp.status_code != 400:
                raise Exception(f"Статус не равен 400: {resp.status_code}")
            print("OK test_FAIL_post_without_text")
        except Exception as e:
            print(f"FAIL test_FAIL_post_without_text: {e}")

    def test_FAIL_post_with_incorrect_timeEnd(self):
        try:
            token = auth.Auth().getAdminToken()
            if not token:
                raise Exception("Токен не получен!")
            resp = requests.post("http://localhost/backend/posts/new", {
                "title": "test title",
                "annotation": "test annotation",
                "text": "test text",
                "timeEnd": "2023-01-03T01:35:00"
            }, headers= {
                "auth-token": token
            })
            if resp.status_code != 400:
                raise Exception(f"Статус не равен 400: {resp.status_code}")
            print("OK test_FAIL_post_with_incorrect_timeEnd")
        except Exception as e:
            print(f"FAIL test_FAIL_post_with_incorrect_timeEnd: {e}")


def launch():
    print("\n--- Posts Controller ---\n")
    TestGettingActualPosts().launch()
    TestCreatePost().launch()
