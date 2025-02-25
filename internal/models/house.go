package models

type House struct {
    ID       int    `json:"id"`
    Address  string `json:"address"`
    OwnerID  int    `json:"owner_id"`
} 