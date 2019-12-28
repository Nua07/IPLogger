package main

import "github.com/labstack/echo"
import "strconv"
import "net/url"
import "net/http"

type ND struct {
	Logs []string
	SiteURL string
}

func main(){
	var DATA map[int]*ND
	DATA=make(map[int]*ND)

	e:=echo.New()

	e.GET("/newID",func(c echo.Context) error {
		ID,err:=strconv.Atoi(c.QueryParam("id"))
		URL:=c.QueryParam("url")

		if(err!=nil){
			return c.String(http.StatusOK,"id error")
		}

		if(URL==""||!IsValidUrl(URL)){
			return c.String(http.StatusOK,"url error")
		}
		if _,ok:=DATA[ID]; ok{
			return c.String(http.StatusOK,"ID is exists")
		}
		DATA[ID]=&ND{
			Logs:[]string{},
			SiteURL:URL,
		}
		return c.String(http.StatusOK,"success")
	})

	e.GET("/viewIP/:id",func(c echo.Context) error {
		ID,err:=strconv.Atoi(c.Param("id"))
                if(err!=nil){
                        return c.String(http.StatusOK,"ERROR")
                }

                if _,ok:=DATA[ID]; ok{
			if !ok{
                       		 return c.String(http.StatusOK,"ID is exists")
			}
                }

		results:=""		
		for i,v := range DATA[ID].Logs {
			results+=strconv.Itoa(i)+" "+v+"\n"
		}
		return c.String(http.StatusOK,results)
	})
	e.GET("/exec/:id", func(c echo.Context) error {
		ID,err:=strconv.Atoi(c.Param("id"))
		if(err!=nil){
			return c.String(http.StatusOK,"ERROR")
		}
		if _,ok:=DATA[ID]; ok{
			if(!ok){
				return c.String(http.StatusOK,"ID is not exists")
			}
		}
		DATA[ID].Logs = append(DATA[ID].Logs,c.RealIP())
		return c.Redirect(302,DATA[ID].SiteURL)
	})
	e.Start(":8080")
}

func IsValidUrl(str string) bool {
u, err := url.Parse(str)
return err == nil && u.Scheme != "" && u.Host != ""
}
