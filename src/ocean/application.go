package main

type application struct {
	Id            string `json:"id"`
}

type AppData struct {
	PageNum   int    `json:"pageNum"`
	PageSize int `json:"pageSize"`
	Size   int    `json:"size"`
	OrderBy string `json:"orderBy"`
	StartRow   int    `json:"startRow"`
	EndRow int `json:"endRow"`
	Total   int    `json:"total"`
	Pages int `json:"pages"`
	List []ListData `json:"list"`
}

type ListData struct {
	Id   string    `json:"id"`
	ApplicationName   string    `json:"applicationName"`
	ApplicationType   int    `json:"applicationType"`
	ApplicationLogo   string    `json:"applicationLogo"`
	ApplicationSummary   string    `json:"applicationSummary"`
	ApplicationWebsite   string    `json:"applicationWebsite"`
	AuthorizationType   int    `json:"authorizationType"`
	ApplicationId   string    `json:"applicationId"`
	ApplicationKey   string    `json:"applicationKey"`
	Enabled   int    `json:"enabled"`
	CreateTime   string    `json:"createTime"`
	UserId   int    `json:"userId"`
	UserType   int    `json:"userType"`
	CallbackAddress   string    `json:"callbackAddress"`
	Remark   string    `json:"remark"`
	Approver   string    `json:"approver"`
	CompanyName   string    `json:"companyName"`
	ContactPhone   string    `json:"contactPhone"`
	ContactEmail   string    `json:"contactEmail"`
	BusinessLicenceNo   string    `json:"businessLicenceNo"`
	BusinessLicencePic   string    `json:"businessLicencePic"`
	LinkMan   string    `json:"linkMan"`
	CertNo   string    `json:"certNo"`
	CertFrontPic   string    `json:"certFrontPic"`
	CertBackPic   string    `json:"certBackPic"`
}

type AppResponse struct {
	Message       string   `json:"message"`
	Data         AppData  `json:"data"`
	Status        uint8   `json:"status"`
	ErrorCode  string   `json:"errorCode"`
	ErrorMessage     string `json:"errorMessage"`
	Exception  string  `json:"exception"`
}

func main() {
	InitApi("http://172.16.2.51:2500/")
	var app application
	app.Id = "1"
	resp, err := GetAppInfo("manager/app/findAppInfo?id=1",&app)
	if err == nil {
		print(resp.Message)
	}
}
