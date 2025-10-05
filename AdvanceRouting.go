package main

//Single Responsibilty Priciple this is. SOLID=> Philoshopy
import (
	"encoding/json"
	"fmt"
	"net/http" // as amra json format a respons pathabao frontend a
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Safari server, Rony")
}

type Product struct {
	ID          int
	Title       string
	Description string
	Price       float64
	ImgURL      string
}

var ProductList []Product

func init() { // product object would be at first
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "I love Orange",
		Price:       100,
		ImgURL:      "https://i.chaldn.com/_mpimage/komola-orange-imported-50-gm-1-kg?src=https%3A%2F%2Feggyolk.chaldal.com%2Fapi%2FPicture%2FRaw%3FpictureId%3D64292&q=best&v=1",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Mango",
		Description: "I love Mango",
		Price:       60,
		ImgURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQX_LYndiiyKxXWCxRmyp8hW5bCaTiM_PF45Q&s",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "I Love Banana",
		Price:       80,
		ImgURL:      "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAkGBxASEhUREBITFRMVGBYYFhcWGBcWFxYYFRUXGBUXGBUaHSggGB0lGxUTITEiJSkrLi4uFx8zODMtNygtLisBCgoKDg0OGxAQGy0mICYtLy0tLS4vLTUvLy0tMi0vLTctLS0tLS0uLS01NS8tLS0tLS0tLS8tLTUvLS0tLi0tLf/AABEIANoA5wMBIgACEQEDEQH/xAAcAAEAAQUBAQAAAAAAAAAAAAAABwMEBQYIAgH/xABEEAABAwICBgcDCQYFBQAAAAABAAIDBBEhMQUGEkFRYQcTInGBkaEyscEUI0JSYoLR4fAIM0NykrJTk6LC0iQ1g8Pj/8QAGgEBAAMBAQEAAAAAAAAAAAAAAAMEBQIBBv/EAC4RAAICAQMBBQgDAQEAAAAAAAABAgMRBBIhMQUTQVFhIjJxgZGx0fChweEjQv/aAAwDAQACEQMRAD8Ag5ERAEREAREQBERAEREAREQBERAEREAREQBERAEREAREQBERAEREAREQBERAEREAREQBERAEREARFuPRxqJLpSYi7o6eO3Wy2v3RsvgXkeAGJ3A+N4BhNW9W6uvk6qkiMjhYuOAawE2u55waM+Ztgpk1b6DqaMB+kJnSvwJjiOxGOIL/AG3d42VJuitG09HC2npY2xxtyAzJ3uc7NxPEqnWVobnj7lTt1OOhLGvJiafUfQ8Q2WUFOQMO0zrD4ufc38VqmunRXRVEbn0MYgqACWtafm5LD2CzJpOQI8lt8lcXZmwHhh8F6oJNtw+r63GRP4KotRPdlMl7tYOTXNINjgRmviyWs0jXVlS5li108xbbKxkcRbwWNWwiqEREAREQBERAEREAREQBERAEREAREQBERAEREAREQF5ofRktVPHTwt2pJXBrRjvzJtkALkncAV1rq/oWHR9KylgHZYMT9J7z7T3cyfLAZBRT+z3q3+90jI3K8UBI/wA148Nlt+bgpdqZszgfS6qaizHBJCOS2ramwzx3/ktfnnJP2twvgfxV3VybRxv5Y+PFY2ZpPZzvjxsORzHisyTyWorB8iHWO4DC43E3x8Peq+sWk/kNDUVLrhzWFsd/a6x9ms5e04HuGKv9HUTcOyLjDHxHHvUTdOusYlnZQRuuyDtS85XDAfdad295G5WNNXvlkjtnhcEVoiLWKoREQBERAEREAREQBERAEREAREQBERAEREAREQBemNJIABJJsAMSScgAvK2/om0UKnStM1wJaxxldy6oF7b8tsMHivG8LIOjNWtENoaGClbnGwB3OR13SHxe5y8V8wtbD9b1kq5+OfNYGtmGPxtbyWTdLLLVaLWS+eNhwJBHj+K8UXaNziTjzwyxGX5rxUbrb+/xsT+KyWiGXOIGFvE3w7/yUGM8Er4R81j0m2go5ap1i6Nh2Q7Jz3GzG4bi4gHkCuV6qd8j3SSOLnvc5znHNznG7ie8kqZf2hNMWbTUTSMbzSAZ4EsjuOH70qFVr6eG2JTm8sIiKc4CIiAIiIAiIgCIiAIiIAiIgCIiAIiIAiIgCIiAKYP2ctH3qKqpI/dxsjHfK7aPpEPNQ+p9/Z5htQ1Mn1p9n+iJh/3qO14iexXJIVa+5OPAcefxWFleSfax/W4q/qn38zv52+CxhjB3N88fULHk8suRWDy4XcBlYbsrnHLLcFn6Jh7AJyu62GI42WCpW9s23WzPcOC2WkYLg4ZAW7yAva17R5M5p6W64y6VqcSRG5sTb7urY1rgPvB58Vp66J090RaPqZJZg+pjklc55Ic1zNt5JJ2XNuRck22gob1z1KqtHSWlG3ET2Jmg7DuR+o77J4YXWrXbCXCZVlFrk1pF6kjc02cCDwIIPkV5UxyEREAREQBERAEREAREQBERAEREAREQBEW59H3R7PpMl+11NOw2dKRtXdgdhjLi5scTkPReNpLLBqNLTSSvEcTHPe7JrGlzjhfBoxOAK2+j6LNMyN2vkuwMLCSSNjjf7LnXHjZT5qzqzS6Ni6ujh2n/AE5XAGR/N0mGHBowCvp69w9uN9uQBFh3KlbrYx6E0amzmjSfR/pWC5kpHkDfGWy+NoyTbwUu/s/f9tnFsql9938KFbU6qDxdpyxIB7XdirvQVmslsACXbR2RYkltto2zNmjE8FDHWd4tuPmjt07eS3mvbDnz3lWLC7l5hXUpwOe/Cx4lWgItl4qqSoqUZd1jsL3c4XuRbtHzWw05t4WGXNu9a5RkdYRfMn3/AJrPscb23W+B+IUlfDOJmWEY3rV9ZtaYKdt5JGMYTsgusNo8v1hmsnpvSAihJJte9zwGZ9FEdRqJU6UkNXWTmBhIEMOztPEOB2nAkbDnDcQSCccrLqct8tmdqXV+PokeRW1bmsvwNtqdLRTgsqYmTMP0ZWteO8X96i3pE1IjgaayhB6i/wA5Hckwk5OaTiWE4Y4g2zBwl3Qmp9JBEIQ6aQNwaZH9oDgNkDDksidX6exYNqxBDg6zmuBwILSMQq9DvpnlSzEns7mcejTOTkUka+9FtTSufPSMM1Ne+y27pIwcwWZuaOIvYZ5XUbrejJSWUZ7WAiIujwIiIAiIgCIiAIiIAiIgCIq1FSvlkZFG3ae9wa0cS42A5IDN6jaqS6SqRTxnYaAXSSEXEbBvtvJJAAvvXTNCaWihjpoQNiJoa0C27Mni4m5J3klaNoKnh0TRGIOZtNBkqZWg9o9+ey0ENaN/C7lG2nukiqlcRT/Ms3GwdIe8nBu7AeayJ326mTjQuF4voWlXGtZs+hPtRrE03w/XmrV+lo35Gx57u4fmubG60V4O18qmvzcSPI4LatTtbKyoqIqV7BM6RwaHCzHDMlxtgQACTgMAVFbpNVFbm0/4JIWUt45RL1UQcTxwcLBw4d+SvdDTOJe1+J2QQeIBtluPaCrU2r7GNAdI9zjwsAPCxwV3S6FbG4ubI7EEG4BGNrY+CrV1S3qWPiSSnHbgxE78xcYE+uPxVlE8byL87D3rYpNBlxdZ4x3EG1x493ksfLq/IzHrGEDvHl2SrTTSyyJSRYteOsOV8CHEZYDLyKzsEpABcRe2JyGGPuCwFbHsOabi1gCbm2e++7FeqqtsNlpB5+FsFE7oxjvydbHJpGQq5zI4H2WtGBdmedlVp9jIOB8RdaNpvSr44pJGDac1rnAG9jsgm2HctBh6Tpwe3BGR9lzm++6iojbdmcI5+ZJZGNeFJ4Oh42DcV72Twuop1b6SYJSG7RjefoyWAPc/LwwPJSNo7SzJM8D+vNS95tlsmtrI3DjdF5RlA9R70gdFsNaXVNMRFUkXIsOrlcN7vquP1h4hb+eO73fkvbHKzXY4vKIXFM48rqOSGR0UzHMkYbOa4WIKoLo3pY1GZXRGoibaqjZ2SP4rW3PVuHHOx44b1zktOuxTWSCSwERFIchERAEREAREQBERAFKPRToINjNa5u1I9xjpwdwGD3jvN233BruKjGGIvcGNzcQB3k2C6T0TRthY0NFmQsEUf3QAT3n8VmdqX93WoL/19vEtaSvdPPkRt0v6U2OroGO4Szn67jcRtPIdo25tKjJZ3Xms62vqH3v84W/5YDP9qwSt6WtQpil5EV0t02wpg/Z60SDJUVbhixrYmH+btSe6MeJUPqeOgEWo5Txmf6RxBc6yW2v4tHlayyTXTjaN8QG5X4r1HU4C5F8/wWPc83ceTc+d18M2V8MNxy+CxVY8ltwMwyazcPesLrLXbDMDwt3uV0Zezv3YH8FrmuEuDf5h6NK51lj7vadaetOZjjVBwIO/fv8ANWV3Nabm5O/kMl4YcFUIuFkxbXBqbUjHuN8DjxUQac0caed8JvZp7J4tOLT5EeqmKSOy1zXnQRnhE8YvLCDcDN0eZ8W4nuLlrdmalVWbX0f3K2up317l1RGK3HU7XiWmc2OYufDgAcS6Pu4jl5cFpyL6G6mFsds0Y0LJQeYnVugNNMma0hwcHC7XA3BB3hZ23DL3Ll3UrW2SieGuJdA49puZaT9JvxG9dEavaYZMxpDg4OALXDEEFY8q5UT2T6Po/wB8S1lWLdH5ozg4HeudemfVn5LWdfG20NTd+GTZR+8b44O+8eC6HcLLVukzQXyygljaLyNHWRYY7bMbDm4bTfvKzTZslyQyjlHL6Ii1CAIiIAiIgCIiAIiIDM6mwh9dTNOXWsP9J2rein10p2AO4nvOJ9SVAGqM2xWQH7Yb4v7I9XKd5H4Du+C+Z7fb3wXoa3ZsU1I590w69RMTvkk/vKs1facZapnHCWQf6yrFfR1+4vgZc/eYU+dBbw2hdzfL/wCsfgoDU1dA1aOplivi2Qm3J7G29WO8lV12e7T9V+Dun3vkSZt9p27Act54Kg+SxGG63lhv7lcQEOe4ZWH47lZ1gts3vfn4/msNvHP71LqKgqLNGdzbcLfjuWH1uN2tP2h/aVeNJ438LKhp9u1C093pcLi/mJJTxNGuRK4iVBgVw0LPkjQyeJ4VRgJaVkALiyovp1zuwdxfgyL9ftVfk7vlMDf+nee0BlE47rbmndwy4X0xdCANILJGhzHCzmkXBBzBCjLXLUR9PeelBkgzLc3xDnvc0cd2/ivpOzu0lNKu18+D8/8Afv8AExtXpHB7odDSFufR5rg6kkEMrj8ncc/8JxPtD7J3jx430xFrW1Rtg4SKMJuDyjr3RVc2RoxF8+NxuKup2XFlA3RZrqY3NpJ3YXtC4nLhGTw4eXBTpS1YkaCLbW8dyxnGVcu7n18PVFvhrdHp9jmHpF0V8m0hUMAs1zusZw2ZO1hyBLm/dWtqTOniC1ZDIMnw7PLsPd/zUZrYolurTKk1iTQREUpyEREAREQBERAVIJixzXt9ppDh3g3HuU/aPqhLDHI3JzWkdxGHoufVJ3Rhp4OjNJIe025j5tOJHgT6jgsftrTuylTj1j9i/wBn27LNr8TT9eabq66YW9oh457bQfeSsCt66V6YCWGUfTa5p+4QR6P9FoqvaGzfp4S9PtwV9THbbJeoW/8AQzpTqa1zDlIy/eYze39LnrQFf6B0h8nqIpvqPBPNuTx/SSpNRXvrcV+vwI63iSZ1HE4tkJxs4YHLu96paQaSMv0QvDXbcLJWi/MZcL/FVJe0ML4jw4r52XKwX0W1jn3H8BzVOrjLoiDmD+BVeGNpGJxAtl3ryw9l47ivJLMMncXiRrMkdivUZV5Usa4B7CC05EcjYjvBBBHJWaq2VtFuueUV2q4jsc1aMcq7CqdkSY9ywKlGXNyV2x+4qp1YOSgUmuobwaRrFqHS1V5ISIJjnYfNuPNv0TzHkVGmntWaujPz8ZDdz29qM9zhl3GxU/upV86k2LSA5pwIOII4EHNa+l7Xtp9mXtL16/Uo3aSufMeGc1AqYej7XwvZ1UzvnmCwv/Fb9b+Yb+OfG3nWTozhmvJRkQSf4Zv1Tu7ez1HIKM9JaJq6KQddG+J4N2u3G29rxg7wK3I3afXQxF+0vqv8KO2dEvaXBIvTXVw1DYXxHadC5zXkW2SJA3K5vgWgHDC6iZVJZ3OJLnE3N1TV+qGyO0gskpSygiuKKilmcGQxvkecmsaXHyCyFJoV7KgR1cckbWG8lwRYDnzyuOK9lZGPDfPXHj9DlLJh0W+1FBQTt2IGMMjiGs2dppbxJG8AXzWkVdO6N743e0xxae9psuare8WcNfE9lFoooiKU5CIiAK70ZpGSnf1kRsbEHgQcx6DyVoi8aT4Z6ngzGsWsU1YWGQBrWCzWt4m204neTYcsFh0ReQhGEdsVhCUnJ5YREXR4dAdEGmRU0PUON5Iez91ttn/Ts+RW30zLgtObT6Fc99GusJoqxjiexIQ13DE9k+pH3iuiWzMuJWey7A8r/n71g6urZZ/Pyf4f9F6qWY/vUtns2XEccRzVKP27bnAhXley1nfVOfJW07MA5uNrFVfNE/XDIkm1kdo7SNRBKCaWR+2RmWGQBxe0d5NxvW9jYe0SRuDmOF2ubiCDkbqNemem2axkgyfGPNrifc5qw+p2uEtE7YcDJTuPaZvbfNzOB5ZH1WjLS9/TGceuPrggjf3djT6Eu7KqMeqdFWQ1EYmgeHxneMweDhuPIo8LBtqaeGasJKSLlsirskWJMhCqx1KryqydtGajm4qu1zSsRHUK4bMFA4tEcomTDQvT4mubsvaHNObXAOB7wcFjuu5qjLpAtyK4lByXBFKOEfKjUnRcuL6OH7odH/YQvMWpOio8W0cP3tp/95Ktp9MykGx8lj63SUgF5Hho4ucGj1KlhDWSW3vJfDLM62UI+BmKzSFPTNLY+rjaNzA1o9LBRzp/WP5W17LWaL2N+1hv5bvyTTul4niwqIONr7YPI2OS1OeolDi7skHC7LbJ34cV9B2b2ZGH/SfvepDCafLMlq23ZqIHh2IkaD3E2PoT5LBayVAkqpnjfI7LLA2v6K4bXuhscARiBvJtmd9lhiVvxXORZJYwj4iIuyEIiIAiIgCIiAIiIAp16JNaRURGnlN5GAB3Fwya8el/zCgpXmidJS00rZoXWe03HA8QeIKr6mjvY48V0/fUkrntfodWsODmOGVgD+XDEK1Y3AtONsDbBYPVDWtlfG2RptJk8E3IP1T5nHfdZqp3PGAOB4g8Fhyjh8+HBeRFvTdRWjp5fquczwc3/wCaiRTv0uMD9HuGbmFjx/UAfQuUELW7Of8Ayx5N/n+yrqV7efNGR0Hpuekk6yB9icHNOLXjg5u/3qUdAa70lTZsloZTucew4/Zf8DiodRS6jSV3L2uvmeU6idT4OhZIVbuhKh3QutdZS2bHJtMH0H9pngM2+BC3fRfSRTvsKiN0R+s3tt/5DyKxL+zbocx5Xp+DVq11cve4NtbdVBIrCm1goZBdlTFjuc4NP9LrFKrS1K0XdUQj/wAjfddZzonnDi/oWt8Gs5RfOqVbz1LQ0ue4NaBcuJsAOJK1PSuu9LGCIiZXcgWt8XH4ArRdNafqKo/Ous0ZMbg0eG88yr+n7LsnzLhfvgUtRq64rEeWbFrFry9xMdH2GZGQjtu/lB9gevctMmlc8lz3FzjmXEknvJXhF9BTp66Y4gjHlJt5YREUx4EREAREQBERAEREAREQBERAEREBktA6bno5RLA6xFrg+y4cHDf8FNernSXQzx/9Q5sUtu017g0HudgHe/koCRQW0Qs5fUkhY48Ey9I+t9H1ToYy2R0sbhaMh2xtDslzss9w4KGkRdVUqtYQsscwiIpSMIiIAiIgCIiAIiIAiIgCIiAIiIAiIgCIiAIiIAiIgCIiAIiIAiIgCIiAIiIAiIgCIiAIiIAiIgCIiAIiID//2Q==",
	}
	//prd1 := Product{ID: 1, Title: "Orange", Price: 100} we can defind also the object like this
	ProductList = append(ProductList, prd1)
	ProductList = append(ProductList, prd2)
	ProductList = append(ProductList, prd3)

}

func handleCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,GET,PUT,PATCH,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Habib")
	w.Header().Set("Content-Type", "Application/json")
}

func preflightRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(200)
		return
	}
}

func sendData(w http.ResponseWriter, data interface{}, statuscode int) {
	w.WriteHeader(statuscode)
	encode := json.NewEncoder(w)
	encode.Encode(ProductList)
}

func getProduct(w http.ResponseWriter, r *http.Request) {

	handleCors(w)          //This function will handle will any cors error.
	preflightRequest(w, r) // This function will handle any cors error.

	if r.Method != http.MethodGet {
		fmt.Println(w, "Please give a get request method.")
	}
	// Othewise product list ta json format a return kore dibo

	sendData(w, ProductList, 200)
}

func createProduct(w http.ResponseWriter, r *http.Request) {

	handleCors(w)          //This function will handle any cors error.
	preflightRequest(w, r) // This function will handle any cors error.

	if r.Method != http.MethodPost { // you can write it "POST"
		http.Error(w, "Please give a post request", 400)
		return
	}

	var newProduct Product // create a new Empty object by the Struct, for the client to view the details after give the information
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me a valid json", 400)
	}

	newProduct.ID = len(ProductList) + 1
	ProductList = append(ProductList, newProduct)

	sendData(w, newProduct, 201)

}

func main() {

	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(getProduct)) //Advance Routing
	mux.Handle("POST /create-products", http.HandlerFunc(createProduct))

	mux.Handle("GET /hey", http.HandlerFunc(helloHandler))

	fmt.Println("Server is running on port: 8000")

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		fmt.Print("Error starting the server", err)
	}
}

/* func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the server, Rony")
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /hello", http.HandlerFunc(helloHandler)) // This is how we can create a Routing.

	fmt.Println(" Server running on port :8000")

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
*/
