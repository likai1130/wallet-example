package v1

import metav1 "wallet-example/internal/wallet/meta/v1"

type User struct {
}

type UserList struct {
	metav1.ListMeta
	Items []*User `json:"items"`
}
