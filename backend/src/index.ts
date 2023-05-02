import { Bootstrap } from "./app"

function Start() {
    Bootstrap()
    .then((app) => {
        app.listen(8001, async () => {
            console.log("Сервер запущен на 8001");
            
        })
    })
}

Start()