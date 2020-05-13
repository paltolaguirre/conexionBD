package automigrateFunction

import (
	"crypto/tls"
	"fmt"
	"github.com/xubiosueldos/framework/configuracion"
	"io/ioutil"
	"net/http"
	"strings"
)

func CrearFunctionEnMicroservicioFormulas(funcion string) error {

	config := configuracion.GetInstance()
	url := configuracion.GetUrlMicroservicio(config.Puertomicroservicioformula) + "formula/public/formulas"
	token := configuracion.GetInstance().Tokenformulapublic

	req, err := http.NewRequest("POST", url, strings.NewReader(funcion))

	if err != nil {
		fmt.Println("Error: ", err)
	}

	req.Header.Add("Authorization", "Bearer " + token)

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("URL:", url)

	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error: ", err)
	}


	return err
}
