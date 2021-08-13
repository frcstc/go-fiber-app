/**
 * @Author: AF
 * @Date: 2021/8/9 14:22
 */

package config

import "os"

func Config(key string) string {
	return os.Getenv(key)
}
