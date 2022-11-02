import requests

def test_OK_get_user_profile():
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

def test_NOT_FOUND_get_user_profile():
    try:
        resp = requests.get("http://localhost/backend/users/profile?idUser=2")
        if resp.status_code != 404:
            raise Exception(f"Статус не равен 404: {resp.status_code}")
        
        print(f"OK test_NOT_FOUND_get_user_profile")
    except Exception as e:
        print(f"FAIL test_NOT_FOUND_get_user_profile: {e}")

toTest = [
    test_OK_get_user_profile,
    test_NOT_FOUND_get_user_profile
]

for func_ in toTest:
    func_()