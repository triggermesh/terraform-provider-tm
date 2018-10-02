package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ghodss/yaml"
	"github.com/gosuri/uitable"
	buildApi "github.com/knative/build/pkg/client/clientset/versioned"
	servingApi "github.com/knative/serving/pkg/client/clientset/versioned"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	confPath = "/.tm/config.json"
	kubeConf = "/.kube/config"
)

var (
	debug     bool
	CfgFile   string
	Namespace string
	output    string
	log       logrus.Logger
	core      *kubernetes.Clientset
	build     *buildApi.Clientset
	serving   *servingApi.Clientset
	table     *uitable.Table
)

type confStruct struct {
	Contexts []struct {
		Context struct {
			Cluster   string `json:"cluster"`
			Namespace string `json:"Namespace"`
		} `json:"context"`
		Name string `json:"name"`
	} `json:"contexts"`
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "tm",
	Short:   "Triggermesh CLI",
	Version: "0.0.3",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(InitConfig)
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable debug logging")
	rootCmd.PersistentFlags().StringVar(&CfgFile, "config", "", fmt.Sprintf("k8s config file (default is $HOME%s)", confPath))
	rootCmd.PersistentFlags().StringVarP(&Namespace, "Namespace", "n", "", "User Namespace")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "Output format")
}

func username() string {
	jsonFile, err := os.Open(CfgFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer jsonFile.Close()

	body, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalln(err)
	}
	if body, err = yaml.YAMLToJSON(body); err != nil {
		log.Fatalln(err)
	}

	var conf confStruct
	if err := yaml.Unmarshal(body, &conf); err != nil {
		log.Fatalln(err)
	}
	for _, v := range conf.Contexts {
		if v.Context.Cluster == "triggermesh" {
			return v.Context.Namespace
		}
	}
	return ""
}

func InitConfig() {
	log = *logrus.New()
	log.Out = os.Stdout

	logFormat := new(logrus.TextFormatter)
	logFormat.TimestampFormat = "2006-01-02 15:04:05"
	logFormat.FullTimestamp = true
	log.Formatter = logFormat

	if debug {
		log.Level = logrus.DebugLevel
	}

	table = uitable.New()
	table.Wrap = true
	table.MaxColWidth = 50

	homeDir := "."
	if dir := os.Getenv("HOME"); dir != "" {
		homeDir = dir
	}
	tmHome := filepath.Dir(homeDir + confPath)
	if _, err := os.Stat(tmHome); os.IsNotExist(err) {
		if err := os.MkdirAll(tmHome, 0755); err != nil {
			log.Fatalln(err)
		}
	}

	if len(CfgFile) == 0 {
		CfgFile = homeDir + confPath
	}

	config, err := rest.InClusterConfig()
	if err != nil {
		if config, err = clientcmd.BuildConfigFromFlags("", CfgFile); err != nil {
			if config, err = clientcmd.BuildConfigFromFlags("", homeDir+kubeConf); err != nil {
				log.Fatalln("Can't read config file")
			}
			CfgFile = homeDir + kubeConf
		}
		if len(Namespace) == 0 {
			Namespace = username()
		}
	}

	build, err = buildApi.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	}
	serving, err = servingApi.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	}
	core, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	}
}
