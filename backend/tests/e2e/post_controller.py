import requests
import auth

class TestGetUserPosts:
    def launch(self):
        print('# Test Get User Posts')
        self.test_FORBIDDEN_getting_user_posts_without_token()
        self.test_OK_admin_gets_own_posts()
        self.test_OK_bloger_gets_own_posts()

    def test_OK_admin_gets_own_posts(self):
        try:
            token = auth.Auth().getAdminToken()
            if not token:
                raise Exception('Токен не получен!')
            resp = requests.get("http://localhost/backend/posts/userPosts", headers={
                'auth-token': token,
            })
            if resp.status_code != 200:
                raise Exception(f"Статус не равен 200: {resp.status_code} - {resp.json()}")
            json_ = resp.json()
            if len(json_) == 0:
                raise Exception(f'Неверная длина массива: {len(json_)}')
            print("OK test_OK_admin_gets_own_posts")
        except Exception as e:
            print(f"FAIL test_OK_admin_gets_own_posts: {e}")

    def test_OK_bloger_gets_own_posts(self):
        try:
            token = auth.Auth().getBlogerToken()
            if not token:
                raise Exception('Токен не получен!')
            resp = requests.get("http://localhost/backend/posts/userPosts", headers={
                'auth-token': token,
            })
            if resp.status_code != 200:
                raise Exception(f"Статус не равен 200: {resp.status_code} - {resp.json()}")
            print("OK test_OK_bloger_gets_own_posts")
        except Exception as e:
            print(f"FAIL test_OK_bloger_gets_own_posts: {e}")

    def test_FORBIDDEN_getting_user_posts_without_token(self):
        try:
            resp = requests.get("http://localhost/backend/posts/userPosts")
            if resp.status_code != 403:
                raise Exception(f"Статус не равен 403: {resp.status_code} - {resp.json()}")
            json_ = resp.json()
            if "message" not in json_:
                raise Exception("В теле запроса нет сообщения об ошибке!")
            print("OK test_FORBIDDEN_getting_user_posts_without_token")
        except Exception as e:
            print(f"FAIL test_FORBIDDEN_getting_user_posts_without_token: {e}")

