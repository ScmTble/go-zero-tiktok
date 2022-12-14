// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	usersFieldNames          = builder.RawFieldNames(&Users{})
	usersRows                = strings.Join(usersFieldNames, ",")
	usersRowsExpectAutoSet   = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	usersRowsWithPlaceHolder = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheDouyinUsersIdPrefix       = "cache:douyin:users:id:"
	cacheDouyinUsersUsernamePrefix = "cache:douyin:users:username:"
)

type (
	usersModel interface {
		Insert(ctx context.Context, data *Users) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Users, error)
		FindOneByUsername(ctx context.Context, username string) (*Users, error)
		Update(ctx context.Context, data *Users) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUsersModel struct {
		sqlc.CachedConn
		table string
	}

	Users struct {
		Id       int64  `db:"id"`       // 用户id，自增主键
		Name     string `db:"name"`     // 用户名
		Password string `db:"password"` // 用户密码
	}
)

func newUsersModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUsersModel {
	return &defaultUsersModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`users`",
	}
}

func (m *defaultUsersModel) Delete(ctx context.Context, id int64) error {
	douyinUsersIdKey := fmt.Sprintf("%s%v", cacheDouyinUsersIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, douyinUsersIdKey)
	return err
}

func (m *defaultUsersModel) FindOne(ctx context.Context, id int64) (*Users, error) {
	douyinUsersIdKey := fmt.Sprintf("%s%v", cacheDouyinUsersIdPrefix, id)
	var resp Users
	err := m.QueryRowCtx(ctx, &resp, douyinUsersIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usersRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) FindOneByUsername(ctx context.Context, username string) (*Users, error) {
	cacheDouyinUsersUsernamePrefix := fmt.Sprintf("%s%s", cacheDouyinUsersIdPrefix, username)
	var resp Users
	err := m.QueryRowCtx(ctx, &resp, cacheDouyinUsersUsernamePrefix, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", usersRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, username)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) Insert(ctx context.Context, data *Users) (sql.Result, error) {
	douyinUsersIdKey := fmt.Sprintf("%s%v", cacheDouyinUsersIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, usersRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Name, data.Password)
	}, douyinUsersIdKey)
	return ret, err
}

func (m *defaultUsersModel) Update(ctx context.Context, data *Users) error {
	douyinUsersIdKey := fmt.Sprintf("%s%v", cacheDouyinUsersIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usersRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Name, data.Password, data.Id)
	}, douyinUsersIdKey)
	return err
}

func (m *defaultUsersModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheDouyinUsersIdPrefix, primary)
}

func (m *defaultUsersModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usersRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUsersModel) tableName() string {
	return m.table
}
