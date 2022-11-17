# blog-project

Проект Блога. Клиент-серверное приложение на платформе Docker.
Сервисы связаны между собой при помощи прокси-сервера Nginx.

<ul>
  <li>Backend (Серверная часть) написана на Golang/Gin/sqlx/go-redis.</li> 
<li>Осуществлено end-to-end тестирование при помощи скриптов на Python.</li> 
<li>Frontend (Клиентская часть) написана на JavaScript/React/Redux/TailwindCSS.</li> 
</ul>
  
Средство администрирования PostgreSQL - adminer.


Базы данных: 
  <ul>
  <li>PostgreSQL как основная база хранения данных.</li> 
  <li>Redis как база для кэширования и хранения сеансов пользователей</li> 
  </ul>
Аутентификация по токену.



![Снимок экрана от 2022-11-17 15-04-51](https://user-images.githubusercontent.com/60382252/202444072-6f2c3eb0-e566-475c-b4c7-08386d2365dd.png)
