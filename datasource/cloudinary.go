package datasource

import (
	"github.com/cloudinary/cloudinary-go/v2"
)

var Cld *cloudinary.Cloudinary

type Storage struct {
    
}

func InitCloudinary() {
	var err error
	Cld, err = cloudinary.NewFromParams(Config.CLD_CLOUD, Config.CLD_KEY, Config.CLD_SECRET)
	if err != nil {
		panic(err)
	}
}
