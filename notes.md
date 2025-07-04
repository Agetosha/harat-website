<!-- установка ангуляр -->

sudo apt install nodejs npm
npm install -g @angular/cli

<!-- запустить приложение -->

ng serve

<!-- создание нового компонента -->

ng generate component todo

<!-- filepath: src/app/app.html -->

<app-todo></app-todo>
<router-outlet />
