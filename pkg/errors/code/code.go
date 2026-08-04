package code

type StatusCode string

// NOTE: StatusCode一覧、必要があれば随時追加する.
var (
	BadRequest   StatusCode = "BAD_REQUEST"  // 400
	Unauthorized StatusCode = "UNAUTHORIZED" // 401
	Forbidden    StatusCode = "FORBIDDEN"    // 403
	NotFound     StatusCode = "NOT_FOUND"    // 404
	Conflict     StatusCode = "CONFLICT"     // 409
	Internal     StatusCode = "INTERNAL"     // 500
)

/*
研修資料
-----

ステータスコードはHTTPの知識である.
Usecase層やDomain層などからステータスコードを呼び出してしまうと、これらがUI層に依存することになる.

仮に、UIがRESTからGraphQLに変更されたらどうだろうか.
GraphQLはいかなる時も200を返却し、エラーコードなどはextensionsに持たせる.

UsecaseやDomain層などからHTTPのステータスコードを呼んでいる場合、全てのステータスコードをGraphQL用のエラーコードに置き換える必要がある.
だが、ステータスコードをあらかじめ自前で用意ておくことによってGraphQL用エラーコードへの変換処理のみを変更すれば対応が可能になる.
このように、依存関係を適切にすることによって変更容易性が向上する.
*/
