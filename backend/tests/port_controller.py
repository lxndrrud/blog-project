import requests


def test_OK_actual_posts():
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

toTest = [
    test_OK_actual_posts
]



for func_ in toTest:
    func_()