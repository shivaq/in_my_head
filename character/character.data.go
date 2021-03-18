package character

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/shivaq/in_my_head/database"
)

func getCharacter(characterID int) (*Character, error) {
	// Context
	// デッドラインを設定したり、シグナルキャンセルしたり、
	// APIやプロセスの境界を超えて、他のリクエストの値を取得できる
	// Context デ、
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// ID 指定で select query
	row := database.DbConn.QueryRowContext(ctx, `SELECT
	characterId,
	UserID,
	characterType,
	circleColorName,
	savedFileName,
	blobImage,
	isPicUsed,
	nickName,
	firstName,
	middleName,
	lastName,
	originalName,
	isImaginary,
	birthDate,
	description,
	level,
	isLiked,
	registeredDate,
	updatedDate,
	onDeleteLock,
	inLimbo,
	reserveNumber01,
	reserveNumber02,
	reserveNumber03,
	reserveNumber04,
	reserveNumber05,
	reserveNumber06,
	reserveNumber07,
	reserveNumber08,
	reserveNumber09,
	reserveNumber10,
	reserveNumber11
	FROM characters
	WHERE characterId = ?`, characterID)

	// character を定義して
	character := &Character{}
	// クエリ結果を character に当てはめていく
	err := row.Scan(
		&character.CharacterID,
		&character.UserID,
		&character.CharacterType,
		&character.CircleColorName,
		&character.SavedFileName,
		&character.BlobImage,
		&character.IsPicUsed,
		&character.NickName,
		&character.FirstName,
		&character.MiddleName,
		&character.LastName,
		&character.OriginalName,
		&character.IsImaginary,
		&character.BirthDate,
		&character.Description,
		&character.Level,
		&character.IsLiked,
		&character.RegisteredDate,
		&character.UpdatedDate,
		&character.OnDeleteLock,
		&character.InLimbo,
		&character.ReserveNumber01,
		&character.ReserveNumber02,
		&character.ReserveNumber03,
		&character.ReserveNumber04,
		&character.ReserveNumber05,
		&character.ReserveNumber06,
		&character.ReserveNumber07,
		&character.ReserveNumber08,
		&character.ReserveNumber09,
		&character.ReserveNumber10,
		&character.ReserveNumber11,
	)
	// クエリ結果 なし だった場合
	if err == sql.ErrNoRows {
		return nil, nil
		// それ以外のエラーの場合
	} else if err != nil {
		log.Println(err)
		return nil, err
	}
	return character, nil
}

func GetTopTenCharacters() ([]Character, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// ORDER BY quantityOnHand DESC LIMIT 10
	results, err := database.DbConn.QueryContext(ctx, `SELECT
	characterId,
	UserID,
	characterType,
	circleColorName,
	savedFileName,
	blobImage,
	isPicUsed,
	nickName,
	firstName,
	middleName,
	lastName,
	originalName,
	isImaginary,
	birthDate,
	description,
	level,
	isLiked,
	registeredDate,
	updatedDate,
	onDeleteLock,
	inLimbo,
	reserveNumber01,
	reserveNumber02,
	reserveNumber03,
	reserveNumber04,
	reserveNumber05,
	reserveNumber06,
	reserveNumber07,
	reserveNumber08,
	reserveNumber09,
	reserveNumber10,
	reserveNumber11
	FROM characters ORDER BY updatedDate DESC LIMIT 10
	`)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer results.Close()
	characters := make([]Character, 0)
	for results.Next() {
		var character Character
		results.Scan(
			&character.CharacterID,
			&character.UserID,
			&character.CharacterType,
			&character.CircleColorName,
			&character.SavedFileName,
			&character.BlobImage,
			&character.IsPicUsed,
			&character.NickName,
			&character.FirstName,
			&character.MiddleName,
			&character.LastName,
			&character.OriginalName,
			&character.IsImaginary,
			&character.BirthDate,
			&character.Description,
			&character.Level,
			&character.IsLiked,
			&character.RegisteredDate,
			&character.UpdatedDate,
			&character.OnDeleteLock,
			&character.InLimbo,
			&character.ReserveNumber01,
			&character.ReserveNumber02,
			&character.ReserveNumber03,
			&character.ReserveNumber04,
			&character.ReserveNumber05,
			&character.ReserveNumber06,
			&character.ReserveNumber07,
			&character.ReserveNumber08,
			&character.ReserveNumber09,
			&character.ReserveNumber10,
			&character.ReserveNumber11)

		characters = append(characters, character)
	}
	return characters, nil
}

