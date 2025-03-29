package services

import (
	"business/api/generated/companies"
	models "business/models/generated"
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
)

type CompanyService interface {
	SignIn(ctx context.Context, requestParams companies.PostCompaniesSignInJSONRequestBody) (statusCode int64, tokenString string, error error)
}

type companyService struct {
	db *sql.DB
}

func NewCompanyService(db *sql.DB) CompanyService {
	return &companyService{db}
}

func (cs *companyService) SignIn(ctx context.Context, requestParams companies.PostCompaniesSignInJSONRequestBody) (statusCode int64, tokenString string, error error) {
	// NOTE: emailからの取得
	company, err := models.Companies(qm.Where("email = ?", requestParams.Email)).One(ctx, cs.db)
	if err != nil {
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
	tokenString, err = token.SignedString([]byte(os.Getenv("JWT_TOKEN_KEY")))
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
