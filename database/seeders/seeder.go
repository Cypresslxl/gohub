// Package seeders 存放数据填充文件
package seeders

import "gohub/pkg/seed"

// Initialize 当我们调用 seeders.Initialize 方法时，就会触发 seeders 包里所有的 init 方法，完成对 Seeder 文件的自动添加操作。
func Initialize() {

	// 触发加载本目录下其他文件中的 init 方法

	// 指定优先于同目录下的其他文件运行
	seed.SetRunOrder([]string{
		"SeedUsersTable",
	})
}
