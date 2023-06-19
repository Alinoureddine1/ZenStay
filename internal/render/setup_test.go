package render

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/Alinoureddine1/ZenStay/internal/config"
	"github.com/alexedwards/scs/v2"
)

var session *scs.SessionManager
var testAppConfig config.AppConfig

func TestMain(m *testing.M) {
	testAppConfig.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false
	testAppConfig.Session = session
	app = &testAppConfig

	os.Exit(m.Run())
}