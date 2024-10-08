package model

type AdAccountModel struct {
	*BaseModel
	ChannelId               uint64  `json:"channel_id" gorm:"column:channel_id;comment:渠道ID"`
	AppId                   uint64  `json:"app_id" gorm:"column:app_id;comment:应用渠道;size:100;"`                                               //应用渠道
	UserAccount             string  `json:"user_account" gorm:"column:user_account;comment:user账号名;size:100;"`                                //user账号名
	ButlerAccount           string  `json:"butler_account" gorm:"column:butler_account;comment:管家账号名;size:100;"`                              //管家账号名
	AccountNumber           string  `json:"account_number"  gorm:"column:account_number;comment:广告主账号;size:100;"`                             //广告主账号
	AccountTitle            string  `json:"account_title" gorm:"column:account_title;comment:账号名称;size:255;"`                                 //账号名称
	Type                    string  `json:"type" gorm:"column:type;comment:账号类型;size:32;"`                                                    //账号类型
	IsOnline                string  `json:"is_online"  gorm:"column:is_online;comment:账号状态 Y正常 N无效;"`                                         //账号状态 Y正常 N无效
	DataId                  string  `json:"data_id"  gorm:"column:data_id;comment:数据源id;size:255;"`                                           //数据源id
	DataId2                 string  `json:"data_id2"  gorm:"column:data_id2;comment:广点通WEB数据源;size:255;"`                                     //广点通WEB数据源
	AccessToken             string  `json:"access_token" gorm:"column:access_token;comment:accessToken值;size:1024;"`                          //accessToken值
	RefreshToken            string  `json:"refresh_token" gorm:"column:refresh_token;comment:refreshToken值;size:1024;"`                       //refreshToken值
	TokenDate               uint64  `json:"token_date"  gorm:"column:token_date;comment:token有效期;size:19;"`                                   //token有效期
	RefreshTokenDate        uint64  `json:"refresh_token_date"  gorm:"column:refresh_token_date;comment:refreshToken有效期;size:19;"`            //refreshToken有效期
	ClientId                string  `json:"client_id"  gorm:"column:client_id;comment:应用id;size:50;"`                                         //应用id
	Secret                  string  `json:"secret"  gorm:"column:secret;comment:应用秘钥;size:500;"`                                              //应用秘钥
	Url                     string  `json:"url"  gorm:"column:url;comment:监测链接;type:text"`                                                    //监测链接
	Url2                    string  `json:"url2"  gorm:"column:url2;comment:监测链接2;type:text"`                                                 //监测链接2
	Balance                 float64 `json:"balance"  gorm:"column:balance;comment:余额;"`                                                       //余额
	SqUrl                   string  `json:"sq_url" gorm:"column:sq_url;comment:授权url;"`                                                       //授权url
	LogidUrl                string  `json:"logid_url"  gorm:"column:logid_url;comment:百度落地页的url;size:512;"`                                   //百度落地页的url
	Token                   string  `json:"token"  gorm:"column:token;comment:百度回传信息的token;size:512;"`                                        //百度回传信息的token
	Mnhc                    string  `json:"mnhc"  gorm:"column:mnhc;comment:模拟回传的状态，Y：已回传 N:未回传;"`                                            //模拟回传的状态，Y：已回传 N:未回传
	Akey                    string  `json:"akey" gorm:"column:akey;comment:百度akey;size:255;"`                                                 //百度akey
	TencentCallbackType     uint64  `json:"tencent_callback_type"  gorm:"column:tencent_callback_type;comment:广点通数据源回传的方式，1：微信数据源，2：web数据源;"` //广点通数据源回传的方式，1：微信数据源，2：web数据源
	IsOpenGlobalCallback    string  `json:"is_open_global_callback"  gorm:"column:is_open_global_callback;comment:是否开启全局回传，Y：开启，N：未开启;"`      //是否开启全局回传，Y：开启，N：未开启
	GlobalCallbackTacticsId uint64  `json:"global_callback_tactics_id"  gorm:"column:global_callback_tactics_id;comment:全局回传策略id;"`           //全局回传策略id
	IsOpenPlanCallback      string  `json:"is_open_plan_callback"  gorm:"column:is_open_plan_callback;comment:是否开启计划回传，Y：开启，N：未开启;"`          //是否开启计划回传，Y：开启，N：未开启
	IsOpenSpreadCallback    string  `json:"is_open_spread_callback" gorm:"column:is_open_spread_callback;comment:是否开启推广链接回传，Y：开启，N：未开启;"`     //是否开启推广链接回传，Y：开启，N：未开启
	Uid                     uint64  `json:"uid"  gorm:"column:uid;comment:uid;"`                                                              //账号归属用户id
	ClueToken               string  `json:"clue_token"  gorm:"column:clue_token;comment:字节回调token;size:500"`
	ByteLink                string  `json:"byte_link" gorm:"column:byte_link;comment:字节落地页的url;size:500;"`
	GroupName               string  `json:"group_name"  gorm:"column:group_name;comment:分配账号组;" `
	NickName                string  `json:"nick_name" gorm:"column:nick_name;comment:用户昵称;" `
	CreatedBy               uint64  `json:"created_by" gorm:"column:created_by;comment:创建者"`
	UpdatedBy               uint64  `json:"updated_by" gorm:"column:updated_by;comment:更新者"`
	DeletedBy               uint64  `json:"deleted_by" gorm:"column:deleted_by;comment:删除者"`
}

// TableName 表名
func (m *AdAccountModel) TableName() string {
	return "nv_ad_account"
}
