package routers

/**
 *  
 * @author Sika Kay
 * @date 28/10/17
 *
 */

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	//Users Routes
	router = SetUserRoutes(router)

	//Listings Routes
	router = SetListingRoutes(router)

	//Transactions Routes
	router = SetTransactionRoutes(router)

}
