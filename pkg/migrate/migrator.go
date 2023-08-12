// Package migrate 处理数据库迁移
package migrate

import (
	"gohub/pkg/console"
	"gohub/pkg/database"
	"gohub/pkg/file"
	"os"

	"gorm.io/gorm"
)

// Migrator 数据迁移操作类
// Migrator 是最终的操作对象，负责创建 migrations 数据表，以及拥有 up 、down、reset、refresh 等动作。
type Migrator struct {
	Folder   string
	DB       *gorm.DB
	Migrator gorm.Migrator
}

// Migration 对应数据的 migrations 表里的一条数据,ID代表编号，Batch代表执行的次数
type Migration struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;"`
	Migration string `gorm:"type:varchar(255);not null;unique;"`
	Batch     int
}

// NewMigrator 创建 Migrator 实例，用以执行迁移操作
func NewMigrator() *Migrator {

	// 初始化必要属性
	migrator := &Migrator{
		Folder:   "database/migrations/",
		DB:       database.DB,
		Migrator: database.DB.Migrator(),
	}
	// migrations 不存在的话就创建它
	migrator.createMigrationsTable()

	return migrator
}

// 创建 migrations 表
func (migrator *Migrator) createMigrationsTable() {

	migration := Migration{}

	// 不存在才创建
	if !migrator.Migrator.HasTable(&migration) {
		migrator.Migrator.CreateTable(&migration)
	}
}

// Up 执行所有未迁移过的文件
func (migrator *Migrator) Up() {
	// 读取所有迁移文件，确保按照时间排序
	migrateFiles := migrator.readAllMigrationFiles()
	// 获取当前批次的值
	batch := migrator.getBatch()

	// 获取所有迁移数据
	migrations := []Migration{}
	migrator.DB.Find(&migrations) //Find finds all records matching given conditions conds
	// 可以通过此值来判断数据库是否已是最新
	runed := false

	// 对迁移文件进行遍历，如果没有执行过，就执行 up 回调
	for _, mfile := range migrateFiles {

		// 对比文件名称，看是否已经运行过
		if mfile.isNotMigrated(migrations) {
			migrator.runUpMigration(mfile, batch)
			runed = true
		}
	}

	if !runed {
		console.Success("database is up to data.")
	}
}

// 获取当前这个批次的值
func (migrator *Migrator) getBatch() int {
	//默认为 1
	batch := 1

	// 取最后执行的一条迁移数据
	lastMigration := Migration{}
	migrator.DB.Order("is DESC").First(&lastMigration)

	// 如果有值的话，加一
	if lastMigration.ID > 0 {
		batch = lastMigration.Batch + 1
	}
	return batch
}

// 从文件目录读取文件，保证正确的时间排序
func (migrator *Migrator) readAllMigrationFiles() (migrateFiles []MigrationFile) {
	// 读取 database/migrations/ 目录下的所有文件
	// 默认是会按照文件名称进行排序
	files, err := os.ReadDir(migrator.Folder)
	console.ExitIf(err)

	//var migrateFiles []MigrationFile
	for _, f := range files {
		//	去除文件后缀 .go
		filename := file.FileNameWithoutExtension(f.Name())

		// 通过迁移文件的名称获取『MigrationFile』对象
		mfile := getMigrationFile(filename)

		// 加个判断，确保迁移文件可用，再放进 migrateFiles 数组中
		if len(mfile.FileName) > 0 {
			migrateFiles = append(migrationFiles, mfile)
		}
	}
	//return migrateFiles
	return
}

// 执行迁移，执行迁移的 up 方法
func (migrator *Migrator) runUpMigration(mfile MigrationFile, batch int) {
	// 执行 up 区块的 SQL
	if mfile.Up != nil {
		//	友好提示
		console.Warning("migrating " + mfile.FileName)
		//执行up方法
		mfile.Up(database.DB.Migrator(), database.SQLDB)
		//提示已经迁移了哪个文件
		console.Success("migrated" + mfile.FileName)
	}

	//	入库
	err := migrator.DB.Create(&Migration{
		Migration: mfile.FileName,
		Batch:     batch,
	}).Error
	console.ExitIf(err)
}
