-- +goose up
ALTER TABLE IF EXISTS "board_column"
RENAME TO "project_column";

ALTER TABLE IF EXISTS "project_column"
RENAME COLUMN board_id to project_id;

ALTER TABLE IF EXISTS `project_column`
ADD CONSTRAINT `fk_name`
    FOREIGN KEY (`fk_table2_id`) REFERENCES `table2` (`t2`) ON DELETE CASCADE;
