# RESTfull API on Golang.
TODO LIST

*Настроить авторизацию по уровню доступа.
уровни доступа нужно настроить через jwt token.
1. Создать клаймс которые хранят уровень доступа.
2. Создать хендлерсы которые будут обрабатывать каждый уровень доступа.
3. Прописть роутинг согласно уроням доспуа к каждому урлу. (лучше поделить на группы и в каждой группе прописать миддлваре с соответствующим обработчиком хендлерсом.)

23.08.2024.
создан миддле варе, package auth/ authAdminMiddlware нужно проверить роботостпособность. Протестировать.

25/08/2024
Создать добавление обьектов.