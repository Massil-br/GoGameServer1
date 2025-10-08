package models



type Item struct{
	Model
	Name string `json:"name"`
	CanUse bool `json:"canUse"`
	CanAttack bool `json:"canAttack"`
	
} 