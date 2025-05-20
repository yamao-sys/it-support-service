package businessservices

import (
	businessapi "apps/api/business"
	models "apps/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SupporterService interface {
	SignIn(requestParams businessapi.PostSupporterSignInJSONRequestBody) (statusCode int, tokenString string, error error)
}

type supporterService struct {
	db *gorm.DB
}

func NewSupporterService(db *gorm.DB) SupporterService {
	return &supporterService{db}
}

func (ss *supporterService) SignIn(requestParams businessapi.PostSupporterSignInJSONRequestBody) (statusCode int, tokenString string, error error) {
	// NOTE: emailからの取得
	var supporter models.Supporter
	ss.db.Where("email = ?", requestParams.Email).Take(&supporter)
	if supporter.ID == 0 {
		return http.StatusBadRequest, "", fmt.Errorf("メールアドレスまたはパスワードに該当する%sが存在しません。", "サポータ")
	}

	// NOTE: パスワードの照合
	if err := ss.compareHashPassword(supporter.Password, requestParams.Password); err != nil {
		return http.StatusBadRequest, "", fmt.Errorf("メールアドレスまたはパスワードに該当する%sが存在しません。", "サポータ")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"supporter_id": supporter.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_TOKEN_KEY")))
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	return http.StatusOK, tokenString, nil
}

// NOTE: パスワードの照合
func (ss *supporterService) compareHashPassword(hashedPassword, requestPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(requestPassword)); err != nil {
		return err
	}
	return nil
}
