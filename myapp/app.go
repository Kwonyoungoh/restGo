package myapp

import "net/http"

type User struct{
	FirstName string	`json:"first_name"`
	LastName string		`json:"last_name"`
	Email string		`json:"email"`
	CreatedAt time.Time	`json:"created_at"`
}

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w  http.ResponseWriter, r *http.Request){ // 포인터로 받아야 원래의 것을 바꿔준다
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err !=nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w,"Bad Request: ", err )
		return
	}
	user.CreatedAt = time.Now()

	data, _ := json.Marshal(user)
	w.Header().Add("content-type","application/json")//header에 json이라고 명시 해준다.
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func barHandler(w http.ResponseWriter, r *http.Request){
	name := r.URL.Query().Get("name") // ㅕurl 정보 쿼리해서 name 이라는 argument를 get하면 아래와같이 response 해주겠다
	if name ==""{
		name = "World Test"
	}
	fmt.Fprintf(w,"hello %s!", name)
}

func NewHTTPHandler() http.Handler {
		mux := http.NewServeMux()//라우터를 등록 mux라는 인스턴스에 라우터를 등록해서 그 인스턴스를 넘겨주는 방식

	// 요청에 따라 무얼 할지 헨들러를 등록하는 곳
	mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){ 
		fmt.Fprint(w, "hello worldddddd")
	})

	mux.HandleFunc("/bar", barHandler)

	mux.Handle("/foo", &fooHandler{}) // & 주소로넘긴다
	return mux
}