package request

// type Sort string

// const (
// 	CREATE Sort = "CREATE" //通过创建日期排序
// 	UPDATE Sort = "UPDATE" //通过修改日期排序
// 	EYE    Sort = "EYE"    //通过浏览量排序
// 	LIKE   Sort = "LIKE"   //通过点赞量排序
// 	BACK   Sort = "BACK"   //通过创建日期倒叙
// 	SIZE   Sort = "SIZE"   //文件大小正序
// 	BSIZE  Sort = "BSIZE"  //文件大小倒叙
// )

// // GetOrderString 博客列表排序方式
// func (sort Sort) GetOrderString(prefix string) string {
// 	switch sort {
// 	case CREATE:
// 		return prefix + "created_at desc"
// 	case UPDATE:
// 		return prefix + "updated_at  desc"
// 	case EYE:
// 		return prefix + "eye_count desc"
// 	case LIKE:
// 		return prefix + "like_count desc"
// 	case BACK:
// 		return prefix + "created_at asc"
// 	case SIZE:
// 		return prefix + "size desc"
// 	case BSIZE:
// 		return prefix + "size asc"
// 	default:
// 		return prefix + "created_at desc"
// 	}
// }
