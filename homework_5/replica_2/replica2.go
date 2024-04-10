package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"slices"
)

func main() {
	mux := http.NewServeMux()
	profile := Profile{userInfo: make(map[int]*User)}
	mux.HandleFunc("/create/", profile.Create)
	mux.HandleFunc("/get/", profile.getUsers)
	mux.HandleFunc("/", profile.Hello)
	mux.HandleFunc("/addfriends/", profile.addFriends)
	mux.HandleFunc("/deletefriends/", profile.deleteFriends)
	mux.HandleFunc("/friends/", profile.showFriends)
	mux.HandleFunc("/agechange/", profile.changeAge)

	http.ListenAndServe(":8082", mux)

}

// Напишите HTTP-сервис, который принимает входящие соединения с JSON-данными и обрабатывает их следующим образом: **
// 1. Сделайте обработчик создания пользователя. У пользователя должны быть следующие поля: имя, возраст и массив друзей.
// Пользователя необходимо сохранять в мапу. Пример запроса:
// POST /create HTTP/1.1
// Content-Type: application/json; charset=utf-8
// Host: localhost:8080
// {"name":"some name","age":"24","friends":[]}
// Данный запрос должен возвращать ID пользователя и статус 201.

type User struct {
	Id      int      `json:"id"`
	Name    string   `json: "name"`
	Age     int      `json: "age"`
	Friends []string `json: "friends"`
	Source  int      `json: "source"`
	Target  int      `json: "target"`
}

type Profile struct {
	userInfo map[int]*User // int ID
}

func (p Profile) Hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" { // non catch all
		http.NotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world ^^"))
}

func (p Profile) Create(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var u User

	if err := json.Unmarshal(content, &u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	p.userInfo[u.Id] = &u

	j, err := json.Marshal(p.userInfo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if err := os.WriteFile("rep2.json", j, 0644); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(201)
	w.Write([]byte("User created "))
	w.Write([]byte(fmt.Sprintf("%d", u.Id)))
	return

}

func (p Profile) getUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var result string
	for _, v := range p.userInfo {
		result += fmt.Sprintf("Id: %d Name: %s Age: %d, Friends: %s \n", v.Id, v.Name, v.Age, v.Friends)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
	return
}

// 2. Сделайте обработчик, который делает друзей из двух пользователей.
// Например, если мы создали двух пользователей и нам вернулись их ID,
// то в запросе мы можем указать ID пользователя, который инициировал запрос на дружбу,
// и ID пользователя, который примет инициатора в друзья. Пример запроса:
// POST /make_friends HTTP/1.1
// Content-Type: application/json; charset=utf-8
// Host: localhost:8080
// {"source_id":"1","target_id":"2"}
// Данный запрос должен возвращать статус 200 и сообщение «username_1 и username_2 теперь друзья».

// handler friends
// {"source": "1", target: "2"}

func (p Profile) addFriends(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	var u User

	if err := json.Unmarshal(content, &u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	p.userInfo[u.Source].Friends = append(p.userInfo[u.Source].Friends, p.userInfo[u.Target].Name)
	p.userInfo[u.Target].Friends = append(p.userInfo[u.Target].Friends, p.userInfo[u.Source].Name)

	j, err := json.Marshal(p.userInfo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if err := os.WriteFile("rep2.json", j, 0644); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v & %v are friends now ^.^/", u.Source, u.Target)))
	return
}

// 3. Сделайте обработчик, который удаляет пользователя.
// Данный обработчик принимает ID пользователя и удаляет его из хранилища,
// а также стирает его из массива friends у всех его друзей.
// Пример запроса:
// DELETE /user HTTP/1.1
// Content-Type: application/json; charset=utf-8
// Host: localhost:8080
// {"target_id":"1"}
// Данный запрос должен возвращать 200 и имя удалённого пользователя.

func (p Profile) deleteFriends(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.Header().Set("Allow", http.MethodDelete)
		http.Error(w, "Method not allowed", http.StatusInternalServerError)
		return
	}

	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	var u User

	if err := json.Unmarshal(content, &u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	_, ok := p.userInfo[u.Target]
	if ok {
		for _, v := range p.userInfo {
			if slices.Contains(v.Friends, p.userInfo[u.Target].Name) {
				v.Friends = slices.DeleteFunc(v.Friends, func(s string) bool {
					return s == p.userInfo[u.Target].Name
				})
			}
		}
		targetName := p.userInfo[u.Target].Name
		delete(p.userInfo, u.Target)
		j, err := json.Marshal(p.userInfo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		if err := os.WriteFile("rep2.json", j, 0644); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(targetName))
		return
	} else {
		http.Error(w, "ID not found", http.StatusInternalServerError)
		return
	}

}

// 4. Сделайте обработчик, который возвращает всех друзей пользователя. Пример запроса:
// GET /friends/user_id HTTP/1.1
// Host: localhost:8080
// Connection: close
// После /friends/ указывается id пользователя, друзей которого мы хотим увидеть.

func (p Profile) showFriends(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method not allowded", http.StatusInternalServerError)
		return
	}
	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	var u User

	if err := json.Unmarshal(content, &u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("User ID: %d friends are: %s", p.userInfo[u.Target].Id, p.userInfo[u.Target].Friends)))
	return
}

// 5. Сделайте обработчик, который обновляет возраст пользователя. Пример запроса:
// PUT /user_id HTTP/1.1
// Content-Type: application/json; charset=utf-8
// Host: localhost:8080
// {"new age":"28"}
// Запрос должен возвращать 200 и сообщение «возраст пользователя успешно обновлён».

func (p Profile) changeAge(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.Header().Set("Allow", http.MethodPut)
		http.Error(w, "Method not allowed", http.StatusInternalServerError)
		return
	}

	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	defer r.Body.Close()

	var u User

	if err := json.Unmarshal(content, &u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	p.userInfo[u.Target].Age = u.Age

	j, err := json.Marshal(p.userInfo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if err := os.WriteFile("rep2.json", j, 0644); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Age of user %d updated", p.userInfo[u.Target].Id)))
	return

}
