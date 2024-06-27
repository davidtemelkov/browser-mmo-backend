package utils

import (
	"browser-mmo-backend/constants"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type FireBaseStorage struct {
	Bucket string
}

// Folder constants
var (
	PlayerFolder    = "player_images"
	WeaponFolder    = "weapon_images"
	ArmourFolder    = "armour_images"
	AccessoryFolder = "accessory_images"
	QuestFolder     = "quest_images"
	MonsterFolder   = "monster_images"
	BossFolder      = "boss_images"
	MountFolder     = "mount_images"
)
var expectedPrefixes = map[string]string{
	"image/jpeg;base64,": "image/jpeg",
	"image/png;base64,":  "image/png",
}

func NewFireBaseStorage(bucket string) *FireBaseStorage {
	return &FireBaseStorage{
		Bucket: bucket,
	}
}

func ValidateAndExtractContentType(photo64 string) (string, string, error) {
	for prefix, contentType := range expectedPrefixes {
		if strings.HasPrefix(photo64, prefix) {
			return contentType, prefix, nil
		}
	}

	return "", "", errors.New(constants.InvalidBase64ImagePrefixError)
}

func UploadFile(photo64, fileFolder, fileName string) (string, error) {
	bucketName := GetFirebaseBucketName()

	fb := NewFireBaseStorage(bucketName)
	ctx := context.Background()
	opt := option.WithCredentialsFile("helpers/serviceAccountKey.json")

	client, err := storage.NewClient(ctx, opt)
	if err != nil {
		panic(constants.FirebaseClientError)
	}

	contentType, prefix, err := ValidateAndExtractContentType(photo64)
	if err != nil {
		return "", err
	}

	photo64 = strings.TrimPrefix(photo64, prefix)

	photoData, err := base64.StdEncoding.DecodeString(photo64)
	if err != nil {
		return "", err
	}

	fileName = strings.ReplaceAll(fileName, " ", "")
	filePath := fmt.Sprintf("%s/%s", fileFolder, fileName)
	wc := client.Bucket(fb.Bucket).Object(filePath).NewWriter(ctx)

	wc.ContentType = contentType

	if _, err := wc.Write(photoData); err != nil {
		return "", err
	}

	if err := wc.Close(); err != nil {
		return "", err
	}

	url, err := generateFirebaseUrl(fileFolder, fileName)
	if err != nil {
		return "", err
	}

	return url, nil
}

func generateFirebaseUrl(fileFolder, fileName string) (string, error) {
	baseUrl := GetFirebaseUrl()

	if fileFolder == "" {
		return "", errors.New(constants.FileFolderEmptyError)
	}
	if fileName == "" {
		return "", errors.New(constants.FileNameEmptyError)
	}

	url := fmt.Sprintf("%s%s%%2F%s?alt=media", baseUrl, fileFolder, fileName)

	return url, nil
}

func GetJWTPrivateKey() []byte {
	jwtPrivateKey := os.Getenv("JWT_PRIVATE_KEY")

	if jwtPrivateKey == "" {
		panic(constants.JWTPrivateKeyError)
	}

	return []byte(jwtPrivateKey)
}

func GetFirebaseUrl() string {
	firebaseUrl := os.Getenv("FIREBASE_URL")

	if firebaseUrl == "" {
		panic(constants.FirebaseURLKeyError)
	}

	return firebaseUrl
}

func GetFirebaseBucketName() string {
	firebaseBucketName := os.Getenv("FIREBASE_BUCKET_NAME")

	if firebaseBucketName == "" {
		panic(constants.FirebaseBucketNameError)
	}

	return firebaseBucketName
}
