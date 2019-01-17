"# xsolla CRUD GO" 

Для работы приложения необходимо создать DB MySql "golangdb"
User: root
Pass: (пустой)

Далее создаем таблицу "Books" c с полями "id" (int), "name" (varchar)
или воспользоваться sql запросом
```sql
CREATE TABLE `books` (
  `id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```
