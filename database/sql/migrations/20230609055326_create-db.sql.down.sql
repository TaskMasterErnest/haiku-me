ALTER TABLE snippets DELETE FROM snippets;
ALTER TABLE snippets DROP INDEX idx_snippets_created ON snippets;
DROP TABLE snippets;
DROP DATABASE snippetbox;