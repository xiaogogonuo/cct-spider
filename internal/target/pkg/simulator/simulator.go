package simulator

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"github.com/tebeka/selenium/firefox"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"strings"
	"sync"
)

// Simulator 模拟浏览器
type Simulator struct {
	driverPath string // 浏览器驱动路径
	port       int    // 浏览器驱动端口
	caps       selenium.Capabilities
}

func (s *Simulator) AddChromeCap(cap chrome.Capabilities) {
	s.caps.AddChrome(cap)
}

func (s *Simulator) NewChromeService(opts ...selenium.ServiceOption) (*selenium.Service, error) {
	return selenium.NewChromeDriverService(s.driverPath, s.port, opts...)
}

func (s *Simulator) AddFirefoxCap(cap firefox.Capabilities) {
	s.caps.AddFirefox(cap)
}

func (s *Simulator) NewFirefoxService(opts []selenium.ServiceOption) (*selenium.Service, error) {
	return selenium.NewGeckoDriverService(s.driverPath, s.port, opts...)
}

func (s *Simulator) NewWebDriver() (selenium.WebDriver, error) {
	return selenium.NewRemote(s.caps, fmt.Sprintf("http://localhost:%d/wd/hub", s.port))
}

// NewSimulator 创建模拟器
func NewSimulator(browserName, driverPath string, port int) *Simulator {
	return &Simulator{
		driverPath: driverPath,
		port:       port,
		caps:       selenium.Capabilities{"browserName": strings.ToLower(browserName)},
	}
}

var once sync.Once
var DriverService *selenium.Service
var WebDriver selenium.WebDriver

func InitPBCSimulator() {
	once.Do(PBCSimulator)
}

func PBCSimulator() {
	sim := NewSimulator("chrome", "./chromedriver", 9999)
	service, err := selenium.NewChromeDriverService(sim.driverPath, sim.port)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	DriverService = service
	//禁止图片加载，加快渲染速度
	imagCaps := map[string]interface{}{
		"profile.managed_default_content_settings.images": 2,
	}
	chromeCaps := chrome.Capabilities{
		Prefs: imagCaps,
		Path:  "",
		Args: []string{
			// 设置Chrome无头模式，在linux下运行，需要设置这个参数，否则会报错
			"--headless",

			//"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36", // 模拟user-agent，防反爬
		},
	}
	sim.AddChromeCap(chromeCaps)

	wb, err := sim.NewWebDriver()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	WebDriver= wb
}
