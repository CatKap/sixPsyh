package handlers 
import (
    "encoding/json"
    "net/http"
		"fmt"
		
		"strings"
    //loger "github.com/CatKap/sixPsyh/loger"
    "github.com/CatKap/sixPsyh/models"

)

func ValidateArg(argument string) (valid string){
	const allowed string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_"
    for _, ch := range argument {
        if !strings.ContainsRune(allowed, ch) {
            return ""
        }
    }
	
	return argument
}

func (h *Handler) Cathegorys(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		querry := models.New(h.db)
		// Meow, meow
		cats, err := querry.GetCathegorys(ctx)	
		if err != nil{
    	w.WriteHeader(http.StatusInternalServerError)
			return
		}
			
		b, err := json.Marshal(cats)
		if err != nil {
    	w.WriteHeader(http.StatusInternalServerError)
			return
		}

    w.WriteHeader(http.StatusOK)
    w.Write(b)
}


func (h *Handler) NewCathegory(w http.ResponseWriter, r *http.Request) {
		type CathegoryJson struct {
			Name string  `json:"name"`
		}
		var cath CathegoryJson
		err := json.NewDecoder(r.Body).Decode(&cath);
		ctx := r.Context()
		querry := models.New(h.db)
		// Meow, meow
		newid, err := querry.NewCathegory(ctx, cath.Name)	

		if err != nil{
    	w.WriteHeader(http.StatusInternalServerError)
			return
		}
			
		b, err := json.Marshal(newid)
		if err != nil {
    	w.WriteHeader(http.StatusInternalServerError)
			return
		}

    w.WriteHeader(http.StatusOK)
    w.Write(b)
}


func (h *Handler) GetEvents(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		querry := models.New(h.db)
		events, err := querry.GetEvents(ctx)	
		if err != nil{
    	w.WriteHeader(http.StatusInternalServerError)
			return
		}
			
    w.WriteHeader(http.StatusOK)
		b, err := json.Marshal(events)
		if err != nil {
    	w.WriteHeader(http.StatusInternalServerError)
			return
		}
    w.Write(b)
}


func (h *Handler) AddEvent(w http.ResponseWriter, r *http.Request){

	type NewEvent struct {
		Name        string `json:"name"`
		Description string `json:"desc"`
		Cathegory   string `json:"ctg"`
		Time        int `json:"time"`
	}

	var  event NewEvent
	err := json.NewDecoder(r.Body).Decode(&event);

	if  err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err)))
		return
	}
	
	ctx := r.Context()
	
	params := models.NewEventParams{
	    Name:        event.Name,
	    Description: event.Description,
	    Name_2:   event.Cathegory,
	    Time:        event.Time, // make sure type matches what DB expects
	}
		
	querry := models.New(h.db)
	id, err := querry.NewEvent(ctx, params)
	if err != nil {
	    http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf(`{"id": %d}`, id)))

}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {

		type DeleteJson struct {
			Table string `json:"table"`
			Ids []int   `json:"ids"`
	 	}

		var del DeleteJson
		err := json.NewDecoder(r.Body).Decode(&del)
		if err != nil {
		    http.Error(w, err.Error(), http.StatusInternalServerError)
				return
		}
		ids := del.Ids

		placeholders := make([]string, len(ids))
		for i := range ids {
		    placeholders[i] = "?"
		}
		inClause := strings.Join(placeholders, ",")
		args := make([]interface{}, len(ids))

		for i, v := range ids {
    args[i] = v
		}

		ctx := r.Context()
		table := ValidateArg(del.Table)
		if table == ""{
		    http.Error(w, "Wrong table name!", http.StatusInternalServerError)
				return
		}
		query := fmt.Sprintf("DELETE FROM %s WHERE id IN (%s)", table, inClause)

		h.log.Info(query)
		h.log.Info(args)
		result, err := h.db.ExecContext(ctx, query, args...)
		if err != nil {
		    http.Error(w, err.Error(), http.StatusInternalServerError)
				return
		}
		deleted, _ := result.RowsAffected()
		h.log.Info("deleted %d rows", deleted)
    w.WriteHeader(http.StatusOK)
		w.Write([]byte(`"status":"OK"`))
}
