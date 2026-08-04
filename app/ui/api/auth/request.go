package auth

// （研修用の説明）
// ui/api/auth/request.go
// 認証に関するエンドポイントのリクエストを定義
// XxxRequest のような命名にする

// Domain層のバリデーションとは別に、リクエストのバリデーションがあればここで行う
// ex. 必須パラメータ、任意パラメータ、ボディにあるべき、クエリパラメータにあるべき、...

type PostLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
