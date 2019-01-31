package sink

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"yunyinyue/spider/entity/common"
)

const (
	DBHost     = "localhost"
	DBPort     = ":3306"
	DBUser     = "root"
	DBPwd      = "zhujiajun"
	DBDatabase = "yunyinyue"
)

type MysqlSink struct {
	database *sql.DB
}

func NewMysqlSink() (sink MysqlSink, err error) {
	dbConnection := fmt.Sprintf("%s:%s@tcp(%s%s)/%s", DBUser, DBPwd, DBHost, DBPort,
		DBDatabase)
	db, err := sql.Open("mysql", dbConnection)
	if err != nil {
		return
	}
	sink = MysqlSink{}
	sink.database = db
	return
}

func (sink MysqlSink) Close() {
	sink.database.Close()
}

func (sink MysqlSink) InsertYunyinyueUser(user common.YunyinyueUser) (err error) {
	_, err = sink.database.Exec("insert into yunyinyue_user(user_id, nickname, user_type, location_info ) values (?,?,?,?)",
		user.UserId, user.Nickname, user.UserType, user.LocationInfo)
	return
}

func (sink MysqlSink) InsertComment(comment common.Comment) (err error) {
	_, err = sink.database.Exec("insert into comment(user_id, `time`, content, song_id) values (?,?,?, ?)",
		comment.User.UserId, comment.Time, comment.Content, comment.SongId)
	if err != nil {
		return
	}
	err = sink.InsertYunyinyueUser(comment.User)
	return

}
