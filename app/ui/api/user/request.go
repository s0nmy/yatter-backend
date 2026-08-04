package user

// （研修用の説明）
// ui/api/user/request.go
// ユーザーに関するエンドポイントのリクエストを定義
// XxxRequest のような命名にする

// Domain層のバリデーションとは別に、リクエストのバリデーションがあればここで行う
// ex. 必須パラメータ、任意パラメータ、ボディにあるべき、クエリパラメータにあるべき、...

// PostUsersRequest: ユーザー新規登録のリクエスト
type PostUsersRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
