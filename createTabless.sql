-- Удалить таблицу News, если она существует
DROP TABLE IF EXISTS "News";

-- Удалить таблицу NewsCategories, если она существует
DROP TABLE IF EXISTS "NewsCategories";

-- Создание таблицы News
CREATE TABLE "News" (
  "Id" bigserial PRIMARY KEY,
  "Title" text NOT NULL,
  "Content" text NOT NULL
);

-- Создание таблицы NewsCategories
CREATE TABLE "NewsCategories" (
  "NewsId" bigint NOT NULL,
  "CategoryId" bigint NOT NULL
);

-- Остальные операторы...


-- Создаем индексы для таблицы "NewsCategories"
CREATE INDEX IF NOT EXISTS "Idx_NewsCategories_NewsId" ON "NewsCategories" ("NewsId");
CREATE INDEX IF NOT EXISTS "Idx_NewsCategories_CategoryId" ON "NewsCategories" ("CategoryId");

CREATE UNIQUE INDEX "Unique_Title_Index" ON "News" ("Title");


INSERT INTO "News" ("Title", "Content") VALUES ('first', 'tratata');

INSERT INTO "NewsCategories" ("NewsId", "CategoryId") VALUES (1,1);

-- Подтвердить транзакцию и сохранить данные
COMMIT;