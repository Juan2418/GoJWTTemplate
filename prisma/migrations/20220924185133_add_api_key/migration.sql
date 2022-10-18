/*
  Warnings:

  - A unique constraint covering the columns `[apiKey]` on the table `Family` will be added. If there are existing duplicate values, this will fail.
  - The required column `apiKey` was added to the `Family` table with a prisma-level default value. This is not possible if the table is not empty. Please add this column as optional, then populate it before making it required.

*/
-- AlterTable
ALTER TABLE `Family` ADD COLUMN `apiKey` VARCHAR(191) NOT NULL;

-- CreateIndex
CREATE UNIQUE INDEX `Family_apiKey_key` ON `Family`(`apiKey`);
