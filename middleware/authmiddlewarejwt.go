package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"project1/config"
	"project1/model/respErr"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// mengambil token dari header Authorization
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, &respErr.ErrorResponse{
				Message: "Unauthorized",
				Status:  http.StatusUnauthorized,
			})
			return
		}

		// split token dari header
		tokenString := authHeader[len("Bearer "):]

		// parsing token dengan secret key
		claims := &config.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return config.JwtKey, nil
		})

		if err != nil {
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, &respErr.ErrorResponse{
					Message: "Unauthorized",
					Status:  http.StatusUnauthorized,
				})
				return
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, &respErr.ErrorResponse{
				Message: "invalid or expired token",
				Status:  http.StatusBadRequest,
			})
			return
		}

		if !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, &respErr.ErrorResponse{
				Message: "Unauthorized (non Valid)",
				Status:  http.StatusUnauthorized,
			})
			return
		}

		// Set data pengguna dari token ke dalam konteks
		ctx.Set("username", claims.Username)
		ctx.Set("user_id", claims.UserID)
		ctx.Set("role", claims.Role) // Menambahkan data peran ke konteks

		// Pengecekan peran admin
		if claims.Role != "admin" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, respErr.ErrorResponse{
				Message: "Unauthorized: Only admin can access this endpoint",
				Status:  http.StatusUnauthorized,
			})
			return
		}

		// Melanjutkan ke handler jika semua pengecekan berhasil
		ctx.Next()
	}
}

func UserMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// mengambil token dari header Authorization
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, &respErr.ErrorResponse{
				Message: "Unauthorized",
				Status:  http.StatusUnauthorized,
			})
			return
		}

		// split token dari header
		tokenString := authHeader[len("Bearer "):]

		// parsing token dengan secret key
		claims := &config.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return config.JwtKey, nil
		})

		if err != nil {
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, &respErr.ErrorResponse{
					Message: "Unauthorized",
					Status:  http.StatusUnauthorized,
				})
				return
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, &respErr.ErrorResponse{
				Message: "invalid or expired token",
				Status:  http.StatusBadRequest,
			})
			return
		}

		if !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, &respErr.ErrorResponse{
				Message: "Unauthorized (non Valid)",
				Status:  http.StatusUnauthorized,
			})
			return
		}

		// Set data pengguna dari token ke dalam konteks
		ctx.Set("username", claims.Username)
		ctx.Set("user_id", claims.UserID)
		ctx.Set("role", claims.Role)

		logrus.Info(claims.UserID)

		// Melanjutkan ke handler jika semua pengecekan berhasil
		ctx.Next()
	}
}

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logrus.Error("Panic occurred:", r)
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, respErr.ErrorResponse{
					Message: "Internal Server Error",
					Status:  http.StatusInternalServerError,
				})
			}
		}()

		ctx.Next()
	}
}
