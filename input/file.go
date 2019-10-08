package input

import (
	"fmt"

	"github.com/spf13/viper"
)

func loadConfig(filePath string) {
	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	*Flags.Host = viper.GetString("host")
	*Flags.Port = viper.GetInt("port")
	*Flags.User = viper.GetString("user")
	*Flags.Passwd = viper.GetString("passwd")
	dbMap := viper.GetStringMapStringSlice("DbMap")
	Flags.DbMap = &dbMap
	*Flags.Table = viper.GetString("Table")
	*Flags.Verbose = viper.GetBool("Verbose")
	*Flags.PackageName = viper.GetString("PackageName")
	*Flags.StructName = viper.GetString("StructName")
	*Flags.JSONAnnotation = viper.GetBool("JSONAnnotation")
	*Flags.GormAnnotation = viper.GetBool("GormAnnotation")
	*Flags.GureguTypes = viper.GetBool("GureguTypes")
	*Flags.TargetFile = viper.GetString("TargetFile")
}
