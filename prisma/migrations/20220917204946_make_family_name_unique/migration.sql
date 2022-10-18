/*
  Warnings:

  - A unique constraint covering the columns `[name]` on the table `Family` will be added. If there are existing duplicate values, this will fail.

*/
-- CreateIndex
CREATE UNIQUE INDEX `Family_name_key` ON `Family`(`name`);
