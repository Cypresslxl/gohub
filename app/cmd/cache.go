package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gohub/pkg/cache"
	"gohub/pkg/console"
)

var CmdCache = &cobra.Command{
	Use:   "cache",
	Short: "Cache management",
}

var CmdCacheClear = &cobra.Command{
	Use:   "clear",
	Short: "Clear cache",
	Run:   runCacheClear,
}

var CmdCacheForget = &cobra.Command{
	Use:   "forget",
	Short: "Delete redis key,example: cache forget cache-key",
	Run:   runCacheForget,
}

// forget 命令选项
var cacheKey string

func init() {
	//    注册cache命令的子命令
	CmdCache.AddCommand(
		CmdCacheClear,
		CmdCacheForget,
	)

	//	设置cache forget 命令的选项
	CmdCacheForget.Flags().StringVarP(&cacheKey, "key", "k", "", "KEY of the cache")

	CmdCacheForget.MarkFlagRequired("key")
}

func runCacheForget(cmd *cobra.Command, args []string) {
	//fmt.Println(cacheKey)
	cache.Forget(cacheKey)
	console.Success(fmt.Sprintf("Cache key [%s] deleted.", cacheKey))
}

func runCacheClear(cmd *cobra.Command, args []string) {
	cache.Flush()
	console.Success("Cache cleared.")
}
