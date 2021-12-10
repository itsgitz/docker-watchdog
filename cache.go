package main

import (
	"os"

	"github.com/docker/docker/api/types"
)

const (
	cacheFile = "/tmp/docker-watchdog.email.cache"
)

//readCache read the email cache file content
func readCache() ([]byte, error) {
	jwtString, err := os.ReadFile(cacheFile)
	if err != nil {
		return nil, err
	}

	return jwtString, nil
}

//writeCache write or overwrite a new cache file
func writeCache(jwtString string) error {
	err := os.WriteFile(cacheFile, []byte(jwtString), 0666)
	if err != nil {
		return err
	}

	return nil
}

//isCacheFileExists check the email cache file in /tmp directory
//return false is not exist, true if exist
func isCacheFileExists(filename string) bool {
	file, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}

	return !file.IsDir()
}

//isCached check if email already sended by get the cache content (jwt)
//This will parse the jwt and get the claims data.
//Claims data example:
// {
//		"containers": [
//			"id": "{$container_id}",
//			"state": "{$container_statue}"
//		]
// }
func isCached(containers []types.Container, config *Config) (bool, error) {
	//If cache file is not exists, create a new cache file
	//and write the token on new cache file
	if !isCacheFileExists(cacheFile) {
		//Encode containers data to string
		tokenString, err := encodeToJWT(containers, config)

		//return error from encoding action
		if err != nil {
			return false, err
		}

		//return error from writing file action
		err = writeCache(*tokenString)
		if err != nil {
			return false, err
		}

		//return false if cache file is not exist
		//should send email alert
		return false, nil
	}

	//If cache file exist, check if the existing cache (token)
	//is equal to new token

	//encode a new data from current containers
	newToken, err := encodeToJWT(containers, config)
	if err != nil {
		return false, err
	}

	isTokenEqual, err := compareContainersData(*newToken)
	if err != nil {
		return false, err
	}

	//If not equal, write a new cache
	//return false, should send email alert
	if !isTokenEqual {
		err = writeCache(*newToken)
		if err != nil {
			return false, err
		}

		return false, nil
	}

	//token equal, this means cached
	//return true, don't send email alert
	return true, nil
}

func compareContainersData(newToken string) (bool, error) {
	//get the jwt data
	tokenString, err := readCache()
	if err != nil {
		return false, err
	}

	//if current token from cache file is not equal to a new token,
	//return false (should send email alert)
	if string(tokenString) != newToken {
		return false, nil
	}

	//return true if token is equal
	//should not send email alert, this means the sended email action is cached
	return true, nil
}
