package controller

import (
	"net/http"
	"strconv"

	"github.com/bagustyo92/auth/utils"
	"github.com/labstack/echo"
)

func (uc *cartController) dockerfile(c echo.Context) error {
	return c.JSON(utils.Response(http.StatusOK, "OK", map[string]string{
		"data": `1. WORKDIR/ADD command seharusnya berada pada /go/src/github.com/telkomdev/indihome/backend sesuai dengan folder yg ditambahkan kedalam dockernya`,
	}))
}

func (uc *cartController) microservice(c echo.Context) error {
	return c.JSON(utils.Response(http.StatusOK, "OK", map[string]string{
		"data": `1. Lebih cepat dalam hal time to market
		Sangat cocok diimplementasikan untuk melihat kondisi market
		microservice dapat plug and play, sehingga developer bisa dengan cepat jika memang harus ada adjustment berdasarkan kebutuhan market.
		Minim akan conflict, dan juga support filosofi cicd dimana bisa dengan cepat di deploy tanpa menganggu system secara kesuluruhan.
		2. Menghindari single point of failure dan meningkatkan availability
		Dengan konsep ini tentu untuk mengurangi single point of failure karena biasanya sebuah service akan di replica dan dilakukan load balancing
		sehingga pemrosesan di sebuah aplikasi dapat dibagi bebanya. dan tentu dengan adanya replica dapat membagi trafic serta tugas dari sebuah aplikasi
		ketika salah satu replica overload atau sedang melakukan proses lain sehingga tidak dapat melayani permintaan lainya.
		3. Fault isolation
		Harapanya dengan microservice failure tidak berimpak kepada service lain dan cepat menemukan sumber masalah jika memang kita mengimplementasikan
		tracing data dengan baik di antar service. Selain itu mengurangi dependency terhadap service lain atau part lain dan functional tiap service dapat 
		terlihat jelas
		4. Membagi pekerjaan dengan jelas
		Dengan adanya microservice pembagian pekerjaan sekaligus tim akan jauh lebih mudah karena beban tanggung jawab dapat dibagi ke beberapa tim ataupun orang.
		Tim atau developer pun dapat jauh lebih fokus dan jauh lebih cepat mentranslate kebutuhan tim product ataupun bisnis kedalam sebuah sistem. Handle atau identifikasi 
		error ataupun bug juga dapat lebih cepat ditangani.`,
	}))
}

func (uc *cartController) indexDB(c echo.Context) error {
	return c.JSON(utils.Response(http.StatusOK, "OK", map[string]string{
		"data": `simpelnya gambaran sebuah index di database adalah seperti buku telpon jadul ataupun seperti kamus.
		dimana adanya pengelompokan dan pengurutan untuk mempermudah sekaligus mempercepat proses pencarian. begitu pula index database,
		ketika kita membuat index di database yang dilakukan adalah mengelompokan sekaligus mengurutkan data, lebih detailnya data index disimpan
		dalam sebuah struktur data yang berisi kumpulan keys beserta referensinya atau alamat ke actual data table. selain itu data tersimpan secara 
		berurutan dan index "biasanya" disimpan kedalam sebuah RAM instead of hardisk. namun yang perlu diperhatikan adalah Jangan sampai karena semua 
		query ingin cepat, kita jadi membuat banyak index, yang jadinya bisa memakan storage berkali lipat dari ukuran table itu sendiri. Selain itu setiap 
		kita melakukan operasi CRUD, si mySQL harus melakukan pengapdetan terhadap key ke semua index yang sudah ada.`,
	}))
}

func (uc *cartController) moneyTest(c echo.Context) error {
	amountString := c.Param("amount")
	amount, err := strconv.Atoi(amountString)
	if err != nil {
		return c.JSON(utils.Response(http.StatusBadRequest, err, nil))
	}
	res := uc.cartController.MoneyTest(amount)
	return c.JSON(utils.Response(http.StatusOK, "OK", res))
}

func (uc *cartController) stringTest(c echo.Context) error {
	type strtest struct {
		Text1 string `json:"text1"`
		Text2 string `json:"text2"`
	}
	a := strtest{}
	if err := c.Bind(&a); err != nil {
		return c.JSON(utils.Response(http.StatusBadRequest, err, nil))
	}
	res := uc.cartController.StringTest(a.Text1, a.Text2)
	return c.JSON(utils.Response(http.StatusOK, "OK", res))
}
