package googlelogin

import "github.com/golang-jwt/jwt/v5"

type UserInfoDto struct {
	Email          string `json:"email"`          // "sony79410@gmail.com"
	Email_verified bool   `json:"email_verified"` // true
	Azp            string `json:"azp"`            // "14099599407-n78q9cn8eslht1ksculubui6oujn4mav.apps.googleusercontent.com"
	Name           string `json:"name"`           // "楊文哲（Sonic）"
	Picture        string `json:"picture"`        // "https://lh3.googleusercontent.com/a/AGNmyxaln4yq6fOwgJblycxQ5wx9RyJ7DmipLKXkkbwN=s96-c"
	Given_name     string `json:"given_name"`     // "文哲"
	Family_name    string `json:"family_name"`    // "楊"

	jwt.RegisteredClaims
}

// aud				這個 ID 權杖適用的目標對象。必須是應用程式的其中一個 OAuth 2.0 用戶端 ID。
// exp				不得接受 ID 權杖的有效期限。以 Unix 時間 (整數秒) 表示。
// iat				ID 權杖的核發時間。以 Unix 時間 (整數秒) 表示。
// iss				回應核發者的核發者 ID。Google ID 權杖一律為 https://accounts.google.com 或 accounts.google.com。

// sub
// 使用者的專屬 ID，必須專屬於所有 Google 帳戶，且不得重複使用。
// Google 帳戶可在不同時間點使用多個電子郵件地址，但 sub 值一律不會變更。
// 在應用程式中使用 sub 做為使用者的專屬 ID 金鑰。長度上限為 255 個區分大小寫的 ASCII 字元。

// at_hash
// 存取權杖雜湊。
// 提供存取權杖已綁定識別權杖的驗證。
// 如果在伺服器流程中使用 access_token 值核發 ID 權杖，則一律會包含這個憑證附加資訊。
// 這項憑證聲明可做為替代機制，藉此防範跨網站偽造攻擊，但如果您按照步驟 1 和步驟 3 操作，則不需要驗證存取權杖。

// azp
// 授權簡報者的 client_id。
// 只有在要求 ID 權杖的方與 ID 權杖的目標對像不同時，才需要使用這項憑證附加資訊。
// 如果在 Google 中混合使用混合式應用程式，即網頁應用程式和 Android 應用程式的 OAuth 2.0 client_id 不同，但共用同一個 Google API 專案。

// email
// 使用者的電子郵件地址。
// 只有在要求中包含 email 範圍時，才需要提供此項目。
// 此憑證附加資訊的值可能與這個帳戶不同，且可能會隨時間改變，因此不應該將此值設為連結至使用者記錄的主要 ID。
// 您也無法根據 email 憑證附加資訊的網域識別 Google Workspace 或 Cloud 機構的使用者，請改用 hd 憑證附加資訊。

// email_verified	如果使用者的電子郵件地址已通過驗證，則為 True，否則為 False。
// family_name		使用者的姓氏或名字。您可以在有 name 聲明時提供。
// given_name		使用者的名字。您可以在有 name 聲明時提供。

// hd
// 與使用者 Google Workspace 或 Cloud 機構相關聯的網域。
// 只有在使用者屬於 Google Cloud 機構時才會提供。
// 將資源的存取權限制為僅限特定網域的成員存取時，您必須檢查這項憑證附加資訊。
// 缺少這項聲明，代表該帳戶不屬於 Google 代管的網域。

// locale			使用者的語言代碼，以 BCP 47 語言標記表示。如有 name 聲明，便可提供。

// name
// 以可顯示格式呈現的使用者全名。可在下列情況提供：
// 要求範圍包含「profile」字串
// 要求更新權杖時，系統會傳回 ID 權杖
// 如果有 name 聲明，您就可以使用這些憑證更新應用程式的使用者記錄。請注意，這個版權聲明不保證一定會顯示。

// nonce		您的應用程式在驗證要求中提供的 nonce 值。您必須確保該訊息僅出現一次，以強制執行重送攻擊。

// picture
// 使用者個人資料相片的網址。可在下列情況提供：
// 要求範圍包含「profile」字串
// 要求更新權杖時，系統會傳回 ID 權杖
// 如果有 picture 聲明，您就可以用這些憑證更新應用程式的使用者記錄。請注意，這個版權聲明不保證一定會顯示。

// profile
// 使用者的個人資料頁面網址。可在下列情況提供：
// 要求範圍包含「profile」字串
// 要求更新權杖時，系統會傳回 ID 權杖
// 如果有 profile 聲明，您就可以用這些憑證更新應用程式的使用者記錄。請注意，這個版權聲明不保證一定會顯示。
