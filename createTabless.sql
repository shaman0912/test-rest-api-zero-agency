-- Структура таблицы "News"
CREATE TABLE IF NOT EXISTS "News" (
  "Id" bigserial PRIMARY KEY,
  "Title" text NOT NULL,
  "Content" text NOT NULL
);

-- Структура таблицы "NewsCategories"
CREATE TABLE IF NOT EXISTS "NewsCategories" (
  "NewsId" bigint NOT NULL,
  "CategoryId" bigint NOT NULL
);

-- Создаем индексы для таблицы "News"
CREATE INDEX IF NOT EXISTS "Idx_News_Id" ON "News" ("Id");

-- Создаем индексы для таблицы "NewsCategories"
CREATE INDEX IF NOT EXISTS "Idx_NewsCategories_NewsId" ON "NewsCategories" ("NewsId");
CREATE INDEX IF NOT EXISTS "Idx_NewsCategories_CategoryId" ON "NewsCategories" ("CategoryId");
