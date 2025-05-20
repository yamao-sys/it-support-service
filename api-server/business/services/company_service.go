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

type CompanyService interface {
	SignIn(requestParams businessapi.PostCompanySignInJSONRequestBody) (statusCode int, tokenString string, error error)
}

type companyService struct {
	db *gorm.DB
}

func NewCompanyService(db *gorm.DB) CompanyService {
	return &companyService{db}
}

func (cs *companyService) SignIn(requestParams businessapi.PostCompanySignInJSONRequestBody) (statusCode int, tokenString string, error error) {
	var company models.Company

	// NOTE: emailからの取得
	cs.db.Where("email = ?", requestParams.Email).Take(&company)
	if company.ID == 0 {
		return http.StatusBadRequest, "", fmt.Errorf("メールアドレスまたはパスワードに該当する%sが存在しません。", "企業")
	}

	// NOTE: パスワードの照合
	if err := cs.compareHashPassword(company.Password, requestParams.Password); err != nil {
		return http.StatusBadRequest, "", fmt.Errorf("メールアドレスまたはパスワードに該当する%sが存在しません。", "企業")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"company_id": company.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_TOKEN_KEY")))
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	return http.StatusOK, tokenString, nil
}

// NOTE: パスワードの照合
func (cs *companyService) compareHashPassword(hashedPassword, requestPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(requestPassword)); err != nil {
		return err
	}
	return nil
}
