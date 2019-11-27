package graph

// 商城站点信息集合 (ms_mall)
// 字段： 商城编号、商城标识（用于自定义访问域名）商城名称、商城备注描述、商城客服电话、客服邮箱、商城域名
type Mall struct {
	MallId   string `bson:"mallId" json:"mallId"`
	Identify string `bson:"identify" json:"identify"`
	Name     string `bson:"name" json:"name"`
	Title    string `bson:"title" json:"title"`
	Remark   string `bson:"remark" json:"remark"`
	LogoUrl  string `bson:"logoUrl" json:"logoUrl"`
	Domain   string `bson:"domain" json:"domain"`

	// 商城广告信息管理，如Banner、侧边广告
	Ad []*Ad `bson:"ad" json:"ad"`
}

// 广告
// key值为广告位名称，如：banner
type Ad struct {
	Id     string `bson:"id" json:"id"`
	Key    string `bson:"key" json:"key"`
	ImgUrl string `bson:"imgUrl" json:"imgUrl"`
	Href   string `bson:"href" json:"href"`
	Title  string `bson:"title" json:"title"`
	Remark string `bson:"remark" json:"remark"`

	// 是否展示(1=展示 0=隐藏)
	Show string
	// 正序排序
	Rank int
}
