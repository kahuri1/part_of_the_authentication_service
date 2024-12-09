package repository

import (
	sql2 "database/sql"
	"encoding/base64"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/kahuri1/part_of_the_authentication_service/iternal/domain"
	"github.com/kahuri1/part_of_the_authentication_service/iternal/model"
	"github.com/spf13/viper"
	"time"
)

func (r *Repository) CreateSessionRepo(auth model.AuthenticationRequest, refreshToken []byte) error {
	sql, args, err := sq.
		Insert("refreshSessions").
		Columns("user_uuid", "ip", "refresh_token", "expires_at", "created_at").
		Values(auth.Uuid, auth.Ip, refreshToken, time.Now().Add(viper.GetDuration("auth.refreshTokenTTL")), time.Now()).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to create refreshSessions  request: %w", err)
	}

	_, err = r.db.Exec(sql, args...)
	if err != nil {
		return fmt.Errorf("failed to process refreshSessions  query: %w", err)
	}
	return nil
}

func (r *Repository) GetRefreshSessionByRefreshTokenRepo(token model.Tokens) (model.RefreshSession, error) {
	var session model.RefreshSession

	decodeString, err := base64.StdEncoding.DecodeString(token.RefreshToken)
	if err != nil {
		return session, fmt.Errorf("failed decoding refresh token: %w", err)
	}
	sql, args, err := sq.
		Select("rs.user_uuid, rs.ip, rs.refresh_token, rs.expires_at, rs.created_at, u.email").
		From("refreshSessions rs").
		Join("users u ON rs.user_uuid = u.uuid").
		Where(sq.Eq{"rs.refresh_token": decodeString}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return session, fmt.Errorf("failed to create query to get refresh session: %w", err)
	}

	err = r.db.QueryRow(sql, args...).Scan(&session.UserUuid, &session.Ip, &session.RefreshToken, &session.ExpiresAt, &session.CreatedAt, &session.Email)
	if err != nil {

		return session, fmt.Errorf("failed to execute query to get refresh session: %w", err)
	}

	return session, nil
}

func (r *Repository) UpdateRefreshSessionRepo(uuid string, ip string, refreshToken []byte) error {
	sql, args, err := sq.
		Update("refreshSessions").
		Set("refresh_token", refreshToken).
		Set("expires_at", time.Now().Add(viper.GetDuration("auth.refreshTokenTTL"))).
		Where(sq.Eq{"user_uuid": uuid}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return fmt.Errorf("failed to create update refresh session query: %w", err)
	}

	_, err = r.db.Exec(sql, args...)
	if err != nil {
		return fmt.Errorf("failed to process update refresh session query: %w", err)
	}

	return nil
}

func (r *Repository) CheckSessionByUuidUserRepo(uuid string) (model.RefreshSession, error) {
	var session model.RefreshSession
	sql, args, err := sq.
		Select("user_uuid, ip, refresh_token, expires_at, created_at").
		From("refreshSessions").
		Where(sq.Eq{"user_uuid": uuid}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return session, fmt.Errorf("failed to create query to check session by user UUID: %w", err)
	}
	err = r.db.QueryRow(sql, args...).Scan(&session.UserUuid, &session.Ip, &session.RefreshToken, &session.ExpiresAt, &session.CreatedAt)

	if err != nil {
		if err == sql2.ErrNoRows {
			return session, nil
		}
		return session, fmt.Errorf("failed to execute query to check session by user UUID: %w", err)
	}
	return session, domain.ErrSessionOpen
}
