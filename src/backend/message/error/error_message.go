package error

const (
	MSG_E00001 = "環境変数の取得に失敗しました。KEY=[%s]"
	MSG_E00002 = "環境変数の変換に失敗しました。KEY=[%s], 変換値=[%s]"
	MSG_E00003 = "SQLの実行に失敗しました。テーブル名=[%s],エラー[%s]"
	MSG_E00004 = "存在チェックに失敗しました。テーブル名=[%s],検索条件=[%s]"
	MSG_E00005 = "対象レコードが存在しません。テーブル名=[%s],検索条件=[%s]"
	MSG_E00006 = "登録内容が重複しています。テーブル名=[%s],KEY=[%s]"
	MSG_E00007 = "DBの接続に失敗しました。DB=[%s],ADDRESS=[%s],PORT=[%s]"
	MSG_E00008 = "サーバー起動に失敗しました。"

	MSG_W00001 = "環境変数の取得に失敗しました。デフォルト値を使用します。KEY=[%s],デフォルト値=[%s]"
	MSG_W00002 = "環境変数の変換に失敗しました。デフォルト値を使用します。KEY=[%s],変換値=[%s],デフォルト値=[%v]"
)
