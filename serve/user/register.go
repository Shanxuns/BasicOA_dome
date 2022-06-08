package user

import (
	"BasicOA/serve/email"
	"BasicOA/serve/md5"
	"BasicOA/serve/mysql"
	_type "BasicOA/type"
	"errors"
	"io/ioutil"
	"strconv"
	"time"
)

func Register(User _type.User) error {
	var err error
	// Mysql查找账户
	RowUser := mysql.Db.QueryRow(mysql.SearchEmail, User.Email)
	if err = RowUser.Scan(&User.Id, &User.Email, &User.Fname, &User.Lname, &User.Phone, &User.Password, &User.Confuse, &User.Avatar, &User.License, &User.Datetime); err == nil {
		return errors.New("账号已注册")
	}
	// 头像
	if User.Avatar != "" {
		Avatar := User.Avatar
		User.Avatar = md5.Hash(strconv.FormatInt(time.Now().UnixNano(), 10))
		err = ioutil.WriteFile("./avatar/"+User.Avatar+".image", []byte(Avatar), 0666)
		if err != nil {
			return errors.New("上传头像失败")
		}
		Avatar = ""
	}
	// 创建加密盐
	User.Confuse = md5.GenerateSubId()
	// 加密明文密码
	md5.SetHash(User.Confuse)
	User.Password = md5.Hash(User.Password)
	// 用户信息写入数据库
	if _, err = mysql.Db.Exec(mysql.Register, &User.Email, &User.Fname, &User.Lname, &User.Phone, &User.Password, &User.Confuse, &User.Avatar); err != nil {
		return errors.New("注册失败")
	}
	// 启动邮箱
	md5.SetHash("Activation")
	go email.SendEmail(User.Email, "BasicOA 协调办公 激活邮件", "<head>\n    <base target=\"_blank\" />\n    <style type=\"text/css\">::-webkit-scrollbar{ display: none; }</style>\n    <style id=\"cloudAttachStyle\" type=\"text/css\">#divNeteaseBigAttach, #divNeteaseBigAttach_bak{display:none;}</style>\n    <style id=\"blockquoteStyle\" type=\"text/css\">blockquote{display:none;}</style>\n    <style type=\"text/css\">\n        body{font-size:14px;font-family:arial,verdana,sans-serif;line-height:1.666;padding:0;margin:0;overflow:auto;white-space:normal;word-wrap:break-word;min-height:100px}\n        td, input, button, select, body{font-family:Helvetica, 'Microsoft Yahei', verdana}\n        pre {white-space:pre-wrap;white-space:-moz-pre-wrap;white-space:-pre-wrap;white-space:-o-pre-wrap;word-wrap:break-word;width:95%}\n        th,td{font-family:arial,verdana,sans-serif;line-height:1.666}\n        img{ border:0}\n        header,footer,section,aside,article,nav,hgroup,figure,figcaption{display:block}\n        blockquote{margin-right:0px}\n    </style>\n</head>\n<body tabindex=\"0\" role=\"listitem\">\n<table width=\"700\" border=\"0\" align=\"center\" cellspacing=\"0\" style=\"width:700px;\">\n    <tbody>\n    <tr>\n        <td>\n            <div style=\"width:700px;margin:0 auto;border-bottom:1px solid #ccc;margin-bottom:30px;\">\n                <table border=\"0\" cellpadding=\"0\" cellspacing=\"0\" width=\"700\" height=\"39\" style=\"font:12px Tahoma, Arial, 宋体;\">\n                    <tbody><tr><td width=\"210\"></td></tr></tbody>\n                </table>\n            </div>\n            <div style=\"width:680px;padding:0 10px;margin:0 auto;\">\n                <div style=\"line-height:1.5;font-size:14px;margin-bottom:25px;color:#4d4d4d;\">\n                    <strong style=\"display:block;margin-bottom:15px;\">尊敬的用户：<span style=\"color:#f60;font-size: 16px;\"></span>您好！</strong>\n                    <strong style=\"display:block;margin-bottom:15px;\">\n                        您正在进行<span style=\"color: red\">激活账号</span>操作，<span style=\"color:#f60;font-size: 24px\"></br>用户邮箱："+User.Email+"<p>激活码："+md5.Hash(User.Confuse)+"</span></a>\n                    </strong>\n                </div>\n                <div style=\"margin-bottom:30px;\">\n                    <small style=\"display:block;margin-bottom:20px;font-size:12px;\">\n                        <p style=\"color:#747474;\">\n                            注意：此操作可能绑定邮箱。如非本人操作，请勿激活，账号将在七天后自动注销。\n                            <br>（请勿泄漏,帮他人注册查证后，账号将注销。)\n                        </p>\n                    </small>\n                </div>\n            </div>\n            <div style=\"width:700px;margin:0 auto;\">\n                <div style=\"padding:10px 10px 0;border-top:1px solid #ccc;color:#747474;margin-bottom:20px;line-height:1.3em;font-size:12px;\">\n                    <p>此为系统邮件，请勿回复<br>\n                        请保管好您的邮箱，避免账号被他人盗用\n                    </p>\n                    <p>天津无铭氏科技有限公司</p>\n                </div>\n            </div>\n        </td>\n    </tr>\n    </tbody>\n</table>\n</body>")
	// 响应结果
	return nil
}
