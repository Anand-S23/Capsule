package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Anand-S23/capsule/internal/models"
)

func (c *Controller) CreateConnection(w http.ResponseWriter, r *http.Request) error {
    currentUserID := r.Context().Value("user_id").(string)

    var connectionData models.ConnectionDto
    err := json.NewDecoder(r.Body).Decode(&connectionData)
    if err != nil {
        return WriteJSON(w, http.StatusBadRequest, "Could not parse connection data")
    }

    // TODO: Validation on create connection data 

    connection := models.NewConnection(connectionData, currentUserID)
    err = c.store.ConnectionRepo.Add(connection)
    if err != nil {
        log.Println("Error while creating connection :: ", err)
        return WriteJSON(w, http.StatusInternalServerError, ErrMsg("Internal server error, please try again later"))
    }

    return WriteJSON(w, http.StatusOK, connection)
}

func (c *Controller) GetConnection(w http.ResponseWriter, r *http.Request) error {
    currentUserID := r.Context().Value("user_id").(string)

    id := r.PathValue("id")
    if id == "" {
        return WriteJSON(w, http.StatusInternalServerError, ErrMsg("Could not parse id from path"))
    }

    connection, err := c.store.ConnectionRepo.GetOneByID(id)
    if err != nil {
        log.Println("Error while getting connection by id :: ", err)
        return WriteJSON(w, http.StatusInternalServerError, ErrMsg("Internal server error, please try again later"))
    }

    if connection.OwnerID != currentUserID {
        return WriteJSON(w, http.StatusUnauthorized, ErrMsg("Unauthorized"))
    }

    return WriteJSON(w, http.StatusOK, connection)
}

func (c *Controller) GetAllConnections(w http.ResponseWriter, r *http.Request) error {
    currentUserID := r.Context().Value("user_id").(string)

    connections, err := c.store.ConnectionRepo.GetAllByOwnerID(currentUserID)
    if err != nil {
        log.Println("Error while getting connection by id :: ", err)
        return WriteJSON(w, http.StatusInternalServerError, ErrMsg("Internal server error, please try again later"))
    }

    return WriteJSON(w, http.StatusOK, connections)
}

func (c *Controller) UpdateConnection(w http.ResponseWriter, r *http.Request) error {
    currentUserID := r.Context().Value("user_id").(string)

    id := r.PathValue("id")
    if id == "" {
        return WriteJSON(w, http.StatusInternalServerError, ErrMsg("Could not parse id from path"))
    }

    connectionFromID, err := c.store.ConnectionRepo.GetOneByID(id)
    if err != nil {
        log.Println("Error while getting connection by id :: ", err)
        return WriteJSON(w, http.StatusInternalServerError, ErrMsg("Internal server error, please try again later"))
    }

    if connectionFromID.OwnerID != currentUserID {
        return WriteJSON(w, http.StatusUnauthorized, ErrMsg("Unauthorized"))
    }

    var connectionData models.ConnectionDto
    err = json.NewDecoder(r.Body).Decode(&connectionData)
    if err != nil {
        return WriteJSON(w, http.StatusBadRequest, "Could not parse connection data")
    }

    // TODO: Validation on create connection data 

    connection := models.NewConnectionWithID(connectionData, currentUserID, id)
    err = c.store.ConnectionRepo.Update(connection)
    if err != nil {
        log.Println("Error while creating connection :: ", err)
        return WriteJSON(w, http.StatusInternalServerError, ErrMsg("Internal server error, please try again later"))
    }

    return WriteJSON(w, http.StatusOK, connection)
}

func (c *Controller) DeleteConnection(w http.ResponseWriter, r *http.Request) error {
    currentUserID := r.Context().Value("user_id").(string)

    id := r.PathValue("id")
    if id == "" {
        return WriteJSON(w, http.StatusInternalServerError, ErrMsg("Could not parse id from path"))
    }

    connection, err := c.store.ConnectionRepo.GetOneByID(id)
    if err != nil {
        log.Println("Error while getting connection by id :: ", err)
        return WriteJSON(w, http.StatusInternalServerError, ErrMsg("Internal server error, please try again later"))
    }

    if connection.OwnerID != currentUserID {
        return WriteJSON(w, http.StatusUnauthorized, ErrMsg("Unauthorized"))
    }

    err = c.store.ConnectionRepo.DeleteByID(id)
    if err != nil {
        log.Println("Error while deleting connection by id :: ", err)
        return WriteJSON(w, http.StatusInternalServerError, ErrMsg("Internal server error, please try again later"))
    }

    return WriteJSON(w, http.StatusOK, "OK")
}