func removeCharacter(characterID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// ExecContext は、Row を返さない
	_, err := database.DbConn.ExecContext(ctx, `DELETE FROM characters where characterId = ?`, characterID)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func getCharacterList() ([]Character, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// すべてを取得
	results, err := database.DbConn.QueryContext(ctx, `SELECT
	characterId,
	UserID,
	characterType,
	circleColorName,
	savedFileName,
	blobImage,
	isPicUsed,
	nickName,
	firstName,
	middleName,
	lastName,
	originalName,
	isImaginary,
	birthDate,
	description,
	level,
	isLiked,
	registeredDate,
	updatedDate,
	onDeleteLock,
	inLimbo,
	reserveNumber01,
	reserveNumber02,
	reserveNumber03,
	reserveNumber04,
	reserveNumber05,
	reserveNumber06,
	reserveNumber07,
	reserveNumber08,
	reserveNumber09,
	reserveNumber10,
	reserveNumber11
	FROM characters`)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	// ここだけ Close しているのはなぜ？？
	defer results.Close()
	// characters を定義して
	characters := make([]Character, 0)
	// for ループで次の row に移動させつつ、
	for results.Next() {
		// このスコープで character を定義して
		var character Character
		// Scan 結果を Character に当てはめて
		results.Scan(
			&character.CharacterID,
			&character.UserID,
			&character.CharacterType,
			&character.CircleColorName,
			&character.SavedFileName,
			&character.BlobImage,
			&character.IsPicUsed,
			&character.NickName,
			&character.FirstName,
			&character.MiddleName,
			&character.LastName,
			&character.OriginalName,
			&character.IsImaginary,
			&character.BirthDate,
			&character.Description,
			&character.Level,
			&character.IsLiked,
			&character.RegisteredDate,
			&character.UpdatedDate,
			&character.OnDeleteLock,
			&character.InLimbo,
			&character.ReserveNumber01,
			&character.ReserveNumber02,
			&character.ReserveNumber03,
			&character.ReserveNumber04,
			&character.ReserveNumber05,
			&character.ReserveNumber06,
			&character.ReserveNumber07,
			&character.ReserveNumber08,
			&character.ReserveNumber09,
			&character.ReserveNumber10,
			&character.ReserveNumber11)
		// Slice に Append していく
		characters = append(characters, character)
	}
	return characters, nil
}

func updateCharacter(character Character) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if character.CharacterID == nil || *character.CharacterID == 0 {
		return errors.New("character has invalid ID")
	}
	_, err := database.DbConn.ExecContext(ctx, `UPDATE characters SET
		UserID=?,
		characterType=?,
		circleColorName=?,
		savedFileName=?,
		blobImage=?,
		isPicUsed=?,
		nickName=?,
		firstName=?,
		middleName=?,
		lastName=?,
		originalName=?,
		isImaginary=?,
		birthDate=?,
		description=?,
		level=?,
		isLiked=?,
		registeredDate=?,
		updatedDate=?,
		onDeleteLock=?,
		inLimbo=?,
		reserveNumber01=?,
		reserveNumber02=?,
		reserveNumber03=?,
		reserveNumber04=?,
		reserveNumber05=?,
		reserveNumber06=?,
		reserveNumber07=?,
		reserveNumber08=?,
		reserveNumber09=?,
		reserveNumber10=?,
		reserveNumber11
		WHERE CharacterID=?`,
		character.UserID,
		character.CharacterType,
		character.CircleColorName,
		character.SavedFileName,
		character.BlobImage,
		character.IsPicUsed,
		character.NickName,
		character.FirstName,
		character.MiddleName,
		character.LastName,
		character.OriginalName,
		character.IsImaginary,
		character.BirthDate,
		character.Description,
		character.Level,
		character.IsLiked,
		character.RegisteredDate,
		character.UpdatedDate,
		character.OnDeleteLock,
		character.InLimbo,
		character.ReserveNumber01,
		character.ReserveNumber02,
		character.ReserveNumber03,
		character.ReserveNumber04,
		character.ReserveNumber05,
		character.ReserveNumber06,
		character.ReserveNumber07,
		character.ReserveNumber08,
		character.ReserveNumber09,
		character.ReserveNumber10,
		character.ReserveNumber11,
		character.CharacterID)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func insertCharacter(character Character) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := database.DbConn.ExecContext(ctx, `INSERT INTO characters 
		(UserID,
		characterType,
		circleColorName,
		savedFileName,
		blobImage,
		isPicUsed,
		nickName,
		firstName,
		middleName,
		lastName,
		originalName,
		isImaginary,
		birthDate,
		description,
		level,
		isLiked,
		registeredDate,
		updatedDate,
		onDeleteLock,
		inLimbo,
		reserveNumber01,
		reserveNumber02,
		reserveNumber03,
		reserveNumber04,
		reserveNumber05,
		reserveNumber06,
		reserveNumber07,
		reserveNumber08,
		reserveNumber09,
		reserveNumber10,
		reserveNumber11
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		character.UserID,
		character.CharacterType,
		character.CircleColorName,
		character.SavedFileName,
		character.BlobImage,
		character.IsPicUsed,
		character.NickName,
		character.FirstName,
		character.MiddleName,
		character.LastName,
		character.OriginalName,
		character.IsImaginary,
		character.BirthDate,
		character.Description,
		character.Level,
		character.IsLiked,
		character.RegisteredDate,
		character.UpdatedDate,
		character.OnDeleteLock,
		character.InLimbo,
		character.ReserveNumber01,
		character.ReserveNumber02,
		character.ReserveNumber03,
		character.ReserveNumber04,
		character.ReserveNumber05,
		character.ReserveNumber06,
		character.ReserveNumber07,
		character.ReserveNumber08,
		character.ReserveNumber09,
		character.ReserveNumber10,
		character.ReserveNumber11)
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	// convert int64 to int
	return int(insertID), nil
}

func searchForCharacterData(characterFilter CharacterReportFilter) ([]Character, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var queryArgs = make([]interface{}, 0)
	// String を効率的に作り上げる strings.Builder を変数として定義
	var queryBuilder strings.Builder

	// クエリのベースを作る
	queryBuilder.WriteString(`SELECT
	characterId,
	UserID,
	characterType,
	circleColorName,
	savedFileName,
	blobImage,
	isPicUsed,
	nickName,
	firstName,
	middleName,
	lastName,
	originalName,
	isImaginary,
	birthDate,
	description,
	level,
	isLiked,
	registeredDate,
	updatedDate,
	onDeleteLock,
	inLimbo,
	reserveNumber01,
	reserveNumber02,
	reserveNumber03,
	reserveNumber04,
	reserveNumber05,
	reserveNumber06,
	reserveNumber07,
	reserveNumber08,
	reserveNumber09,
	reserveNumber10,
	reserveNumber11
	FROM characters WHERE `)

	// クエリのフィルターを作る
	if characterFilter.UserIDFilter != "" {
		queryBuilder.WriteString(`UserID LIKE ? `)
		// フィルターを % % で囲うことで、渡されるデータのうち、マッチするすべてを追加する
		queryArgs = append(queryArgs, "%"+strings.ToLower(characterFilter.UserIDFilter)+"%")
	}
	if characterFilter.CharacterTypeFilter != "" {
		if len(queryArgs) > 0 {
			queryBuilder.WriteString(" AND ")
		}
		queryBuilder.WriteString(`characterType LIKE ? `)
		queryArgs = append(queryArgs, "%"+strings.ToLower(characterFilter.CharacterTypeFilter)+"%")
	}
	if characterFilter.CircleColorNameFilter != "" {
		if len(queryArgs) > 0 {
			queryBuilder.WriteString(" AND ")
		}
		queryBuilder.WriteString(`circleColorName LIKE ? `)
		queryArgs = append(queryArgs, "%"+strings.ToLower(characterFilter.CircleColorNameFilter)+"%")
	}
	// if characterFilter.IsPicUsedFilter != "" {
	// 	if len(queryArgs) > 0 {
	// 		queryBuilder.WriteString(" AND ")
	// 	}
	// 	queryBuilder.WriteString(`sku LIKE ? `)
	// 	queryArgs = append(queryArgs, "characterFilter.IsPicUsedFilter")
	// }
	if characterFilter.NickNameFilter != "" {
		if len(queryArgs) > 0 {
			queryBuilder.WriteString(" AND ")
		}
		queryBuilder.WriteString(`nickName LIKE ? `)
		queryArgs = append(queryArgs, "%"+strings.ToLower(characterFilter.NickNameFilter)+"%")
	}
	if characterFilter.FirstNameFilter != "" {
		if len(queryArgs) > 0 {
			queryBuilder.WriteString(" AND ")
		}
		queryBuilder.WriteString(`firstName LIKE ? `)
		queryArgs = append(queryArgs, "%"+strings.ToLower(characterFilter.FirstNameFilter)+"%")
	}
	if characterFilter.MiddleNameFilter != "" {
		if len(queryArgs) > 0 {
			queryBuilder.WriteString(" AND ")
		}
		queryBuilder.WriteString(`middleName LIKE ? `)
		queryArgs = append(queryArgs, "%"+strings.ToLower(characterFilter.MiddleNameFilter)+"%")
	}
	if characterFilter.LastNameFilter != "" {
		if len(queryArgs) > 0 {
			queryBuilder.WriteString(" AND ")
		}
		queryBuilder.WriteString(`lastName LIKE ? `)
		queryArgs = append(queryArgs, "%"+strings.ToLower(characterFilter.LastNameFilter)+"%")
	}
	if characterFilter.OriginalNameFilter != "" {
		if len(queryArgs) > 0 {
			queryBuilder.WriteString(" AND ")
		}
		queryBuilder.WriteString(`originalName LIKE ? `)
		queryArgs = append(queryArgs, "%"+strings.ToLower(characterFilter.OriginalNameFilter)+"%")
	}
	// if characterFilter.IsImaginaryFilter != "" {
	// 	if len(queryArgs) > 0 {
	// 		queryBuilder.WriteString(" AND ")
	// 	}
	// 	queryBuilder.WriteString(`sku LIKE ? `)
	// 	queryArgs = append(queryArgs, "%"+strings.ToLower(characterFilter.IsImaginaryFilter)+"%")
	// }
	// if characterFilter.BirthDateFilter != "" {
	// 	if len(queryArgs) > 0 {
	// 		queryBuilder.WriteString(" AND ")
	// 	}
	// 	queryBuilder.WriteString(`sku LIKE ? `)
	// 	queryArgs = append(queryArgs, "%"+strings.ToLower(characterFilter.BirthDateFilter)+"%")
	// }
	// if characterFilter.LevelFilter != "" {
	// 	if len(queryArgs) > 0 {
	// 		queryBuilder.WriteString(" AND ")
	// 	}
	// 	queryBuilder.WriteString(`sku LIKE ? `)
	// 	queryArgs = append(queryArgs, "%"+strings.ToLower(characterFilter.LevelFilter)+"%")
	// }
	// if characterFilter.IsLikedFilter != "" {
	// 	if len(queryArgs) > 0 {
	// 		queryBuilder.WriteString(" AND ")
	// 	}
	// 	queryBuilder.WriteString(`sku LIKE ? `)
	// 	queryArgs = append(queryArgs, "%"+strings.ToLower(characterFilter.IsLikedFilter)+"%")
	// }
	// if characterFilter.RegisteredDateFilter != "" {
	// 	if len(queryArgs) > 0 {
	// 		queryBuilder.WriteString(" AND ")
	// 	}
	// 	queryBuilder.WriteString(`sku LIKE ? `)
	// 	queryArgs = append(queryArgs, "%"+strings.ToLower(characterFilter.RegisteredDateFilter)+"%")
	// }
	// if characterFilter.UpdatedDateFilter != "" {
	// 	if len(queryArgs) > 0 {
	// 		queryBuilder.WriteString(" AND ")
	// 	}
	// 	queryBuilder.WriteString(`sku LIKE ? `)
	// 	queryArgs = append(queryArgs, "%"+strings.ToLower(characterFilter.UpdatedDateFilter)+"%")
	// }
	// if characterFilter.OnDeleteLockFilter != "" {
	// 	if len(queryArgs) > 0 {
	// 		queryBuilder.WriteString(" AND ")
	// 	}
	// 	queryBuilder.WriteString(`sku LIKE ? `)
	// 	queryArgs = append(queryArgs, "%"+strings.ToLower(characterFilter.OnDeleteLockFilter)+"%")
	// }
	// if characterFilter.InLimboFilter != "" {
	// 	if len(queryArgs) > 0 {
	// 		queryBuilder.WriteString(" AND ")
	// 	}
	// 	queryBuilder.WriteString(`sku LIKE ? `)
	// 	queryArgs = append(queryArgs, "%"+strings.ToLower(characterFilter.InLimboFilter)+"%")
	// }

	// クエリを実行する
	results, err := database.DbConn.QueryContext(ctx, queryBuilder.String(), queryArgs...)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer results.Close()
	characters := make([]Character, 0)
	for results.Next() {
		var character Character
		results.Scan(&character.CharacterID,
			&character.UserID,
			&character.CharacterType,
			&character.CircleColorName,
			&character.SavedFileName,
			&character.BlobImage,
			&character.IsPicUsed,
			&character.NickName,
			&character.FirstName,
			&character.MiddleName,
			&character.LastName,
			&character.OriginalName,
			&character.IsImaginary,
			&character.BirthDate,
			&character.Description,
			&character.Level,
			&character.IsLiked,
			&character.RegisteredDate,
			&character.UpdatedDate,
			&character.OnDeleteLock,
			&character.InLimbo,
			&character.ReserveNumber01,
			&character.ReserveNumber02,
			&character.ReserveNumber03,
			&character.ReserveNumber04,
			&character.ReserveNumber05,
			&character.ReserveNumber06,
			&character.ReserveNumber07,
			&character.ReserveNumber08,
			&character.ReserveNumber09,
			&character.ReserveNumber10,
			&character.ReserveNumber11,
		)

		characters = append(characters, character)
	}
	return characters, nil
}
