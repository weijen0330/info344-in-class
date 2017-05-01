package handlers

import (
	"net/http"
)

// HandleTasks will handle requests for the /v1/tasks resource
// Always pass Context as a pointer (*), not the entire struct
// 64-bit number vs a large data structure
func (ctx *Context) HandleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		decoder := json.NewDecoder(r.Body)

		// & is used to "get an address" to the variable
		newtask := &tasks.NewTaks{}

		// perform the decoode and check for error
		if err := decoder.Decode(newtask); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		if err := newtask.Validate(); err != nil {
			http.Error(w, "error validating task: " + err.Error(), http.StatusBadRequest)
			return
		}

		ctx.TasksStore.Insert(newtask)
		if err != nil {
			http.Error(w, "error inserting task: " + err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().add(headerContentType, contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(task)
	case "GET":
		// Returns slice of tasks and error
		tasks, err := ctx.TasksStore.GetAll()
		if err != nil {
			http.Error(w, "error getting tasks: " + err.Error(), http.StatusInternalServerError)
			return 
		}
		w.Header().add(headerContentType, contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(task)
	}

}

//HandleSpecificTask will handle requests for the /v1/tasks/some-task-id resource
func (ctx *Context) HandleSpecificTask(w http.ResponseWriter, r *http.Request) {

	// _ will store the directory, which we really don't care about
	_, id := path.Split(r.URL.Path)

	swith r.Method {
	case "GET": 
		task. err := ctx.TasksStore.Get(id)
		if err != nil {
			http.Error(w, "error finding task: " + err.Error(), http.StatusBadRequest)
			return
		}
		// let them know that it is json 
		w.Header().add(headerContentType, contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(task)
	case "PATCH":
		decoder := json.NewDecoder(r.Body)
		task := &tasks.Task{}
		if err := decoder.Decode(task); err != nil {
			http.Error(w, "error decoding JSON: " + err.Error(), http.StatusBadRequest)
			return
		}
		task.ID = id

		if err := ctx.TasksStore.Update(task)l err != nil {
			http.Error(w, "error updating: " + error Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("update Successful!"))
	}

}
