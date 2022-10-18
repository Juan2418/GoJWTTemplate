-- CreateTable
CREATE TABLE `Category` (
    `id` INTEGER NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(191) NOT NULL,
    `description` VARCHAR(191) NOT NULL,
    `imageURL` VARCHAR(191) NOT NULL,
    `monthlySpendingLimit` INTEGER NOT NULL DEFAULT 0,
    `familyId` INTEGER NOT NULL,
    `createdAt` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updatedAt` DATETIME(3) NOT NULL,
    `deleted` DATETIME(3) NULL,

    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- AddForeignKey
ALTER TABLE `Category` ADD CONSTRAINT `Category_familyId_fkey` FOREIGN KEY (`familyId`) REFERENCES `Family`(`id`) ON DELETE RESTRICT ON UPDATE CASCADE;