class TestGetActualPosts:
    def launch(self):
        print("# Test Get Actual Posts")
        self.test_OK_get_actual_posts()

    def test_OK_get_actual_posts(self):
        try:
            resp = requests.get("http://localhost/backend/posts/actual")
            if resp.status_code != 200:
                raise Exception(f"Статус не равен 200: {resp.status_code} - {resp.json()}")
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
            resp = requests.post("http://localhost/backend/posts/new", json={
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
            resp = requests.post("http://localhost/backend/posts/new", json={
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
            resp = requests.post("http://localhost/backend/posts/new", json={
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
            resp = requests.post("http://localhost/backend/posts/new", json={
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
            resp = requests.post("http://localhost/backend/posts/new", json={
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
            resp = requests.post("http://localhost/backend/posts/new", json={
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
            resp = requests.post("http://localhost/backend/posts/new", json={
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
            resp = requests.post("http://localhost/backend/posts/new", json={
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

class TestGetActualPost:
    def launch(self):
        print("# Test Get Actual Post")
        self.test_NOT_FOUND_get_actual_post_not_approved()
        self.test_OK_get_actual_post()

    def test_OK_get_actual_post(self):
        try:
            resp = requests.get("http://localhost/backend/posts/get", {
                "idPost": 2,
            })
            if resp.status_code != 200:
                raise Exception(f"Статус не равен 200: {resp.status_code} - {resp.json()}")
            json_ = resp.json()
            if not json_["post"]:
                raise Exception("В теле ответа нет поста!")
            print("OK test_OK_get_actual_post")
        except Exception as e:
            print(f"FAIL test_OK_get_actual_post: {e}")

    def test_NOT_FOUND_get_actual_post_not_approved(self):
        try:
            resp = requests.get("http://localhost/backend/posts/get", {
                "idPost": 1,
            })
            if resp.status_code != 404:
                raise Exception(f"Статус не равен 404: {resp.status_code} - {resp.json()}")
            json_ = resp.json()
            if not json_["message"]:
                raise Exception("В теле ответа нет сообщения об ошибке!")
            print("OK test_NOT_FOUND_get_actual_post_not_approved")
        except Exception as e:
            print(f"FAIL test_NOT_FOUND_get_actual_post_not_approved: {e}")

class TestGetPostsNeedToApprove:
    def launch(self):
        print("# Test Get Posts Need To Approve")
        self.test_FORBIDDEN_bloger_get_posts_need_to_approve()
        self.test_NOT_FOUND_get_posts_need_to_approve_unknown_token()
        self.test_FORBIDDEN_get_posts_need_to_approve_without_token()
        self.test_OK_admin_get_posts_need_to_approve()

    def test_OK_admin_get_posts_need_to_approve(self):
        try:
            token = auth.Auth().getAdminToken()
            if not token:
                raise Exception("Токен не получен!")
            resp = requests.get("http://localhost/backend/posts/needToApprove", headers={
                "auth-token": token
            })
            if resp.status_code != 200:
                raise Exception(f"Статус не равен 200: {resp.status_code} - {resp.json()}")
            json_ = resp.json()
            if not json_["posts"]:
                raise Exception(f"В теле ответа нет постов: {json_}")
            print("OK test_OK_admin_get_posts_need_to_approve")
        except Exception as e:
            print(f"FAIL test_OK_admin_get_posts_need_to_approve: {e}")

    def test_FORBIDDEN_bloger_get_posts_need_to_approve(self):
        try:
            token = auth.Auth().getBlogerToken()
            if not token:
                raise Exception("Токен не получен!")
            resp = requests.get("http://localhost/backend/posts/needToApprove", headers={
                "auth-token": token
            })
            if resp.status_code != 403:
                raise Exception(f"Статус не равен 403: {resp.status_code} - {resp.json()}")
            json_ = resp.json()
            if not json_["message"]:
                raise Exception(f"В теле ответа нет сообщения об ошибке: {json_}")
            print("OK test_FORBIDDEN_bloger_get_posts_need_to_approve")
        except Exception as e:
            print(f"FAIL test_FORBIDDEN_bloger_get_posts_need_to_approve: {e}")

    def test_NOT_FOUND_get_posts_need_to_approve_unknown_token(self):
        try:
            token = "unknown_token"
            if not token:
                raise Exception("Токен не получен!")
            resp = requests.get("http://localhost/backend/posts/needToApprove", headers={
                "auth-token": token
            })
            if resp.status_code != 404:
                raise Exception(f"Статус не равен 404: {resp.status_code} - {resp.json()}")
            json_ = resp.json()
            if not json_["message"]:
                raise Exception(f"В теле ответа нет сообщения об ошибке: {json_}")
            print("OK test_NOT_FOUND_get_posts_need_to_approve_unknown_token")
        except Exception as e:
            print(f"FAIL test_NOT_FOUND_get_posts_need_to_approve_unknown_token: {e}")

    def test_FORBIDDEN_get_posts_need_to_approve_without_token(self):
        try:
            resp = requests.get("http://localhost/backend/posts/needToApprove")
            if resp.status_code != 403:
                raise Exception(f"Статус не равен 403: {resp.status_code} - {resp.json()}")
            json_ = resp.json()
            if not json_["message"]:
                raise Exception(f"В теле ответа нет сообщения об ошибке: {json_}")
            print("OK test_FORBIDDEN_get_posts_need_to_approve_without_token")
        except Exception as e:
            print(f"FAIL test_FORBIDDEN_get_posts_need_to_approve_without_token: {e}")

class TestGetPostNeedToApprove:
    def launch(self):
        print("# Test Get Post Need To Approve")
        self.test_FORBIDDEN_bloger_get_post_need_to_approve()
        self.test_NOT_FOUND_get_post_need_to_approve_unknown_token()
        self.test_FORBIDDEN_get_post_need_to_approve_without_token()
        self.test_NOT_FOUND_get_unexisting_post_need_to_approve()
        self.test_OK_admin_get_post_need_to_approve()

    def test_OK_admin_get_post_need_to_approve(self):
        try:
            token = auth.Auth().getAdminToken()
            if not token:
                raise Exception("Токен не получен!")
            resp = requests.get("http://localhost/backend/posts/needToApprove/get?idPost=1", headers={
                "auth-token": token
            })
            if resp.status_code != 200:
                raise Exception(f"Статус не равен 200: {resp.status_code} - {resp.json()}")
            json_ = resp.json()
            if not json_["post"]:
                raise Exception(f"В теле ответа нет постов: {json_}")
            print("OK test_OK_admin_get_post_need_to_approve")
        except Exception as e:
            print(f"FAIL test_OK_admin_get_post_need_to_approve: {e}")

    def test_FORBIDDEN_bloger_get_post_need_to_approve(self):
        try:
            token = auth.Auth().getBlogerToken()
            if not token:
                raise Exception("Токен не получен!")
            resp = requests.get("http://localhost/backend/posts/needToApprove/get?idPost=1", headers={
                "auth-token": token
            })
            if resp.status_code != 403:
                raise Exception(f"Статус не равен 403: {resp.status_code} - {resp.json()}")
            json_ = resp.json()
            if not json_["message"]:
                raise Exception(f"В теле ответа нет сообщения об ошибке: {json_}")
            print("OK test_FORBIDDEN_bloger_get_post_need_to_approve")
        except Exception as e:
            print(f"FAIL test_FORBIDDEN_bloger_get_post_need_to_approve: {e}")

    def test_NOT_FOUND_get_post_need_to_approve_unknown_token(self):
        try:
            token = "unknown_token"
            if not token:
                raise Exception("Токен не получен!")
            resp = requests.get("http://localhost/backend/posts/needToApprove/get?idPost=1", headers={
                "auth-token": token
            })
            if resp.status_code != 404:
                raise Exception(f"Статус не равен 404: {resp.status_code} - {resp.json()}")
            json_ = resp.json()
            if not json_["message"]:
                raise Exception(f"В теле ответа нет сообщения об ошибке: {json_}")
            print("OK test_NOT_FOUND_get_post_need_to_approve_unknown_token")
        except Exception as e:
            print(f"FAIL test_NOT_FOUND_get_post_need_to_approve_unknown_token: {e}")

    def test_FORBIDDEN_get_post_need_to_approve_without_token(self):
        try:
            resp = requests.get("http://localhost/backend/posts/needToApprove/get?idPost=1")
            if resp.status_code != 403:
                raise Exception(f"Статус не равен 403: {resp.status_code} - {resp.json()}")
            json_ = resp.json()
            if not json_["message"]:
                raise Exception(f"В теле ответа нет сообщения об ошибке: {json_}")
            print("OK test_FORBIDDEN_get_post_need_to_approve_without_token")
        except Exception as e:
            print(f"FAIL test_FORBIDDEN_get_post_need_to_approve_without_token: {e}")

    def test_NOT_FOUND_get_unexisting_post_need_to_approve(self):
        try:
            token = auth.Auth().getAdminToken()
            if not token:
                raise Exception("Токен не получен!")
            resp = requests.get("http://localhost/backend/posts/needToApprove/get?idPost=2213", headers={
                "auth-token": token,
            })
            if resp.status_code != 404:
                raise Exception(f"Статус не равен 404: {resp.status_code} - {resp.json()}")
            json_ = resp.json()
            if not json_["message"]:
                raise Exception(f"В теле ответа нет сообщения об ошибке: {json_}")
            print("OK test_NOT_FOUND_get_unexisting_post_need_to_approve")
        except Exception as e:
            print(f"FAIL test_NOT_FOUND_get_unexisting_post_need_to_approve: {e}")

class TestApprovePost:
    def launch(self):
        print("# Test Approve Post")
        self.test_FORBIDDEN_bloger_approves_post()
        self.test_FORBIDDEN_approve_post_without_token()
        self.test_NOT_FOUND_approve_post_with_unknown_token()
        self.test_OK_admin_approves_post()

    def test_OK_admin_approves_post(self):
        try:
            token = auth.Auth().getAdminToken()
            if not token:
                raise Exception("Токен не получен!")
            resp = requests.post("http://localhost/backend/posts/approve", json={
                "idPost": 1,
            }, headers={
                "auth-token": token,
            })
            if resp.status_code != 200:
                raise Exception(f"Статус не равен 200: {resp.status_code} - {resp.json()}")
            print("OK test_OK_admin_approves_post")
        except Exception as e:
            print(f"FAIL test_OK_admin_approves_post: {e}")

    def test_FORBIDDEN_bloger_approves_post(self):
        try:
            token = auth.Auth().getBlogerToken()
            if not token:
                raise Exception("Токен не получен!")
            resp = requests.post("http://localhost/backend/posts/approve", json={
                "idPost": 1,
            }, headers={
                "auth-token": token,
            })
            if resp.status_code != 403:
                raise Exception(f"Статус не равен 403: {resp.status_code} - {resp.json()}")
            json_ = resp.json()
            if not json_["message"]:
                raise Exception(f"В теле ответа нет сообщения об ошибке: {json_}")
            print("OK test_FORBIDDEN_bloger_approves_post")
        except Exception as e:
            print(f"FAIL test_FORBIDDEN_bloger_approves_post: {e}")

    def test_FORBIDDEN_approve_post_without_token(self):
        try:
            resp = requests.post("http://localhost/backend/posts/approve", json={
                "idPost": 1,
            })
            if resp.status_code != 403:
                raise Exception(f"Статус не равен 403: {resp.status_code} - {resp.json()}")
            json_ = resp.json()
            if not json_["message"]:
                raise Exception(f"В теле ответа нет сообщения об ошибке: {json_}")
            print("OK test_FORBIDDEN_approve_post_without_token")
        except Exception as e:
            print(f"FAIL test_FORBIDDEN_approve_post_without_token: {e}")

    def test_NOT_FOUND_approve_post_with_unknown_token(self):
        try:
            token = "unknown_token"
            resp = requests.post("http://localhost/backend/posts/approve", json={
                "idPost": 1,
            }, headers={
                "auth-token": token,
            })
            if resp.status_code != 404:
                raise Exception(f"Статус не равен 404: {resp.status_code} - {resp.json()}")
            json_ = resp.json()
            if not json_["message"]:
                raise Exception(f"В теле ответа нет сообщения об ошибке: {json_}")
            print("OK test_NOT_FOUND_approve_post_with_unknown_token")
        except Exception as e:
            print(f"FAIL test_NOT_FOUND_approve_post_with_unknown_token: {e}")

class TestRejectPost:
    def launch(self):
        print("# Test Reject Post")
        self.test_FORBIDDEN_bloger_rejects_post()
        self.test_FORBIDDEN_reject_post_without_token()
        self.test_NOT_FOUND_reject_post_with_unknown_token()
        self.test_OK_admin_rejects_post()

    def test_OK_admin_rejects_post(self):
        try:
            token = auth.Auth().getAdminToken()
            if not token:
                raise Exception("Токен не получен!")
            resp = requests.post("http://localhost/backend/posts/reject", json={
                "idPost": 1,
            }, headers={
                "auth-token": token,
            })
            if resp.status_code != 200:
                raise Exception(f"Статус не равен 200: {resp.status_code} - {resp.json()}")
            print("OK test_OK_admin_rejects_post")
        except Exception as e:
            print(f"FAIL test_OK_admin_rejects_post: {e}")

    def test_FORBIDDEN_bloger_rejects_post(self):
        try:
            token = auth.Auth().getBlogerToken()
            if not token:
                raise Exception("Токен не получен!")
            resp = requests.post("http://localhost/backend/posts/reject", json={
                "idPost": 1,
            }, headers={
                "auth-token": token,
            })
            if resp.status_code != 403:
                raise Exception(f"Статус не равен 403: {resp.status_code} - {resp.json()}")
            json_ = resp.json()
            if not json_["message"]:
                raise Exception(f"В теле ответа нет сообщения об ошибке: {json_}")
            print("OK test_FORBIDDEN_bloger_rejects_post")
        except Exception as e:
            print(f"FAIL test_FORBIDDEN_bloger_rejects_post: {e}")

    def test_FORBIDDEN_reject_post_without_token(self):
        try:
            resp = requests.post("http://localhost/backend/posts/reject", json={
                "idPost": 1,
            })
            if resp.status_code != 403:
                raise Exception(f"Статус не равен 403: {resp.status_code} - {resp.json()}")
            json_ = resp.json()
            if not json_["message"]:
                raise Exception(f"В теле ответа нет сообщения об ошибке: {json_}")
            print("OK test_FORBIDDEN_reject_post_without_token")
        except Exception as e:
            print(f"FAIL test_FORBIDDEN_reject_post_without_token: {e}")

    def test_NOT_FOUND_reject_post_with_unknown_token(self):
        try:
            token = "unknown_token"
            resp = requests.post("http://localhost/backend/posts/reject", json={
                "idPost": 1,
            }, headers={
                "auth-token": token,
            })
            if resp.status_code != 404:
                raise Exception(f"Статус не равен 404: {resp.status_code} - {resp.json()}")
            json_ = resp.json()
            if not json_["message"]:
                raise Exception(f"В теле ответа нет сообщения об ошибке: {json_}")
            print("OK test_NOT_FOUND_reject_post_with_unknown_token")
        except Exception as e:
            print(f"FAIL test_NOT_FOUND_reject_post_with_unknown_token: {e}")

class TestDeletePost:
    def launch(self):
        print("# Test Delete Post")
        self.test_NOT_FOUND_delete_unexisting_post()
        self.test_FORBIDDEN_delete_post_without_token()
        self.test_FORBIDDEN_user_deletes_another_users_post()
        self.test_OK_author_deletes_own_post()

    def test_OK_author_deletes_own_post(self):
        try:
            token = auth.Auth().getAdminToken()
            if not token:
                raise Exception('Токен не получен!')
            resp = requests.delete("http://localhost/backend/posts/delete", params={
                "idPost": 3,
            }, headers={
                "auth-token": token,
            })
            if resp.status_code != 200:
                raise Exception(f"Статус не равен 200: {resp.status_code} - {resp.json()}")
            print("OK test_OK_author_deletes_own_post")
        except Exception as e:
            print(f"FAIL test_OK_author_deletes_own_post: {e}")

    def test_NOT_FOUND_delete_unexisting_post(self):
        try:
            token = auth.Auth().getAdminToken()
            if not token:
                raise Exception('Токен не получен!')
            resp = requests.delete("http://localhost/backend/posts/delete", params={
                "idPost": 22321,
            }, headers={
                "auth-token": token,
            })
            if resp.status_code != 404:
                raise Exception(f"Статус не равен 404: {resp.status_code} - {resp.json()}")
            print("OK test_NOT_FOUND_delete_unexisting_post")
        except Exception as e:
            print(f"FAIL test_NOT_FOUND_delete_unexisting_post: {e}")

    def test_FORBIDDEN_delete_post_without_token(self):
        try:
            token = auth.Auth().getAdminToken()
            if not token:
                raise Exception('Токен не получен!')
            resp = requests.delete("http://localhost/backend/posts/delete", params={
                "idPost": 3,
            }, headers={
                #! "auth-token": token,
            })
            if resp.status_code != 403:
                raise Exception(f"Статус не равен 403: {resp.status_code} - {resp.json()}")
            print("OK test_FORBIDDEN_delete_post_without_token")
        except Exception as e:
            print(f"FAIL test_FORBIDDEN_delete_post_without_token: {e}")

    def test_FORBIDDEN_user_deletes_another_users_post(self):
        try:
            token = auth.Auth().getBlogerToken()
            if not token:
                raise Exception('Токен не получен!')
            resp = requests.delete("http://localhost/backend/posts/delete", params={
                "idPost": 3,
            }, headers={
                "auth-token": token,
            })
            if resp.status_code != 403:
                raise Exception(f"Статус не равен 403: {resp.status_code} - {resp.json()}")
            print("OK test_FORBIDDEN_user_deletes_another_users_post")
        except Exception as e:
            print(f"FAIL test_FORBIDDEN_user_deletes_another_users_post: {e}")


def launch():
    print("\n--- Posts Controller ---\n")
    TestGetUserPosts().launch()
    TestGetActualPosts().launch()
    TestGetActualPost().launch()
    TestGetPostsNeedToApprove().launch()
    TestGetPostNeedToApprove().launch()
    TestCreatePost().launch()
    TestApprovePost().launch()
    TestRejectPost().launch()
    TestDeletePost().launch()
