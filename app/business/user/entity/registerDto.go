/**
 * @Author: AF
 * @Date: 2021/8/11 10:08
 */

package entity

type RegisterDto struct {
	Mobile       string `json:"mobile"`
	MobilePrefix string `json:"mobilePrefix"`
	Password     string `json:"password"`
}
