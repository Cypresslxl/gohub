package migrate

import (
	"database/sql"

	"gorm.io/gorm"
)

// migrationFunc 定义 up 和 down 回调方法的类型
type migrationFunc func(gorm.Migrator, *sql.DB)

// migrationFiles 所有的迁移文件数组
var migrationFiles []MigrationFile

// MigrationFile 代表着单个迁移文件
// 我们先来创建 MigrationFile 对象，这个对象对应单个的迁移文件。
type MigrationFile struct {
	Up       migrationFunc
	Down     migrationFunc
	FileName string
}

// Add 新增一个迁移文件，所有的迁移文件都需要调用此方法来注册
func Add(name string, up migrationFunc, down migrationFunc) {
	migrationFiles = append(migrationFiles, MigrationFile{
		FileName: name,
		Up:       up,
		Down:     down,
	})
}

// getMigrationFile 通过迁移文件的名称来获取到 MigrationFile 对象
func getMigrationFile(fileName string) MigrationFile {
	for _, file := range migrationFiles {
		if fileName == file.FileName {
			return file
		}
	}
	return MigrationFile{}
}

// isNotMigrated 判断迁移是否已执行
func (mfile MigrationFile) isNotMigrated(migrations []Migration) bool {
	for _, v := range migrations {
		if v.Migration == mfile.FileName { //还能在数组中找到说明还没有迁移，迁移之后是会把值从数组中移除的
			return false
		}
	}
	return true
}
