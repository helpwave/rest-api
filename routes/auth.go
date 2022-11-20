package routes

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"regexp"
	"rest-api/auth"
	"rest-api/logging"
	"rest-api/models"
	"rest-api/util"
	"time"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log, _ := logging.GetRequestLogger(ctx)
		authHeader := ctx.GetHeader("Authorization")

		regex := regexp.MustCompile(`^Bearer (\w+)$`)

		var err error
		var token string

		if len(authHeader) == 0 {
			err = errors.New("missing Authorization header")
		} else {
			matches := regex.FindStringSubmatch(authHeader)
			if len(matches) != 2 {
				err = errors.New("authorization header invalid")
			} else {
				token = matches[1]
			}
		}

		if err != nil {
			log.Warn().Err(err).Send()
			SendError(ctx, http.StatusUnauthorized, err)
			return
		}

		// TODO: validate token and set user
		ctx.Set("authToken", token)
	}
}

type LoginRequest struct {
	Email    string `binding:"required,email"`
	Password string `binding:"required"`
}

type TokenResponse struct {
	Token string    `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...."`
	Exp   time.Time `json:"exp"`
}

type UserResponse struct {
	models.UserBase
	Role string `json:"role" example:"user"`
}

func (userResponse UserResponse) toClaim() auth.UserOrOrgClaim {
	return auth.UserOrOrgClaim{
		ID:   userResponse.ID,
		Role: userResponse.Role,
	}
}

type LoginResponse struct {
	User         UserResponse        `json:"user"`
	Organization auth.UserOrOrgClaim `json:"organization"`
	RefreshToken TokenResponse       `json:"refreshToken"`
	AccessToken  TokenResponse       `json:"accessToken"`
}

func getGlobalRole(user *models.User) string {
	for _, role := range user.GlobalRoles {
		if role.Role == models.Admin {
			return "admin"
		}
	}
	return "user"
}

// Login godoc
// @Summary    log in using email and password - get refresh-token and initial access-token
// @Tags       auth
// @Produce    json
// @Param      credentials                     body        LoginRequest      true    "User Credentials"
// @Success    200                             {object}    LoginResponse
// @Failure    400                             {object}    HTTPErrorResponse
// @Failure    401                             {object}    HTTPErrorResponse
// @Failure    500                             {object}    HTTPErrorResponse
// @Router     /auth/login                     [post]
func Login(ctx *gin.Context) {
	log, logCtx := logging.GetRequestLogger(ctx)

	body := LoginRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		log.Warn().Err(err).Msg("validation failed")
		SendError(ctx, http.StatusBadRequest, err)
		return
	}

	user := models.User{
		UserBase: models.UserBase{
			Email: body.Email,
		},
	}

	db := models.GetDB(logCtx)
	tx := db.
		Preload("Organizations").
		Preload("GlobalRoles").
		Where(&user).
		First(&user)

	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// we return the same error message in case the email was not found and in case the password was wrong
			// thus an unauthorized third party is not able to check whether the email is registered with us or not
			SendError(ctx, http.StatusUnauthorized, errors.New("email or password invalid"))
		} else {
			HandleDBError(ctx, logCtx, err)
		}
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PwBcrypt), []byte(body.Password)); err != nil {
		log.
			Info().
			Err(err).
			Str("user_id", user.ID.String()).
			Msg("Invalid Password")
		SendError(ctx, http.StatusUnauthorized, errors.New("email or password invalid"))
		return
	}

	if len(user.Organizations) == 0 {
		log.
			Warn().
			Str("user_id", user.ID.String()).
			Msg("user does not belong to any organizations")
		SendError(ctx, http.StatusBadRequest, errors.New("user is not authorized for any organizations"))
		return
	} else if len(user.Organizations) > 1 {
		log.
			Warn().
			Str("user_id", user.ID.String()).
			Msg("user is in more than one Organization")
	}

	organizationClaim := auth.UserOrOrgClaim{
		ID:   user.Organizations[0].ID,
		Role: "user",
	}

	userResponse := UserResponse{
		UserBase: user.UserBase,
		Role:     getGlobalRole(&user),
	}

	userClaim := userResponse.toClaim()

	refreshTokenString, refreshTokenClaim, err := auth.IssueRefreshToken(logCtx, userClaim, organizationClaim)
	if err != nil {
		log.Error().Err(err).Str("refresh_claim", util.Formatted(refreshTokenClaim)).Msg("failed to generate refresh token")
		SendError(ctx, http.StatusInternalServerError, errors.New("error generating token"))
		return
	}

	accessTokenString, accessTokenClaim, err := auth.IssueAccessToken(logCtx, userClaim, organizationClaim, refreshTokenClaim.ID)

	if err != nil {
		log.Error().Err(err).Str("refresh_claim", util.Formatted(refreshTokenClaim)).Msg("failed to generate access token")
		SendError(ctx, http.StatusInternalServerError, errors.New("error generating token"))
		return
	}

	response := LoginResponse{
		User:         userResponse,
		Organization: organizationClaim,
		RefreshToken: TokenResponse{
			Token: refreshTokenString,
			Exp:   refreshTokenClaim.Exp,
		},
		AccessToken: TokenResponse{
			Token: accessTokenString,
			Exp:   accessTokenClaim.Exp,
		},
	}

	ctx.JSON(http.StatusOK, response)
}
