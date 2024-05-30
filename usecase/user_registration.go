package usecase

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"log"
	"strings"

	"git.finsoft.id/finsoft.id/go-example/db"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/argon2"
)

func Login(ctx context.Context, email, password string) error {
	user, err := Queries.GetUserByUserEmail(ctx, email)
	if err != nil {
		return errors.New("invalid username or password")
	}

	if !comparePasswordAndHash(password, user.Password) {
		return errors.New("invalid username or password")
	}

	return nil
}

func Register(ctx context.Context, email, password string, roleIds []string) error {
	tx, err := DbConn.Begin(ctx)
	if err != nil {
		log.Println(err)
		return err
	}

	defer tx.Rollback(ctx)
	queryTx := Queries.WithTx(tx)

	newUser, err := queryTx.CreateUser(ctx, db.CreateUserParams{
		Email:    email,
		Password: hashPassword(password),
	})

	if err != nil {
		log.Println(err)
		return err
	}

	for _, role := range roleIds {
		userRole := db.CreateUserRoleParams{
			UserID: newUser.ID,
			RoleID: uuid.MustParse(role),
		}

		queryTx.CreateUserRole(ctx, userRole)
	}

	if len(roleIds) == 0 {
		queryTx.CreateUserRole(ctx, db.CreateUserRoleParams{
			UserID: newUser.ID,
			RoleID: uuid.MustParse("0400185d-9276-4d23-af74-2e9f99353d12"),
		})
	}

	err = tx.Commit(ctx)
	if err != nil {
		log.Println("commit error:", err)
		return err
	}

	return nil
}

type argon2Params struct {
	Version       int
	MinMemorySize uint32
	MinIterations uint32
	Parallelism   uint8
	KeyLength     uint32
}

func hashPassword(password string) string {
	params := argon2Params{
		Version:       argon2.Version,
		MinMemorySize: uint32(12 * 1024),
		MinIterations: uint32(3),
		Parallelism:   uint8(1),
		KeyLength:     uint32(32),
	}

	salt := make([]byte, 16)
	rand.Read(salt)

	hash := argon2.IDKey([]byte(password), salt, params.MinIterations, params.MinMemorySize, params.Parallelism, params.KeyLength)
	b64Salt := base64.RawStdEncoding.Strict().EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.Strict().EncodeToString(hash)

	encoded := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", params.Version, params.MinMemorySize, params.MinIterations, params.Parallelism, b64Salt, b64Hash)
	return encoded
}

// func decodeHash(encoded string) (params argon2Params, decodedSalt, decodedHash []byte) {
// 	var salt, hash []byte

// 	_, err := fmt.Sscanf(encoded, "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", &params.Version, &params.MinMemorySize, &params.MinIterations, &params.Parallelism, &salt, &hash)
// 	if err != nil {
// 		return argon2Params{}, nil, nil
// 	}

// 	decodedSalt, _ = base64.RawStdEncoding.Strict().DecodeString(string(salt))
// 	decodedHash, _ = base64.RawStdEncoding.Strict().DecodeString(string(hash))

// 	return
// }

func decodeHash(encodedHash string) (params *argon2Params, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, errors.New("Invalid hash")
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, errors.New("Incompatible version")
	}

	params = &argon2Params{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &params.MinMemorySize, &params.MinIterations, &params.Parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	params.KeyLength = uint32(len(hash))

	return params, salt, hash, nil
}

func comparePasswordAndHash(password, encoded string) bool {
	params, salt, hash, err := decodeHash(encoded)
	if err != nil {
		return false
	}

	if len(salt) == 0 || len(hash) == 0 {
		return false
	}

	comparisonHash := argon2.IDKey([]byte(password), salt, params.MinIterations, params.MinMemorySize, params.Parallelism, params.KeyLength)
	if subtle.ConstantTimeCompare(hash, comparisonHash) == 1 {
		return true
	}

	return false
}
